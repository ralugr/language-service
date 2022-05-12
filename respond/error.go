package respond

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Error(w http.ResponseWriter, status int, error string) {
	resp := errorResponse{
		Success: false,
		Error:   error,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}
