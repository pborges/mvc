package person

import (
	"net/http"
	"strconv"
	"strings"
	"errors"
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

func (this *Person) RemoveNote(id int) {
	this.Notes = append(this.Notes[:id], this.Notes[id + 1:]...)
}

func (this *Person) LeroyJenkins() {
	this.FirstName = "Leeerrooyy"
	this.LastName = "Jennnkinssssss"
}

func (this *Person)Validate() (errs []error) {
	errs = make([]error, 0)
	if this.Age <= 0 || this.Age >= 120 {
		errs = append(errs, errors.New("Invalid Age!"))
	}
	if this.FirstName == "bob" {
		errs = append(errs, errors.New("Thats a dumb first name"))
	}
	return
}

func (this *Person) ParseForm(prefix string, r *http.Request) (errs []error) {
	errs = make([]error, 0)
	err := r.ParseForm()
	if err != nil {
		errs = append(errs, err)
		return
	}

	// Parse Form
	this.FirstName = r.FormValue(prefix + ".FirstName")
	this.LastName = r.FormValue(prefix + ".LastName")
	this.Age, err = strconv.Atoi(r.FormValue(prefix + ".Age"))
	if err != nil {
		errs = append(errs, err)
	}
	notes := r.Form[prefix + ".Notes[]"]
	this.Notes = make([]string, len(notes))
	for i, note := range notes {
		this.Notes[i] = note
	}

	// Process Actions
	args := strings.Split(r.FormValue(prefix + "._action"), ".")
	switch args[0]{
	case "AddNote":
		this.AddNote()
	case "RemoveNote":
		if len(args) == 2 {
			var i int
			i, err = strconv.Atoi(args[1])
			if err != nil {
				errs = append(errs, err)
			}else {
				this.RemoveNote(i)
			}
		}
	case "LeroyJenkins":
		this.LeroyJenkins()
	}
	return
}