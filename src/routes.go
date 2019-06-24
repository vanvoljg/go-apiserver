package main

import (
	"net/http"

	"github.com/gorilla/mux"
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

// NewRouter returns a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

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
		"OneCategory",
		"GET",
		"/categories/{ID}",
		OneCategoryHandler,
	},
}
