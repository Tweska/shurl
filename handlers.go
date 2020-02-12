package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

func addRedirect(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("longURL")

	redirection, _ := DBAddRedirection(longURL)

	fmt.Fprintf(w, "localhost:8000/%s", redirection.Hash)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	redirection, err := DBGetRedirection(hash)

	if err != nil {
		http.Redirect(w, r, "/", 501)
	} else {
		http.Redirect(w, r, redirection.URL, 301)
	}
}
