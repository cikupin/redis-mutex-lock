package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RestAPIRouter is the struct that act as router for rest API server
type RestAPIRouter struct {
	Handlers map[string]http.HandlerFunc
}

// NewRestAPIRouter create a new rest API router
func NewRestAPIRouter(c map[string]http.HandlerFunc) *RestAPIRouter {
	return &RestAPIRouter{
		Handlers: c,
	}
}

// GetHandler is the pointer receiver of `RestAPIRouter` struct.
// This function is the function that used to get the router handler.
func (r *RestAPIRouter) GetHandler() (http.Handler, error) {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", r.Handlers["ok"]).Methods(http.MethodGet)
	muxRouter.HandleFunc("/get-data", r.Handlers["get-data"]).Methods(http.MethodGet)
	muxRouter.HandleFunc("/set-data", r.Handlers["set-data"]).Methods(http.MethodGet)

	return muxRouter, nil
}
