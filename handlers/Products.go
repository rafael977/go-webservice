package handlers

import (
	"encoding/json"
	"go-webservice/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	enc := json.NewEncoder(rw)
	err := enc.Encode(lp)
	if err != nil {
		http.Error(rw, "Cannot marshal espone", http.StatusInternalServerError)
	}
}
