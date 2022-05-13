package respond

import (
	"encoding/json"
	"github.com/ralugr/language-service/logger"
	"net/http"
)

// successResponse used for storing a success response
type successResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}

// Success creates a new response
func Success(w http.ResponseWriter, response interface{}) {
	resp := successResponse{
		Success:  true,
		Response: response,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if _, err := w.Write(body); err != nil {
		logger.Warning.Printf("Could not create success response %v", err)
	}
}
