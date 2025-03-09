package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	// w http.ResponseWriter --> Provides methods for assembling a HTTP response and sending it to the user
	// r *http.Request is a pointer to a struct which holds information about the current request (HTTP method, url, etc)

	// A check to non-existent routes
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// cwd, _ := os.Getwd()
	// log.Println("Current working directory:", cwd) //for debugging errors parsing the html templates

	ts, err := template.ParseFiles(files...) // Reading or parsing the template into a template set
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil) // Writing the template content as the response body and inserting data in the template (Second argument)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("Hello from Snippetbox")) // Placeholder message to write to the html body
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
