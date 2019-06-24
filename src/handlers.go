package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CategoriesHandler sends JSON of categories
func CategoriesHandler(response http.ResponseWriter, request *http.Request) {
	categories := Categories{
		Category{
			4,
			"fish",
			"Fish!",
			"Fish that are fish",
		},
		Category{
			5,
			"bob",
			"BOB!",
			"Bob!",
		},
	}
	requestVars := mux.Vars(request)
	switch request.Method {
	case "GET":
		json.NewEncoder(response).Encode(categories)
	case "POST":
		fmt.Println(requestVars)
	}
}

// OneCategoryHandler sends one category by ID
func OneCategoryHandler(response http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)
	json.NewEncoder(response).Encode(requestVars)
}

// Hello World
func Hello(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Hello World")
}
