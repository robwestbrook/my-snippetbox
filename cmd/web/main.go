package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robwestbrook/snippetbox/pkg/models/sqlite"
)

// Define an application struct
// Holds all application-wide dependencies needed
// throughout the application
// Dependencies:
//	errorLog		error logger
//  infoLog			info logget
//	snippets		snippet model
type application struct {
	errorLog	*log.Logger
	infoLog		*log.Logger
	snippets	*sqlite.SnippetModel
}

func main() {
	// Define command line flags and parse command line
	addr := flag.String("addr", ":8000", "HTTP network address")
	dsn := flag.String("dsn", "snippetbox.sqlite", "Sqlite database name")
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
		snippets: &sqlite.SnippetModel{DB: db},
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
	app.errorLog.Fatal(err)
}

// Wrap sql.Open() and return a connection pool to DB
func openDb(dsn string) (*sql.DB, error) {
	// Create a logger for informational messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		errorLog.Println("Could not open DB")
		return nil, err
	} else {
		infoLog.Println("Database is open.")
	}

	if err = db.Ping(); err != nil {
		errorLog.Println("Could not ping DB")
		return nil, err
	} else {
		infoLog.Println("Database was successfully pinged.")
	}

	res, err := db.Query("SELECT * FROM snippets")
	if err != nil {
		errorLog.Println("Error", err)
	} else {
		infoLog.Println(res)
	}



	return db, nil
}