package person

import (
	"net/http"
	"strconv"
)

func ParseForm(r *http.Request, p *Person) {
	p.errors = make([]error, 0)
	err := r.ParseForm()
	if err != nil {
		p.errors = append(p.errors, err)
		return
	}
	p.FirstName = r.FormValue("FirstName")
	p.LastName = r.FormValue("LastName")
	p.Age, err = strconv.Atoi(r.FormValue("Age"))
	if err != nil {
		p.errors = append(p.errors, err)
	}
	notes := r.Form["Notes[]"]
	p.Notes = make([]string, len(notes))
	for i, note := range notes {
		p.Notes[i] = note
	}
	return
}

