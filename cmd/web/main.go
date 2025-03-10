package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // To initialize a servermux or router

	fileServer := http.FileServer(http.Dir("./ui/static/")) // To serve static files within that folder, so we can use them from the template (.tpml) files.

	mux.HandleFunc("/", home) // Registering a handler or a controller to a route
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server in :4000")
	err := http.ListenAndServe(":4000", mux) // To start a web server with the router
	log.Fatal(err)
}
