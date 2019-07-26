package main

import (
	"net/http"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array
type Routes []Route

var routes = Routes{
	Route{
		"GetAllCategories",
		"GET",
		"/categories",
		GetAllCategories,
	},
	Route{
		"CreateCategory",
		"POST",
		"/categories",
		CreateCategory,
	},
	Route{
		"GetOneCategory",
		"GET",
		"/categories/{id}",
		GetOneCategory,
	},
	Route{
		"DeleteCategory",
		"DELETE",
		"/categories/{id}",
		DeleteCategory,
	},
	Route{
		"DatabaseInit",
		"GET",
		"/database/initialize",
		DatabaseInit,
	},
}
