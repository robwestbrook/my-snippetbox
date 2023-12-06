package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Home handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("In Home handler")
	// Check for exact path "/"
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containg paths to template files
	// File path relative to ROOT DIR
	files := []string{
		"./ui/html/home.page.go.tmpl",
		"./ui/html/base.layout.go.tmpl",
		"./ui/html/footer.partial.go.tmpl",
	}

	// Read files, parse, and check for errors
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Execute template set snd write content to response
	// Check for errors
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Show Snippet Handler
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("In showSnippet handler")
	// Get ID from query string and convert to integer.
	// If it can't be converted to integer or less
	// than 1, return 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.serverError(w, err)
		app.notFound(w)
		return
	}
	app.infoLog.Println("showSnippet: Route accessed with ID:", id)
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
	
}

//
// Create Snippet Handler
// A method of the application type
//
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("In createSnippet handler")
	// Create must be a POST request
	// Check for POST method
	// If not a POST, set headers
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	
	// Dummy data
	title := "Time Test"
	content := "This is a test of the expires date and time"
	expires := "30"

	// Pass the data to the SnippetModel Insert() method
	// Receive the ID of the created snippet
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect to new snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}