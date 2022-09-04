package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/scmn-dev/core/model"

	"github.com/scmn-dev/core/app"
)

// SetArgs ...
func SetArgs(r *http.Request, fields []string) (map[string]string, map[string]int) {
	// String type query params
	search := r.FormValue("Search")
	sort := r.FormValue("Sort")
	order := r.FormValue("Order")
	argsStr := map[string]string{
		"search": search,
		"order":  setOrder(fields, sort, order),
	}

	// Integer type query params
	offset := r.FormValue("Offset")
	limit := r.FormValue("Limit")
	argsInt := map[string]int{
		"offset": setOffset(offset),
		"limit":  setLimit(limit),
	}

	return argsStr, argsInt
}

// Offset returns the starting number of result for pagination
func setOffset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)

	if err != nil {
		return -1
	}

	// don't allow negative values
	// except -1 which cancels offset condition
	if offsetInt < 0 {
		offsetInt = -1
	}

	return offsetInt
}

// Limit returns the number of result for pagination
func setLimit(limit string) int {
	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		// -1 cancels limit condition
		return -1
	}

	// min limit should be 1
	if limitInt < 1 {
		limitInt = 1
	}

	return limitInt
}

// SortOrder returns the string for sorting and ordering data
func setOrder(fields []string, sort, order string) string {
	orderValues := []string{"desc", "asc"}

	if include(fields, ToSnakeCase(sort)) && include(orderValues, ToSnakeCase(order)) {
		return ToSnakeCase(sort) + " " + ToSnakeCase(order)
	}

	return "updated_at desc"
}

// include ...
func include(vs []string, t string) bool {
	return app.FindIndex(vs, t) >= 0
}

// ToSnakeCase changes string to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// ToPayload unmarshal request body to payload
func ToPayload(r *http.Request) (model.Payload, error) {
	var payload model.Payload
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		return model.Payload{}, err
	}

	return payload, nil
}

// ToBody decrypts payload data and updates r.Body
func ToBody(r *http.Request, env, transmissionKey string) error {
	// Check environment
	if env == "dev" {
		return nil
	}

	// Unmarshall r.Body to model.Payload
	var payload model.Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		return err
	}

	// Decrypt payload
	dec, err := app.DecryptPayload(transmissionKey, []byte(payload.Data))
	if err != nil {
		return err
	}

	// Update r.Body
	r.Body = ioutil.NopCloser(strings.NewReader(string(dec)))

	return nil
}
