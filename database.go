package main

import (
	"fmt"
	"os"
)
import "github.com/go-pg/pg/v9/orm"
import "github.com/go-pg/pg/v9"

var Database *pg.DB

func StartDatabase() {

	databaseUrl := os.Getenv("DATABASE_URL")

	options, err := pg.ParseURL(databaseUrl)
	if err != nil {
		panic(err)
	}

	Database = pg.Connect(options)

	err = createSchema(Database)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Category)(nil), (*Product)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// var currentID int

// var categories = make(Categories)
// var products = make(Products)

// func init() {
// 	DatabaseCategoryCreate(Category{
// 		Name:        "test1",
// 		DisplayName: "Test Category 1",
// 		Description: "Description 1",
// 	})
// 	DatabaseCategoryCreate(Category{
// 		Name:        "test2",
// 		DisplayName: "Test Category 2",
// 		Description: "Description 2",
// 	})
// }

func DatabaseGetCategories() []Category {
	var categories []Category

	err := Database.Model(&categories).Select()
	if err != nil {
		panic(err)
	}

	return categories
}

func DatabaseCategoryCreate(newCategory Category) Category {
	inserted, err := Database.Model(newCategory).Returning("*").Insert()
	if err != nil {
		panic(err)
	}
	fmt.Print(inserted)
	return newCategory
}

//
// func DatabaseCategoryFindOne(id int) Category {
// 	return categories[id]
// }
//
// func DatabaseCategoryFindAndDelete(id int) error {
// 	category := DatabaseCategoryFindOne(id)
// 	if category.Id != id {
// 		return fmt.Errorf("Could not find category with id %d to delete", id)
// 	}
// 	delete(categories, id)
// 	return nil
// }
