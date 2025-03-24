package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Aliasing database drive package with a "_" because anything within the package is used, but the "init()" function is needed to register the driver with the "database/sql" package
)

func openDB(dsn string, infoLog *log.Logger) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn) // Returns a "sql.DB" object, which is not a database connection, but a pool of many connections. Go automatically manages these pool conections as needed, so it supports concurrent conections.

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	} // sql.Open does not establish a connection until is actually used (lazy connection), so to be able to check for any errors during the connection, the "db.Ping" statement is executed.

	infoLog.Printf("SQL connection pool established successfully.")

	return db, nil
}
