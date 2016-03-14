package mvc

import (
	"net/http"
	"html/template"
	"github.com/pborges/log"
	"strings"
	"crypto/sha256"
	"encoding/hex"
)

var SharedTemplatePrefix string = "tmpl/shared/"

func RenderLayout(model *ViewModel, files... string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		files = append(files, SharedTemplatePrefix + "master.tmpl.html")
		var err error
		sha := sha256.Sum256([]byte(strings.Join(files, ".")))
		t, ok := templateCache[sha]
		if !CacheTemplates || !ok {
			log.WithField("files", files).WithField("sha", hex.EncodeToString(sha[:])).Debug("generating template")
			t, err = template.New("master.tmpl.html").Funcs(FuncMap).ParseFiles(files...)
			if err != nil {
				log.WithError(err).WithField("files", files).WithField("model", model).Panic("error parsing template")
			}
			templateCache[sha] = t
		}
		err = t.Execute(w, model)
		if err != nil {
			log.WithError(err).WithField("files", files).WithField("model", model).Panic("error executing template")
		}
	}
}
