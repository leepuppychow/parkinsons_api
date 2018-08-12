// Re-creating backend for Parkinson's App in Go

package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"parkinsons/models"
)

func ConvertToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func routes(w http.ResponseWriter, r *http.Request) {
	switch (r.URL.Path) {
	case "/": 
		fmt.Fprintf(w, "Welcome to Parkinson's API built in GO")
	case "/api/v1/patients":
		patients := ConvertToJSON(patients.GetAllPatients())
		w.Header().Set("Content-Type", "application/json")
		w.Write(patients)
	case "/api/v1/users":
		fmt.Fprintf(w, "Show all users")
	case "/api/v1/providers":
		fmt.Fprintf(w, "Show all providers")
	default:
		fmt.Fprintf(w, "Path not found")
	}
}

func main() {
	http.HandleFunc("/", routes)
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

