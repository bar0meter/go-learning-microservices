package handlers

import (
	"log"
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

// Products struct
type Products struct {
	l *log.Logger
}

// NewProducts type
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts method
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Request")

	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct method
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Request")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

// UpdateProduct method
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Put Request")
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	err = prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Server error", http.StatusInternalServerError)
		return
	}
}
