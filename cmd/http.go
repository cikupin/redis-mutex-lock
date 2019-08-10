package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cikupin/redis-mutex-lock/routes"
)

// Exec executes apps
func Exec() {
	app := routes.NewRestAPI()
	defer app.UnloadResources()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":8080",
		Handler: app.Handler,
	}

	go func() {
		log.Println("starting web server on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error while shutting down server")
		os.Exit(1)
	}
	log.Println("Server gracefully stopped")
	os.Exit(0)
}
