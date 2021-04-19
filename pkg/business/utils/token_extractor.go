package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// TokenExtractor , for extract data inside token
func TokenExtractor(encodedToken string) string {
	if encodedToken != "" {
		encodedPayload := strings.Split(encodedToken, ".")
		decodedPayload, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(encodedPayload[1])

		if err != nil {
			panic(err)
		}

		payload := make(map[string]string)
		err = json.Unmarshal(decodedPayload, &payload)

		if err != nil {
			panic(err)
		}

		return payload["userId"]
	}
	return ""
}
