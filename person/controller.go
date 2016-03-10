package person

import (
	"net/http"
	"html/template"
	"strconv"
	"strings"
	"github.com/pborges/rest/common"
)

func RegisterController() (c *Controller) {
	c = new(Controller)
	c.TemplatePrefix = "people"
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
	db             []*Person
	TemplatePrefix string
}

func (this *Controller)List(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("list.tmpl.html").ParseFiles(this.TemplatePrefix + "/list.tmpl.html")
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	t.Execute(w, this.db)
}

func (this *Controller)Show(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}

	t, err := template.New("show.tmpl.html").ParseFiles(this.TemplatePrefix + "/show.tmpl.html")
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	t.Execute(w, this.db[id])
}

func (this *Controller)Edit(w http.ResponseWriter, r *http.Request) {
	id, err := this.GetIdFromURL(w, r)
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	p := new(Person)
	*p = *this.db[id]
	if r.Method == "POST" {
		ParseForm(r, p)
		if len(p.Errors()) == 0 {
			args := strings.Split(r.FormValue("_action"), ".")
			switch args[0]{
			case "save":
				this.db[id] = p
				http.Redirect(w, r, "/person", 302)
			case "addnote":
				p.AddNote()
			case "removenote":
				if len(args) == 2 {
					var i int
					i, err = strconv.Atoi(args[1])
					if err != nil {
						return
					}
					p.RemoveNote(i)
				}
			case "leroyjenkins":
				p.LeroyJenkins()
			}
		}
	}
	t, err := template.New("edit.tmpl.html").ParseFiles("person/edit.tmpl.html")
	if err != nil {
		this.Error(err)(w, r)
		return
	}
	t.Execute(w, p)
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