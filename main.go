// Re-creating backend for Parkinson's App in Go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"parkinsons/models"
	"path"
	"strconv"
	"strings"
)

func ConvertToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func CheckForIDInRoute(route string) (bool, int) {
	i, err := strconv.Atoi(route)
	if err != nil {
		return false, 0
	}
	return true, i
}

func routes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	route := r.URL.Path
	method := r.Method

	switch true {
	case route == "/":
		fmt.Fprintf(w, "Welcome to Parkinson's API built in GO")
	case strings.HasPrefix(route, "/api/v1/patients") && method == "GET":
		found, id := CheckForIDInRoute(path.Base(r.URL.Path))
		var response []byte
		if !found {
			response = ConvertToJSON(patients.GetAllPatients())
		} else {
			patient := patients.GetOnePatient(id)
			if patient.Id == 0 {
				response = ConvertToJSON("Patient not found")
			} else {
				response = ConvertToJSON(patients.GetOnePatient(id))
			}
		}
		w.Write(response)
	case strings.HasPrefix(route, "/api/v1/users"):
		response := ConvertToJSON("Get all users")
		w.Write(response)
	case strings.HasPrefix(route, "/api/v1/providers"):
		response := ConvertToJSON("Get all users")
		w.Write(response)
	default:
		fmt.Fprintf(w, "Path not found")
	}
}

func main() {
	http.HandleFunc("/", routes)
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
