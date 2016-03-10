package common

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Controller struct {
}

func (this *Controller)Error(err error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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