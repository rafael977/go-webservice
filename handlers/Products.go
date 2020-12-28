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

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	enc := json.NewEncoder(rw)
	err := enc.Encode(lp)
	if err != nil {
		http.Error(rw, "Cannot marshal espone", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var prod *data.Product
	if err := dec.Decode(&prod); err != nil {
		p.l.Fatalln("Unable to read body")
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
	}
	p.l.Printf("%#v", prod)

	data.AddProduct(prod)
}
