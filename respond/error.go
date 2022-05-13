package respond

import (
	"encoding/json"
	"github.com/ralugr/language-service/logger"
	"net/http"
)

// errorResponse type, used for storing an error response
type errorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Error creates an error response
func Error(w http.ResponseWriter, status int, error string) {
	resp := errorResponse{
		Success: false,
		Error:   error,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		logger.Warning.Printf("Could not marshal error response %v", err)
		return
	}
}
