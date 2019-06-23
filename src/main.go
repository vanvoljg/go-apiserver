package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"http"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	var router = mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")

	PORT := os.Getenv(PORT)
	fmt.Printf("Server Running on %v", PORT)
}

func helloWorld(request, response, next) {

}
