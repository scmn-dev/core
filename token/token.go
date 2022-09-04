package token

import (
	"strings"
	"net/http"
	"encoding/json"

	"github.com/scmn-dev/core/logger"
	"github.com/scmn-dev/core/constants"
)

func Find(r *http.Request) string {
	token := ExtractFromCookie(r)
	if token == "" {
		token = ExtractFromHeader(r)
	}

	return token
}

func ExtractFromCookie(r *http.Request) string {
	c, err := r.Cookie(constants.CookieName)
	if err != nil {
		return ""
	}

	return c.Value
}

func ExtractFromHeader(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func ExtractRefreshToken(r *http.Request) string {
	mapToken := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&mapToken); err != nil {
		logger.Errorf("Error while extracting refresh token from body: %v", err)
		return ""
	}

	defer r.Body.Close()

	return mapToken["refresh_token"]
}
