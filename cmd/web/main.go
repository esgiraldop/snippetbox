package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Defining a struct to hold the application-wide dependencies with dependency injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Defining a command-line flag with default value for 4000 for the port number
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() // Parsing the command-line flag

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // Creating a logger for writting info messages

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // Creating a logger for writting error messages

	// Instantiating my application
	app := application{errorLog, errorLog}

	mux := app.routes()

	// Creating a custom struct for setting up the configuration for the http server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server in %s", *addr)
	// err := http.ListenAndServe(*addr, mux) // To start a web server with the router. If a custom http.Server struct has been set up, this line is not needed anymore, but something like "srv.ListenAndServe()"
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // Should be used only on main()
}
