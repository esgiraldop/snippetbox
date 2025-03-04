package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Display a specific snippet..."))

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id) // Does the same than w.Write but on top of that, it interpolates the string with variables
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST") // To let the user know which request methods are supported
		w.Header().Set("Deny", "DELETE, PUT, PATCH, GET")

		// w.WriteHeader(405) // To change the default response from 200 to 405 (Not allowed)
		// w.Write([]byte("Method not allowed"))

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Does exactly the same than w.WriteHeader and w.Write combined
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
