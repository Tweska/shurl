package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Index will display the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

// AddRedirect will add a new redirection to the database and return the hash.
func AddRedirect(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("longURL")

	redirection, _ := DBAddRedirection(longURL)

	fmt.Fprintf(w, "localhost:8000/%s", redirection.Hash)
}

// Redirect will redirect the user to the url that corresponds to the hash.
func Redirect(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	redirection, err := DBGetRedirection(hash)

	if err != nil {
		http.Redirect(w, r, "/", 501)
	} else {
		http.Redirect(w, r, redirection.URL, 301)
	}
}
