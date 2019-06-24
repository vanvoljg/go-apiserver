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
		"Index",
		"GET",
		"/",
		Hello,
	},
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
		CategoriesCreate,
	},
	// Route{
	// 	"OneCategory",
	// 	"GET",
	// 	"/categories/{ID}",
	// 	OneCategoryHandler,
	// },
}
