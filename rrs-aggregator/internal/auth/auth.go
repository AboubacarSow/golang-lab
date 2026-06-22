package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// Pattern: [Authorization: ApiKey {key_value_here}]
func GetApiKey(header http.Header) (string, error) {
	value := header.Get("Authorization")

	if value == "" {
		return "", fmt.Errorf("Not Authentication Inforamtion Found")
	}
	authValues := strings.Split(strings.TrimSpace(value), " ")

	if len(authValues) != 2 || authValues[0] != "ApiKey" {
		return "", fmt.Errorf("Wrong Authentication Format")
	}
	if authValues[0] != "ApiKey" {
		return "", fmt.Errorf("Wrong Authentication Format:missing ApiKey")
	}
	return authValues[1], nil

}
