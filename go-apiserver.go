package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"})

	port := os.Getenv("PORT")
	url := ":" + port

	Router := NewRouter()

	StartDatabase()

	fmt.Printf("Server Running on %v\n", port)
	log.Fatal(http.ListenAndServe(url, handlers.CORS(originsOk, headersOk, methodsOk)(Router)))
}
