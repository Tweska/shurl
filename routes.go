package main

import (
	"net/http"
)

// Route holds a single route used by the router.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of routes.
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AddRedirect",
		"GET",
		"/add/",
		AddRedirect,
	},
	Route{
		"Redirect",
		"GET",
		"/{hash:[A-Za-z0-9]+}",
		Redirect,
	},
}
