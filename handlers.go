package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const contentType = "Content-Type"
const applicationJson = "application/json; charset=UTF-8"

// CategoriesHandler sends JSON of categories
func CategoriesHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set(contentType, applicationJson)
	response.WriteHeader(http.StatusOK)

	categories := DatabaseGetCategories()

	if err := json.NewEncoder(response).Encode(categories); err != nil {
		panic(err)
	}
}

func CategoriesCreate(response http.ResponseWriter, request *http.Request) {

	var category Category
	messageLimit := int64(1024 * 1024) // 1024 KiB is a reasonable message size

	requestBody, err := ioutil.ReadAll(io.LimitReader(request.Body, messageLimit))
	if err != nil {
		panic(err)
	}

	if err := request.Body.Close(); err != nil {
		panic(err)
	}

	// Unmarshal takes a JSON string and parses it into the category created above
	if err := json.Unmarshal(requestBody, &category); err != nil {
		response.Header().Set(contentType, applicationJson)
		response.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(response).Encode(err); err != nil {
			panic(err)
		}
	}

	newCategory := DatabaseCategoryCreate(category)

	response.Header().Set(contentType, applicationJson)
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(newCategory); err != nil {
		panic(err)
	}
}
