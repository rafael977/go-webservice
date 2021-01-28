package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-webservice/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	enc := json.NewEncoder(rw)
	err := enc.Encode(lp)
	if err != nil {
		http.Error(rw, "Cannot marshal espone", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.UpdateProduct(id, prod)
}

type KeyProduct struct{}

func (p *Products) ValidateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		var prod *data.Product
		if err := dec.Decode(&prod); err != nil {
			p.l.Fatalln("Unable to read body")
			http.Error(rw, "Unable to read body", http.StatusBadRequest)
		}
		p.l.Printf("%#v", prod)

		if err := prod.Validate(); err != nil {
			http.Error(rw, fmt.Sprintf("Error validting data %s", err), http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
