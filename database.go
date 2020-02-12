package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DBAddRedirection(url string) (Redirection, error) {
	db, err := DBOpen()
	if err != nil {
		return EmptyRedirection, err
	}

	hash := randomString(5)

	existingRedirection, err := DBGetRedirection(hash)
	if err != sql.ErrNoRows && err != nil {
		db.Close()
		return EmptyRedirection, err
	}

	if existingRedirection != EmptyRedirection {
		db.Close()
		return DBGetRedirection(url)
	}

	stmt, err := db.Prepare("INSERT INTO redirections(hash, url) values(?, ?)")
	if err != nil {
		db.Close()
		return EmptyRedirection, err
	}

	_, err = stmt.Exec(hash, url)
	db.Close()
	if err != nil {
		return EmptyRedirection, err
	}

	redirection := Redirection{
		hash,
		url,
	}

	return redirection, nil
}

func DBGetRedirection(hash string) (Redirection, error) {
	var url string

	db, err := DBOpen()
	if err != nil {
		return EmptyRedirection, err
	}

	stmt, err := db.Prepare("SELECT url FROM redirections WHERE hash == ?")
	if err != nil {
		db.Close()
		return EmptyRedirection, err
	}

	err = stmt.QueryRow(hash).Scan(&url)
	db.Close()
	if err != nil {
		return EmptyRedirection, err
	}

	redirection := Redirection{
		hash,
		url,
	}

	return redirection, nil
}

func DBOpen() (*sql.DB, error) {
	return sql.Open("sqlite3", "./redirections.db")
}
