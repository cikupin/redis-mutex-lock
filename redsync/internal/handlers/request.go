package handlers

import (
	"log"
	"net/http"

	"github.com/cikupin/redis-mutex-lock/redsync/internal/api"
	"github.com/cikupin/redis-mutex-lock/redsync/internal/constants"
	"github.com/cikupin/redis-mutex-lock/redsync/internal/repositories"
)

type (
	// IRequestHandler defines interface for request handler
	IRequestHandler interface {
		OK(wr http.ResponseWriter, req *http.Request)
		GetData(wr http.ResponseWriter, req *http.Request)
		SetData(wr http.ResponseWriter, req *http.Request)
		GetDataWithThunderingHerdUpdate(wr http.ResponseWriter, req *http.Request)
	}

	// RequestHandler defines struct for request handler
	RequestHandler struct {
		cacheRepo repositories.ICache
	}
)

// NewRequestHandler create instance of request handler
func NewRequestHandler(c repositories.ICache) *RequestHandler {
	return &RequestHandler{
		cacheRepo: c,
	}
}

// OK defines OK handler
func (r *RequestHandler) OK(wr http.ResponseWriter, req *http.Request) {
	log.Println("ok was requested...")

	api.GeneralSuccess.WriteJSON(wr)
	return
}

// GetData will get data from cache
func (r *RequestHandler) GetData(wr http.ResponseWriter, req *http.Request) {
	log.Println("get-data was requested...")

	data, err := r.cacheRepo.GetCache(constants.UserCacheKey)
	if err != nil {
		panic(err)
	}

	resp := api.GeneralSuccess
	resp.Message = data
	resp.WriteJSON(wr)
	return

}

// SetData will set data to cache
func (r *RequestHandler) SetData(wr http.ResponseWriter, req *http.Request) {
	log.Println("set-data was requested...")

	err := r.cacheRepo.UpdateCache(constants.UserCacheKey)
	if err != nil {
		panic(err)
	}

	resp := api.GeneralSuccess
	resp.Message = "cache has been set"
	resp.WriteJSON(wr)
	return
}

// GetDataWithThunderingHerdUpdate will get data from database
// If data is not exist, it will save to cache with thundering herd
func (r *RequestHandler) GetDataWithThunderingHerdUpdate(wr http.ResponseWriter, req *http.Request) {
	log.Println("get-data-with-thundering-herd was requested...")

	data, err := r.cacheRepo.GetCacheWithThunderingHerd(constants.UserCacheKey)
	if err != nil {
		panic(err)
	}

	resp := api.GeneralSuccess
	resp.Message = data
	resp.WriteJSON(wr)
	return
}
