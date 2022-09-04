package api

import (
	"net/http"

	"github.com/scmn-dev/core/app"
	"github.com/scmn-dev/core/model"
	"github.com/spf13/viper"
)

// FindSamePassword ...
// func FindSamePassword(s db.Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var password model.Password

// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&password); err != nil {
// 			RespondWithError(w, http.StatusBadRequest, InvalidRequestPayload)
// 			return
// 		}
// 		defer r.Body.Close()

// 		schema := r.Context().Value("schema").(string)
// 		urls, err := app.FindSamePassword(s, password, schema)

// 		if err != nil {
// 			RespondWithError(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		RespondWithJSON(w, http.StatusOK, urls)
// 	}
// }

// GeneratePassword generates new password
func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	generatedPass, err := app.GenerateSecureKey(viper.GetInt("server.generatedPasswordLength"))

	if err != nil {
		RespondWithError(w, http.StatusSeeOther, err.Error())
	}

	password := generatedPass
	response := model.Response{
		Code:    http.StatusOK,
		Status:  Success,
		Message: password,
	}

	RespondWithJSON(w, http.StatusOK, response)
}
