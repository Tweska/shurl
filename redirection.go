package main

// Redirection holds the redirection data.
type Redirection struct {
	Hash string
	URL  string
}

var emptyRedirection = Redirection{
	"",
	"",
}
