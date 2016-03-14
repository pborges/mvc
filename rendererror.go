package mvc

import (
	"net/http"
	"github.com/pborges/log"
)

func RenderError(err error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.WithError(err).Error("error route")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
