package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // To initialize a servermux or router
	mux.HandleFunc("/", home) // Registering a handler or a controller to a route
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server in :4000")
	err := http.ListenAndServe(":4000", mux) // To start a web server with the router
	log.Fatal(err)
}
