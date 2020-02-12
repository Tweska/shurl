package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Set the seed of the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Create a router.
	router := NewRouter()

	// Create a file server.
	fileServer := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Start listening for incomming requests.
	log.Fatal(http.ListenAndServe(":8000", router))
}
