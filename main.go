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

func home(w http.ResponseWriter, r *http.Request) {
	// w http.ResponseWriter --> Provides methods for assembling a HTTP response and sending it to the user
	// r *http.Request is a pointer to a struct which holds information about the current request (HTTP method, url, etc)

	// A check to non-existent routes
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
