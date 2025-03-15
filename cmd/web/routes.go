package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux() // To initialize a servermux or router

	fileServer := http.FileServer(http.Dir("./ui/static/")) // To serve static files within that folder, so we can use them from the template (.tpml) files.

	mux.HandleFunc("/", app.home)                    // Registering a handler or a controller to a route
	mux.HandleFunc("/snippet/view", app.snippetView) // HandleFunc is syntantic sugar of "http.Handle" and "http.HandlerFunc", so it transforms a function into a handler and registers it in the server mux in a single step
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
