package utils

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type BaseEndPoint struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message "`
	Response interface{} `json:"response"`
}

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, result bool, message string, payload interface{}) {
	response, err := json.Marshal(&BaseEndPoint{
		Success:  result,
		Message:  message,
		Response: payload,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, success bool, message string, messageBody string) {
	RespondJSON(w, code, success, message, map[string]string{"error": messageBody})
}

// RespondSuccess makes success message response with payload as json format
func RespondSuccess(w http.ResponseWriter, code int, success bool, message string, messageBody string) {
	RespondJSON(w, code, success, message, map[string]string{"success": messageBody})
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
