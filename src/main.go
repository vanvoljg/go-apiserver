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
	godotenv.Load()

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"})

	PORT := os.Getenv("PORT")
	fmt.Printf("Server Running on %v", PORT)

	Router := NewRouter()

	url := ":" + PORT
	log.Fatal(http.ListenAndServe(url, handlers.CORS(originsOk, headersOk, methodsOk)(Router)))
}
