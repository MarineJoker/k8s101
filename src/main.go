package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the homepage!")
	})

	mux.HandleFunc("/localhost/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've hit the hello route!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// it will be ErrServerClosed when Shutdown or Close is called
			fmt.Fprintf(os.Stderr, "Could not listen on %s: %v\n", server.Addr, err)
			os.Exit(1)
		}
	}()

	// Create a channel to receive OS signals
	signals := make(chan os.Signal, 1)
	// `signal.Notify` will catch SIGINT and SIGTERM signals and relay it to the `signals` channel
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-signals

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout.
	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped")
}

