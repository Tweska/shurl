package main

type Redirection struct {
	Hash string
	URL  string
}

var EmptyRedirection = Redirection{
	"",
	"",
}
