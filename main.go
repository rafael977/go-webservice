package main

import (
	"context"
	"go-webservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)

	mux := http.NewServeMux()
	mux.Handle("/", hh)
	mux.Handle("/products", ph)

	ss := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		if err := ss.ListenAndServe(); err != nil {
			l.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := ss.Shutdown(tc); err != nil {
		l.Fatalln("Shutdown failed", err)
	}
}
