package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json; charset=UTF-8"
)

// GetAllCategories sends JSON of categories
func GetAllCategories(writer http.ResponseWriter, request *http.Request) {

	results := DatabaseGetCategories()

	categories := map[string][]Category{
		"categories": results,
	}

	writer.Header().Set(contentType, applicationJson)
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(categories); err != nil {
		log.Fatal(err)
	}
}

func GetOneCategory(writer http.ResponseWriter, request *http.Request) {

	idString := mux.Vars(request)["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}

	result := DatabaseGetOneCategory(id)

	writer.Header().Set(contentType, applicationJson)
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}

func CreateCategory(writer http.ResponseWriter, request *http.Request) {

	var category Category

	// Parse json encoded request body
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
		writer.Header().Set(contentType, applicationJson)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	newCategory := DatabaseCategoryCreate(category)

	writer.Header().Set(contentType, applicationJson)
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(newCategory); err != nil {
		log.Fatal(err)
	}
}

func DeleteCategory(writer http.ResponseWriter, request *http.Request) {

	idString := mux.Vars(request)["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}

	deletedCategory := DatabaseDeleteCategory(id)

	writer.Header().Set(contentType, applicationJson)
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(deletedCategory); err != nil {
		log.Fatal(err)
	}

}
