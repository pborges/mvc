package common

import (
	"net/http"
	"strconv"
	"html/template"
	"github.com/gorilla/mux"
	log "github.com/Sirupsen/logrus"
)

const TemplatePrefix string = "tmpl/common/"

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

func (this *Controller)ViewLayout(model *ViewModel, files... string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		files = append(files, TemplatePrefix + "/master.tmpl.html")
		t, err := template.New("master.tmpl.html").Funcs(FuncMap).ParseFiles(files...)
		if err != nil {
			this.Error(err)(w, r)
			return
		}
		err = t.Execute(w, model)
		if err != nil {
			panic(err)
		}
	}
}