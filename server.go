package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/pborges/rest/person"
	"fmt"
)


func main() {
	r := mux.NewRouter()

	personController := person.RegisterController()

	r.HandleFunc("/person", personController.List).Methods("GET")
	r.HandleFunc("/person/{id:[0-9]+}", personController.Show).Methods("GET")
	r.HandleFunc("/person/{id:[0-9]+}/edit", personController.Edit).Methods("GET", "POST")
	r.HandleFunc("/person/{id:[0-9]+}/delete", personController.Delete).Methods("GET")

	panic(http.ListenAndServe(":8080", Log(r)))
}

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}