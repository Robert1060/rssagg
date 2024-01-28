package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from headers of request
// Example:
// Authorization: ApiKey {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	authString := headers.Get("Authorization")
	if authString == "" {
		return "", errors.New("no auth info found")
	}
	values := strings.Split(authString, " ")
	if len(values) != 2 {
		return "", errors.New("malformed auth header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return values[1], nil
}
