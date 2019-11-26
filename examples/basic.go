package main

import (
	"log"
	"net/http"

	"github.com/Tamarou/blackarachnia"
	"github.com/gorilla/mux"
)

type SimpleResource struct{}

func (sr SimpleResource) content() []byte {
	return []byte("Hello World!\n")
}

func main() {
	r := mux.Router()
	resource := SimpleResource{}
	handler := blackarachnia.NewHandler(resource)
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
