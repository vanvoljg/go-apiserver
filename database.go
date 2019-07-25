package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var database *pg.DB

func StartDatabaseConnection() {

	databaseUrl := os.Getenv("DATABASE_URL")

	options, err := pg.ParseURL(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	database = pg.Connect(options)
}

func CloseDatabaseConnection() {
	if err := database.Close(); err != nil {
		log.Fatal(err)
	}
}

func DatabaseInit(response http.ResponseWriter, request *http.Request) {
	if err := createSchema(database); err != nil {
		log.Fatal(err)
	}

	if _, err := io.WriteString(response, "Database Initialized"); err != nil {
		log.Fatal(err)
	}

	if err := request.Body.Close(); err != nil {
		log.Fatal(err)
	}
}

func DatabaseGetCategories() []Category {
	var categories []Category

	err := database.Model(&categories).Select()
	if err != nil {
		log.Fatal(err)
	}

	return categories
}

func DatabaseCategoryCreate(newCategory Category) Category {
	inserted, err := database.Model(&newCategory).Returning("*").Insert()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted)
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
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Category)(nil), (*Product)(nil)} {
		if err := db.CreateTable(model, &orm.CreateTableOptions{}); err != nil {
			return err
		}
	}
	return nil
}
