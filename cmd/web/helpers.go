package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Server Error
// Writes a message and stack trace
// Sends a 500 error to the user
// Set stack trace output to frame depth of 2
// to get file name and line number one step back
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Client Error
// Writes a specific status code and description
// to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Not Found Error
// Convenience wrapper arounf clientWrapper to send
// 404 Not Found response to user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}