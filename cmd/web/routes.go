package main

import "net/http"

// Routes
func (app *application) routes() *http.ServeMux {
	// Create new MUX
	mux := http.NewServeMux()

	// Create routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server to serve static resources
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}