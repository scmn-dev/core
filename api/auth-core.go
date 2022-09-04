package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"

	"github.com/scmn-dev/core/app"
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/logger"
	"github.com/scmn-dev/core/model"
)

var (
	verifySuccess = "Email verified successfully"
)

// Signup ...
func Signup(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Decode request body to userDTO object
		userSignup := new(model.UserSignup)
		decoderr := json.NewDecoder(r.Body)

		if err := decoderr.Decode(&userSignup); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid request payload")

			return
		}

		defer r.Body.Close()
		// 2. Check if email is verified
		if err := isMailVerified(userSignup.Email); err != nil {
			logger.Errorf("email %s is not verified error %v\n", userSignup.Email, err)
			RespondWithError(w, http.StatusUnauthorized, "Email is not verified")

			return
		}

		// 2. Run validator according to model.UserDTO validator tags
		err := app.PayloadValidator(userSignup)
		if err != nil {
			errs := GetErrors(err.(validator.ValidationErrors))
			RespondWithErrors(w, http.StatusBadRequest, InvalidRequestPayload, errs)

			return
		}

		// 4. Check if user exist in database
		userDTO := model.ConvertUserDTO(userSignup)
		_, err = s.Users().FindByEmail(userDTO.Email)

		if err == nil {
			RespondWithError(w, http.StatusBadRequest, "User couldn't created!")
			return
		}

		// 5. Create new user
		createdUser, err := app.CreateUser(s, userDTO)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// 6. Send email to admin about new user subscription
		notifyAdminEmail(createdUser)
		// Return success message
		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: signupSuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Create email verification code
func CreateCode(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Decode json to email
		var signup model.AuthEmail
		if err := json.NewDecoder(r.Body).Decode(&signup); err != nil {
			RespondWithError(w, http.StatusBadRequest, InvalidRequestPayload)
			return
		}

		// 2. Check if user exist in database
		_, err := s.Users().FindByEmail(signup.Email)
		if err == nil {
			logger.Errorf("email %s already exist in database\n", signup.Email)
			RespondWithError(w, http.StatusBadRequest, "User couldn't created!")
			return
		}

		// 2. Generate a random code
		rand.Seed(time.Now().Unix())
		min := 100000
		max := 999999
		code := strconv.Itoa(rand.Intn(max-min+1) + min)

		log.Printf("verification code %s generated for email %s\n", code, signup.Email)

		// 3. Save code in cache
		c.Set(signup.Email, code, cache.DefaultExpiration)

		// 4. Send verification email to user
		subject := "$PASSWORD_MANAGER_NAME - Email Verification"
		body := EmailVerification(signup.Email, code)

		if err = app.SendMail("$PASSWORD_MANAGER_NAME Verification Code", signup.Email, subject, body); err != nil {
			logger.Errorf("can't send email to %s error: %v\n", signup.Email, err)
			RespondWithError(w, http.StatusBadRequest, "Couldn't send email")
			return
		}

		// Return success message
		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: codeSuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

func SendFMP(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Decode json to email
		var fmp model.AuthEmail
		if err := json.NewDecoder(r.Body).Decode(&fmp); err != nil {
			RespondWithError(w, http.StatusBadRequest, InvalidRequestPayload)
			return
		}

		// 2. Check if user exist in database
		_, err := s.Users().FindByEmail(fmp.Email)
		if err != nil {
			logger.Errorf("email %s doesn't exist in database\n", fmp.Email)
			RespondWithError(w, http.StatusBadRequest, "User couldn't found!")

			return
		}

		// 2. Generate a random code
		rand.Seed(time.Now().Unix())
		min := 100000
		max := 999999
		code := strconv.Itoa(rand.Intn(max-min+1) + min)

		log.Printf("restore master password verification code %s generated for email %s\n", code, fmp.Email)

		// 3. Save code in cache
		c.Set(fmp.Email, code, cache.DefaultExpiration)

		// 4. Send verification email to user
		subject := "$PASSWORD_MANAGER_NAME - Reset master password instructions"
		body := ResetMasterPasswordInstructions(fmp.Email, code)

		if err = app.SendMail("$PASSWORD_MANAGER_NAME reset master password instructions", fmp.Email, subject, body); err != nil {
			logger.Errorf("can't send email to %s error: %v\n", fmp.Email, err)
			RespondWithError(w, http.StatusBadRequest, "Couldn't send email")
			return
		}

		// Return success message
		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: codeSuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Create user deletion code
func CreateDeleteCode(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Decode json to email
		var signup model.AuthEmail

		if err := json.NewDecoder(r.Body).Decode(&signup); err != nil {
			RespondWithError(w, http.StatusBadRequest, InvalidRequestPayload)
			return
		}

		// 2. Check if user exist in database
		_, err := s.Users().FindByEmail(signup.Email)
		if err != nil {
			logger.Errorf("email %s does not exist in database error %v\n", signup.Email, err)
			RespondWithError(w, http.StatusBadRequest, "User couldn't be found!")
			return
		}

		// 2. Generate a random code
		rand.Seed(time.Now().Unix())
		min := 100000
		max := 999999
		code := strconv.Itoa(rand.Intn(max-min+1) + min)

		logger.Infof("deletion code %s generated for email %s\n", code, signup.Email)

		// 3. Save code in cache
		c.Set(signup.Email, code, cache.DefaultExpiration)

		// 4. Send verification email to user
		subject := "$PASSWORD_MANAGER_NAME User Deletion Verification"
		body := "$PASSWORD_MANAGER_NAME user deletion code: " + "<strong>" + code + "</strong>"

		if err = app.SendMail("$PASSWORD_MANAGER_NAME user deletion Code", signup.Email, subject, body); err != nil {
			logger.Errorf("can't send email to %s error: %v\n", signup.Email, err)
			RespondWithError(w, http.StatusBadRequest, "Couldn't send email")
			return
		}

		// Return success message
		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: codeSuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Verify Email
func VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userCode := mux.Vars(r)["code"]
		email := r.FormValue("email")

		code, ok := c.Get(email)
		if !ok {
			RespondWithError(w, http.StatusBadRequest, "Code couldn't found!")
			return
		}

		confirmationCode, ok := code.(string)
		if !ok {
			RespondWithError(w, http.StatusInternalServerError, "Server error!")
			return
		}

		if userCode != confirmationCode {
			RespondWithError(w, http.StatusBadRequest, "Code doesn't match!")
			return
		}

		c.Set(email, "verified", cache.DefaultExpiration)

		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: verifySuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

func RecoverDelete(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get route variables
		vars := mux.Vars(r)
		// Get email variable
		email := vars["email"]

		// Check if email is verified
		if err := isMailVerified(email); err != nil {
			logger.Errorf("email %s is not verified error %v\n", email, err)
			RespondWithError(w, http.StatusUnauthorized, "Email is not verified")

			return
		}

		// Check if user exist in database
		user, err := s.Users().FindByEmail(email)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// Delete user
		err = s.Users().Delete(user.ID, user.Schema)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		response := model.Response{
			Code:    http.StatusOK,
			Status:  "Success",
			Message: "User deleted successfully!",
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}
