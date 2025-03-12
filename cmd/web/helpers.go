package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Writes an error message and stack trace to the error log and sends a 500 internal server error response to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack()) // debug.Stack" gets a stack trace for the current goroutine
	// app.errorLog.Println(trace)
	app.errorLog.Output(2, trace) // Replaces "app.errorLog.Println" to obtain the previous before the last one's stack message and better locate where the error was originated from. So in the error trace, "helpers.go" is not mentioned, but the file that is calling it, which may be "handlers.go"

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status) // "http.StatusText" transforms a status code into a message
}

func (app *application) notFoundError(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
