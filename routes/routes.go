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
	{ "Index", "GET", "/api/v1/patients", patients.Index },
	{ "Show", "GET", "/api/v1/patients/{id}", patients.Show },
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

