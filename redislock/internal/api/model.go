package api

import (
	"encoding/json"
	"net/http"
)

// Response structure for API response
type Response struct {
	HTTPStatusCode int         `json:"-"`
	Code           int         `json:"code"`
	Message        interface{} `json:"message"`
}

// WriteJSON write response in JSON format
func (r Response) WriteJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.HTTPStatusCode)
	json.NewEncoder(w).Encode(r)
}
