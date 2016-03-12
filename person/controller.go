package person

import (
	"net/http"
	"strings"
	"github.com/pborges/mvc/common"
	log "github.com/Sirupsen/logrus"
)

const TemplatePrefix string = "tmpl/person/"

func RegisterController() (c *Controller) {
	log.WithField("controller", "person").Info("register controller")
	c = new(Controller)
	c.db = make([]*Person, 2)

	c.db[0] = new(Person)
	c.db[0].FirstName = "jim"
	c.db[0].LastName = "jones"
	c.db[0].Notes = make([]string, 0)

	c.db[1] = new(Person)
	c.db[1].FirstName = "bob"
	c.db[1].LastName = "dole"
	c.db[1].Notes = make([]string, 0)

	return c
}

type Controller struct {
	common.Controller
	db []*Person
}

func (this *Controller)List(w http.ResponseWriter, r *http.Request) {
	model := common.CreateViewModel()
	model.Model = this.db

	this.ViewLayout(model, TemplatePrefix + "list.tmpl.html")(w, r)
}

func (this *Controller)Show(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}

	model := common.CreateViewModel()
	model.Model = this.db[id]

	this.ViewLayout(model,
		TemplatePrefix + "show.tmpl.html",
		TemplatePrefix + "person.show.tmpl.html",
	)(w, r)
}

func (this *Controller)Edit(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	p := new(Person)
	*p = *this.db[id]

	model := common.CreateViewModel()
	model.ViewBag["prefix"] = "p1"
	model.Model = p

	if r.Method == "POST" {
		model.Errors = p.ParseForm(model.Prefix, r)
		if len(model.Errors) == 0 {
			// run form actions
			args := strings.Split(r.FormValue("_action"), ".")
			if len(args) > 0 {
				switch args[0]{
				case "Save":
					if errs := p.Validate(); len(errs) == 0 {
						this.db[id] = p
						http.Redirect(w, r, "/person", 302)
					}else {
						model.Errors = append(model.Errors, errs...)
					}
				case "Cancel":
					http.Redirect(w, r, "/person", 302)
				default:
					log.WithField("args", args).WithField("model", "person").Warn("unknown action")
				}
			}
		}
		for _, e := range model.Errors {
			log.WithField("id", id).WithField("model", "person").WithError(e).Warn("model has errors")
		}
	}
	this.ViewLayout(model,
		TemplatePrefix + "edit.tmpl.html",
		TemplatePrefix + "person.edit.tmpl.html",
	)(w, r)
}

func (this *Controller)Delete(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	this.db = append(this.db[:id], this.db[id + 1:]...)
	http.Redirect(w, r, "/person", 302)
}