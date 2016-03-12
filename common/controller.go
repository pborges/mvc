package common

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/pborges/log"
)

type Controller struct {
}

func (this *Controller)Error(err error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithError(err).Error("error route")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (this *Controller)GetIdFromURL(w http.ResponseWriter, r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	id, err = strconv.Atoi(vars["id"])
	return
}
