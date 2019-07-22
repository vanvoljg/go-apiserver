package main

import "fmt"

var currentID int

var categories = make(Categories)
var products = make(Products)

func init() {
	DatabaseCategoryCreate(Category{
		Name:        "test1",
		DisplayName: "Test Category 1",
		Description: "Description 1",
	})
	DatabaseCategoryCreate(Category{
		Name:        "test2",
		DisplayName: "Test Category 2",
		Description: "Description 2",
	})
}

func DatabaseCategoryCreate(newCategory Category) Category {
	currentID++
	newCategory.ID = currentID
	categories[currentID] = newCategory
	return newCategory
}

func DatabaseCategoryFindOne(id int) Category {
	return categories[id]
}

func DatabaseCategoryFindAndDelete(id int) error {
	category := DatabaseCategoryFindOne(id)
	if category.ID != id {
		return fmt.Errorf("Could not find category with id %d to delete", id)
	}
	delete(categories, id)
	return nil
}
