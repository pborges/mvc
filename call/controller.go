package call

import (
	"net/http"
	"strings"
	"time"
	"github.com/pborges/mvc/common"
	"github.com/pborges/mvc/person"
	"github.com/pborges/log"
	"github.com/pborges/mvc/view"
)

const TemplatePrefix string = "tmpl/call/"

func RegisterController() (c *Controller) {
	log.WithField("controller", "call").Info("register controller")
	c = new(Controller)
	c.db = make([]*Call, 1)

	c.db[0] = new(Call)
	c.db[0].Caller.FirstName = "jim"
	c.db[0].Caller.LastName = "jones"
	c.db[0].Caller.Notes = make([]string, 0)
	c.db[0].ReasonForCalling = "Stick in a vice"
	n := time.Now()
	c.db[0].CreatedOn = &n
	c.db[0].LastModified = &n
	c.db[0].PhoneNumber = 2096750475
	return c
}

type Controller struct {
	common.Controller
	db []*Call
}

func (this *Controller)List(w http.ResponseWriter, r *http.Request) {
	model := view.NewViewModel()
	model.Model = this.db

	view.Layout(model, TemplatePrefix + "list.tmpl.html")(w, r)
}

func (this *Controller)Show(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}

	model := view.NewViewModel()
	model.Model = this.db[id]

	view.Layout(model,
		TemplatePrefix + "show.tmpl.html",
		TemplatePrefix + "call.show.tmpl.html",
		person.TemplatePrefix + "person.show.tmpl.html",
	)(w, r)
}

func (this *Controller)Edit(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	p := new(Call)
	*p = *this.db[id]
	model := view.NewViewModel()
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
						http.Redirect(w, r, "/call", 302)
					}else {
						model.Errors = append(model.Errors, errs...)
					}
				case "Cancel":
					http.Redirect(w, r, "/call", 302)
				default:
					log.WithField("args", args).WithField("model", "call").Warn("unknown action")
				}
			}
		}
		for _, e := range model.Errors {
			log.WithField("id", id).WithField("model", "call").WithError(e).Warn("model has errors")
		}
	}
	view.Layout(model,
		TemplatePrefix + "edit.tmpl.html",
		TemplatePrefix + "call.edit.tmpl.html",
		person.TemplatePrefix + "person.edit.tmpl.html",
	)(w, r)
}

func (this *Controller)Delete(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	this.db = append(this.db[:id], this.db[id + 1:]...)
	http.Redirect(w, r, "/call", 302)
}