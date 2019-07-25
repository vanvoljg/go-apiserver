package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json; charset=UTF-8"
)

// CategoriesHandler sends JSON of categories
func CategoriesHandler(response http.ResponseWriter, request *http.Request) {

	categories := DatabaseGetCategories()
	result := {
		"Results": []Category,
	}

	if len(categories) == 0 {

	}

	response.Header().Set(contentType, applicationJson)
	response.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(response).Encode(categories); err != nil {
		log.Fatal(err)
	}
}

func CategoryCreate(response http.ResponseWriter, request *http.Request) {

	var category Category
	// messageLimit needs to be set, so use something reasonable
	messageLimit := int64(1024 * 1024) // 1024 KiB is a reasonable message size

	requestBody, err := ioutil.ReadAll(io.LimitReader(request.Body, messageLimit))
	if err != nil {
		log.Fatal(err)
	}

	if err := request.Body.Close(); err != nil {
		log.Fatal(err)
	}

	// Unmarshal takes a JSON string and parses it into the category created above
	if err := json.Unmarshal(requestBody, &category); err != nil {
		response.Header().Set(contentType, applicationJson)
		response.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(response).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	newCategory := DatabaseCategoryCreate(category)

	response.Header().Set(contentType, applicationJson)
	response.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(response).Encode(newCategory); err != nil {
		log.Fatal(err)
	}
}
