package api

import (
	"net/http"
	"strconv"

	"github.com/scmn-dev/core/app"
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/model"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

const (
	// SubscriptionDeleteSuccess represents message when deletind subscription successfully
	SubscriptionDeleteSuccess = "Subscription deleted successfully!"
)

// PostSubscription ...
func PostSubscription(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 0. API Key Check
		keys, ok := r.URL.Query()["api_key"]

		if !ok || len(keys[0]) < 1 {
			RespondWithError(w, http.StatusBadRequest, "API Key is missing")
			return
		}

		if keys[0] != viper.GetString("server.apiKey") {
			RespondWithError(w, http.StatusUnauthorized, "API Key is wrong")
			return
		}

		if err := r.ParseForm(); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Could not parse form.")
			return
		}

		var code int
		var msg string

		switch r.FormValue("alert_name") {
		case "subscription_created":
			code, msg = app.CreateSubscription(s, r)
		case "subscription_updated":
			code, msg = app.UpdateSubscription(s, r)
		case "subscription_cancelled":
			code, msg = app.CancelSubscription(s, r)
		case "subscription_payment_succeeded":
			code, msg = app.PaymentSucceedSubscription(s, r)
		case "subscription_payment_failed":
			code, msg = app.PaymentFailedSubscription(s, r)
		default:
			RespondWithError(w, http.StatusBadRequest, "unknown alert_name")
			return
		}

		if code != http.StatusOK {
			RespondWithError(w, code, msg)
			return
		}

		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: msg,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}

// FindAllSubscriptions ...
func FindAllSubscriptions(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var subscriptionList []model.Subscription

		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// Encrypt payload
		var payload model.Payload
		key := r.Context().Value("transmissionKey").(string)
		encrypted, err := app.EncryptJSON(key, subscriptionList)

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		payload.Data = string(encrypted)

		RespondWithJSON(w, http.StatusOK, payload)
	}
}

// FindSubscriptionByID ...
func FindSubscriptionByID(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		subscription, err := s.Subscriptions().FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		// Decrypt subscription side encrypted fields
		decSubscription, err := app.DecryptModel(subscription)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		subscriptionDTO := model.ToSubscriptionDTO(decSubscription.(*model.Subscription))

		// Encrypt payload
		var payload model.Payload
		key := r.Context().Value("transmissionKey").(string)
		encrypted, err := app.EncryptJSON(key, subscriptionDTO)

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		payload.Data = string(encrypted)

		RespondWithJSON(w, http.StatusOK, payload)
	}
}

// CreateSubscription ...
func CreateSubscription(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := ToPayload(r)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, InvalidRequestPayload)
			return
		}

		defer r.Body.Close()

		// Decrypt payload
		var subscriptionDTO model.SubscriptionDTO
		key := r.Context().Value("transmissionKey").(string)
		err = app.DecryptJSON(key, []byte(payload.Data), &subscriptionDTO)

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		createdSubscription, err := s.Subscriptions().Create(model.ToSubscription(&subscriptionDTO))

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		createdSubscriptionDTO := model.ToSubscriptionDTO(createdSubscription)

		// Encrypt payload
		encrypted, err := app.EncryptJSON(key, createdSubscriptionDTO)

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		payload.Data = string(encrypted)

		RespondWithJSON(w, http.StatusOK, payload)
	}
}

// DeleteSubscription ...
func DeleteSubscription(s db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		subscription, err := s.Subscriptions().FindByID(uint(id))

		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = s.Subscriptions().Delete(subscription.ID)

		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		response := model.Response{
			Code:    http.StatusOK,
			Status:  Success,
			Message: SubscriptionDeleteSuccess,
		}

		RespondWithJSON(w, http.StatusOK, response)
	}
}
