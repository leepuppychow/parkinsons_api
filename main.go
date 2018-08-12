// Re-creating backend for Parkinson's App in Go

package main

import (
	"fmt"
	"log"
	"net/http"
	"parkinsons/routes"
)

func main() {
	router := routes.NewRouter()

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
