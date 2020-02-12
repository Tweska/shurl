package main

import (
	"database/sql"
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
	params := r.URL.Query()
	fmt.Printf(params.Get("longURL"))

	longURL := r.URL.Query().Get("longURL")
	hash := randomString(5)

	db, _ := sql.Open("sqlite3", "./redirections.db")
	stmt, _ := db.Prepare("INSERT INTO redirections(hash, url) values(?, ?)")
	stmt.Exec(hash, longURL)

	fmt.Fprintf(w, "localhost:8000/%s", hash)
	db.Close()
}

func redirect(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]

	var url string

	db, _ := sql.Open("sqlite3", "./redirections.db")
	stmt, _ := db.Prepare("SELECT url FROM redirections WHERE hash == ?")
	err := stmt.QueryRow(hash).Scan(&url)

	fmt.Printf("%s\n", hash)
	fmt.Printf("%s\n", url)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", 404)
	} else {
		http.Redirect(w, r, url, 301)
	}
	db.Close()
}
