package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type encoder struct{}

// errorResponse will encapsulate errors to be transferred over HTTP.
type errorResponse struct {
	Message string      `json:"message"`
	Trace   interface{} `json:"trace,omitempty"`
	Trail   interface{} `json:"trail,omitempty"`
}

func (e encoder) StatusResponse(ctx context.Context, w http.ResponseWriter, response interface{}, status int) {
	if response == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("error encoding/writing response %s", err.Error()), status)
	}
}

func (e encoder) Error(ctx context.Context, w http.ResponseWriter, err error, status int) {
	resp := errorResponse{
		Message: err.Error(),
	}

	w.Header().Set("Content-type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("error encoding/writing response %s", err.Error()), status)
	}
}
