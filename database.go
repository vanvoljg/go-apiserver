package main

import (
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

func DatabaseGetCategories() []Category {
	var categories []Category

	if err := database.Model(&categories).Select(); err != nil {
		log.Fatal(err)
	}

	return categories
}

func DatabaseGetOneCategory(id int) Category {
	foundCategory := Category{Id: id}

	if err := database.Select(&foundCategory); err != nil {
		log.Fatal(err)
	}

	return foundCategory
}

func DatabaseCategoryCreate(newCategory Category) Category {

	_, err := database.Model(&newCategory).Returning("*").Insert()

	if err != nil {
		log.Fatal(err)
	}

	return newCategory
}

func DatabaseDeleteCategory(id int) Category {

	deletedCategory := Category{Id: id}

	if _, err := database.Model(&deletedCategory).WherePK().Returning("*").Delete(); err != nil {
		panic(err)
	}

	return deletedCategory
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

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Category)(nil), (*Product)(nil)} {
		if err := db.CreateTable(model, &orm.CreateTableOptions{}); err != nil {
			return err
		}
	}
	return nil
}
