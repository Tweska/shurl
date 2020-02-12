package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DBAddRedirection will create a hash and add a new redirection to the database.
func DBAddRedirection(url string) (Redirection, error) {
	db, err := DBOpen()
	if err != nil {
		return emptyRedirection, err
	}

	hash := RandomString(5)

	existingRedirection, err := DBGetRedirection(hash)
	if err != sql.ErrNoRows && err != nil {
		db.Close()
		return emptyRedirection, err
	}

	if existingRedirection != emptyRedirection {
		db.Close()
		return DBGetRedirection(url)
	}

	stmt, err := db.Prepare("INSERT INTO redirections(hash, url) values(?, ?)")
	if err != nil {
		db.Close()
		return emptyRedirection, err
	}

	_, err = stmt.Exec(hash, url)
	db.Close()
	if err != nil {
		return emptyRedirection, err
	}

	redirection := Redirection{
		hash,
		url,
	}

	return redirection, nil
}

// DBGetRedirection will return the redirection corresponding to the hash.
func DBGetRedirection(hash string) (Redirection, error) {
	var url string

	db, err := DBOpen()
	if err != nil {
		return emptyRedirection, err
	}

	stmt, err := db.Prepare("SELECT url FROM redirections WHERE hash == ?")
	if err != nil {
		db.Close()
		return emptyRedirection, err
	}

	err = stmt.QueryRow(hash).Scan(&url)
	db.Close()
	if err != nil {
		return emptyRedirection, err
	}

	redirection := Redirection{
		hash,
		url,
	}

	return redirection, nil
}

// DBOpen will return a database connection.
func DBOpen() (*sql.DB, error) {
	return sql.Open("sqlite3", "./redirections.db")
}
