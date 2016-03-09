package person

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"strconv"
)

func RegisterController() (c *Controller) {
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
	db []*Person
}

func (this *Controller)List(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("list.tmpl.html").ParseFiles("person/list.tmpl.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, this.db)
}

func (this *Controller)Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	t, err := template.New("show.tmpl.html").ParseFiles("person/show.tmpl.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, this.db[id])
}

func (this *Controller)Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	p := this.db[id]
	if r.Method == "POST" {
		err = p.ParseForm(r)
		if err != nil {
			panic(err)
		}
		switch r.FormValue("_action"){
		case "save":
			this.db[id] = p
			http.Redirect(w, r, "/person", 302)
		default:
			err = p.ProcessActions(r)
			if err != nil {
				panic(err)
			}
		}
	}
	t, err := template.New("edit.tmpl.html").ParseFiles("person/edit.tmpl.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, p)
}

func (this *Controller)Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	this.db = append(this.db[:id], this.db[id + 1:]...)
	http.Redirect(w, r, "/person", 302)
}