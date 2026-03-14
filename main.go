package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortner/handlers"

	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/{short}", handlers.HandleShortLink)

	// TODO implement CORS middleware
	run(mux)
}

// running server on goroutine with graceful shutdown
func run(mux *http.ServeMux) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn't get environment variable", err)
	}

	PORT := os.Getenv("PORT")
	log.Println("Hmmmm..... starting server")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	server := http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Error listening :", err)
		}
		log.Println("Stoped accepting new connections")
	}()

	<-shutdown

	err = server.Shutdown(ctx)
	if err != nil {
		log.Println("Error shuting down server: ", err)
	}
	log.Println("Graceful shutdown achieved")
}
