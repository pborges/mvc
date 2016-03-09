package person

import (
	"net/http"
	"strconv"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Notes     []string
}

func (this *Person) AddNote() {
	this.Notes = append(this.Notes, "")
}

func (this *Person) LeroyJenkins() {
	this.FirstName = "Leeerrooyy"
	this.LastName = "Jennnkinssssss"
}

func (this *Person) ParseForm(r *http.Request) (err error) {
	r.ParseForm()
	this.FirstName = r.FormValue("FirstName")
	this.LastName = r.FormValue("LastName")
	this.Age, err = strconv.Atoi(r.FormValue("Age"))
	if err != nil {
		return
	}
	notes := r.Form["Notes[]"]
	this.Notes = make([]string, len(notes))
	for i, note := range notes {
		this.Notes[i] = note
	}
	return
}

func (this *Person) ProcessActions(r *http.Request) (err error) {
	switch r.FormValue("_action") {
	case "addnote":
		this.AddNote()
	case "leroyjenkins":
		this.LeroyJenkins()
	}
	return
}