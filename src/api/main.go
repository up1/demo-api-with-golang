package main

import (
  "net/http"
  "log"
  "os"
  "context"
	"syscall"
  "os/signal"
  "handlers"
)

func main() {
  log.Print("Starting the service")

  port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  router := handlers.Router()
  srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

  log.Print("The service is ready to listen and serve.")

  killSignal := <-stop
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Server gracefully stopped")
}
