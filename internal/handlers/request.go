package handlers

import (
	"log"
	"net/http"

	"github.com/cikupin/redis-mutex-lock/internal/api"
)

type (
	// IRequestHandler defines interface for request handler
	IRequestHandler interface {
		OK(w http.ResponseWriter, r *http.Request)
	}

	// RequestHandler defines struct for request handler
	RequestHandler struct {
	}
)

// NewRequestHandler create instance of request handler
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

// OK defines OK handler
func (r *RequestHandler) OK(wr http.ResponseWriter, req *http.Request) {
	log.Println("OK was requested...")
	api.GeneralSuccess.WriteJSON(wr)
	return
}
