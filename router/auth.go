package router

import (
	"context"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/scmn-dev/core/app"
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/token"
	"github.com/urfave/negroni"
)

// Auth is a middleware that checks for a valid JWT token
func Auth(s db.Store) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		tokenStr := token.Find(r)

		token, err := app.TokenValid(tokenStr)

		if err != nil {
			if token != nil {
				claims, _ := token.Claims.(jwt.MapClaims)
				uuid, _ := claims["uuid"].(string)
				s.Tokens().DeleteByUUID(uuid)
			}

			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		uuid, _ := claims["uuid"].(string)

		// Get User UUID from claims
		ctxUserUUID, ok := claims["user_uuid"].(string)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Get user details from db by User UUID
		user, err := s.Users().FindByUUID(ctxUserUUID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check token from tokens db table
		tokenRow, err := s.Tokens().FindByUUID(uuid)

		// Token invalidation for old token usage
		if err != nil {
			s.Tokens().Delete(int(user.ID))
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		// Admin or Member
		ctxAuthorized, ok := claims["authorized"].(bool)

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctxSchema := user.Schema
		ctxTransmissionKey := tokenRow.TransmissionKey

		ctx := r.Context()
		ctxWithUUID := context.WithValue(ctx, "uuid", ctxUserUUID)
		ctxWithAuthorized := context.WithValue(ctxWithUUID, "authorized", ctxAuthorized)
		ctxWithSchema := context.WithValue(ctxWithAuthorized, "schema", ctxSchema)
		ctxWithTransmissionKey := context.WithValue(ctxWithSchema, "transmissionKey", ctxTransmissionKey)

		next(w, r.WithContext(ctxWithTransmissionKey))
	})
}
