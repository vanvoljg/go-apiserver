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
		"Categories",
		"GET",
		"/categories",
		CategoriesHandler,
	},
	Route{
		"CreateCategory",
		"POST",
		"/categories",
		CategoryCreate,
	},
	// Route{
	// 	"OneCategory",
	// 	"GET",
	// 	"/categories/{Id}",
	// 	OneCategoryHandler,
	// },
	Route{
		"InitializeDatabase",
		"GET",
		"/database/initialize",
		DatabaseInit,
	},
}
