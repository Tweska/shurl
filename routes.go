package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index,
	},
	Route{
		"AddRedirect",
		"GET",
		"/add/",
		addRedirect,
	},
	Route{
		"Redirect",
		"GET",
		"/{hash:[A-Za-z0-9]+}",
		redirect,
	},
}
