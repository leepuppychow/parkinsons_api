package patients

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"parkinsons/models"

	"github.com/gorilla/mux"
)

func ToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write(ToJSON(patient.All()))
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	w.Write(ToJSON(patient.Find(id)))
}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	w.Write(ToJSON(patient.Update(id)))		 
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	w.Write(ToJSON(patient.Delete(id)))
}
