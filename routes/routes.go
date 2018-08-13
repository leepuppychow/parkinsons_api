package routes

import (
	"net/http"
	"parkinsons/controllers"

	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{ "PatientsIndex", "GET", "/api/v1/patients", patients.Index },
	{ "PatientsShow", "GET", "/api/v1/patients/{id}", patients.Show },
	{ "PatientsCreate", "POST", "/api/v1/patients", patients.Create },
	{ "PatientsUpdate", "PATCH", "/api/v1/patients/{id}", patients.Update },
	{ "PatientsDelete", "DELETE", "/api/v1/patients/{id}", patients.Delete },
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

