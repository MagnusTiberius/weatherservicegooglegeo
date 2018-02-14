package main

import (
	"net/http"

	"github.com/MagnusTiberius/weatherservicegooglegeo/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.GoogleGeoHandler)
	r.HandleFunc("/weather/{addr}", handler.GoogleGeoHandler)
	http.ListenAndServe(":8095", r)
}
