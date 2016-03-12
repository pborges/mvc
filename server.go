package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/pborges/mvc/home"
	"github.com/pborges/mvc/call"
	"github.com/pborges/mvc/person"
	"github.com/pborges/log"
	"github.com/pborges/mvc/view"
)

const ListenAddress string = ":8080"

func main() {
	log.LogLevel = log.LevelDebug
	view.CacheTemplates = false
	log.Info("starting")
	r := mux.NewRouter()

	homeController := home.RegisterController()
	r.HandleFunc("/", homeController.Index).Methods("GET")

	personController := person.RegisterController()
	r.HandleFunc("/person", personController.List).Methods("GET")
	r.HandleFunc("/person/{id:[0-9]+}", personController.Show).Methods("GET")
	r.HandleFunc("/person/{id:[0-9]+}/edit", personController.Edit).Methods("GET", "POST")
	r.HandleFunc("/person/{id:[0-9]+}/delete", personController.Delete).Methods("GET")

	callController := call.RegisterController()
	r.HandleFunc("/call", callController.List).Methods("GET")
	r.HandleFunc("/call/{id:[0-9]+}", callController.Show).Methods("GET")
	r.HandleFunc("/call/{id:[0-9]+}/edit", callController.Edit).Methods("GET", "POST")
	r.HandleFunc("/call/{id:[0-9]+}/delete", callController.Delete).Methods("GET")

	// static
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Warn("listening on", ListenAddress)
	log.Panic(http.ListenAndServe(ListenAddress, Log(r)))
}

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField("ip", r.RemoteAddr).WithField("url", r.URL).WithField("method", r.Method).Info("http request")
		h.ServeHTTP(w, r)
	})
}