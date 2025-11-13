package util

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"
)

type ErrorResponse struct {
	Message    string    `json:"message"`
	StatusCode int       `json:"status_code"`
	Timestamp  time.Time `json:"timestamp"`
}

func ToJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func FromJSON(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func JSONError(w http.ResponseWriter, status int, err error) {
	response := &ErrorResponse{
		Message:    err.Error(),
		StatusCode: status,
		Timestamp:  time.Now(),
	}

	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON error response", http.StatusInternalServerError)
	}
}
