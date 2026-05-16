package main

import (
	"log"
	"net/http"

	"github.com/faysal0x1/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", nil))

}
