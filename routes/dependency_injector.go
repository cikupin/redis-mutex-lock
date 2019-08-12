package routes

import (
	"log"
	"net/http"
	"sync"

	"github.com/cikupin/redis-mutex-lock/drivers"
	"github.com/cikupin/redis-mutex-lock/internal/handlers"
	"github.com/cikupin/redis-mutex-lock/internal/repositories"
)

// RestAPI is a type of struct that contain rest API handler
type RestAPI struct {
	Handler http.Handler
}

var (
	redisDriver *drivers.RedisPool
	once        sync.Once
)

// NewRestAPI is the function to make new RestAPI type and fill it with the handler
func NewRestAPI() *RestAPI {
	initializeDrivers()

	handler, err := makeHandler()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &RestAPI{
		Handler: handler,
	}
}

func initializeDrivers() {
	once.Do(func() {
		redisDriver, _ = drivers.NewRedisConn()
	})
}

func makeHandler() (http.Handler, error) {
	cacheRepo := repositories.NewCacheRepo(redisDriver)
	requestHandler := handlers.NewRequestHandler(cacheRepo)

	r := NewRestAPIRouter(make(map[string]http.HandlerFunc))
	r.Handlers["ok"] = requestHandler.OK
	r.Handlers["get-data"] = requestHandler.GetData
	r.Handlers["set-data"] = requestHandler.SetData
	r.Handlers["update-with-thundering-herd"] = requestHandler.GetDataWithThunderingHerdUpdate

	return r.GetHandler()
}

// UnloadResources will unload any used resources
func (rest *RestAPI) UnloadResources() {
}
