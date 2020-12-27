package main

import (
	"go-webservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)

	mux := http.NewServeMux()
	mux.Handle("/", hh)

	http.ListenAndServe(":9090", mux)
}
