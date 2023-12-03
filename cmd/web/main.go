package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Define an application struct
// Holds all application-wide dependencies needed
// throughout the application
type application struct {
	errorLog	*log.Logger
	infoLog		*log.Logger
}

func main() {
	// Define command line flags and parse command line
	addr := flag.String("addr", ":8000", "HTTP network address")
	dsn := flag.String("dsn", "snippetbox", "Sqlite database name")
	flag.Parse()

	// Create a logger for informational messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create database connection pool using name
	// from command line flag.
	db, err := openDb(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Defer a call to db.Close(), so the connection pool
	// closes before the main function exits
	defer db.Close()

	// Initialize a new instance of application
	// containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	// Initialize a new http.Server struct so that
	// the server users the custom errorLog
	srv := &http.Server{
		Addr:			*addr,
		ErrorLog: errorLog,
		Handler: 	app.routes(),
	}

	// Start server
	infoLog.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Wrap sql.Open() and return a connection pool to DB
func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}