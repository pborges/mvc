package mvc

import (
	"encoding/json"
	"net/http"
	"github.com/pborges/log"
)

func RenderError(err error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		entry := log.SetCallDepth(1).WithError(err)
		RenderDetailedError(entry)(w, r)
		return
	}
}

func RenderDetailedError(err *log.Entry) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err.Error("Error executing request")
		w.WriteHeader(http.StatusInternalServerError)
		var thrownError error
		for i, e := range err.Keys {
			if e == "err" {
				thrownError = err.Values[i].(error)
			}
		}

		e := struct {
			Heading     string
			Error      string
			Package    string
			Filename   string
			LineNumber int
			Fields     map[string]string
		}{
			Heading:err.Msg,
			Error:thrownError.Error(),
			Package:err.Package,
			Filename:err.Filename,
			LineNumber:err.Line,
		}
		e.Fields = make(map[string]string)
		for i, k := range err.Keys {
			j, _ := json.MarshalIndent(err.Values[i], "", "  ")
			e.Fields[k] = string(j)
		}

		model := NewViewModel()
		model.Model = e
		RenderLayout(model, SharedTemplatePrefix + "error.tmpl.html")(w, r)
		return
	}
}