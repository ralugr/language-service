package respond

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}

func Success(w http.ResponseWriter, response interface{}) {
	resp := successResponse{
		Success:  true,
		Response: response,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}
