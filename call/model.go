package call

import (
	"net/http"
	"strings"
	"github.com/pborges/mvc/person"
	"time"
)

type Call struct {
	Caller           person.Person
	ReasonForCalling string
	PhoneNumber      int
	CreatedOn        *time.Time
	LastModified     *time.Time
}

func (this *Call) ParseForm(prefix string, r *http.Request) (errs []error) {
	errs = make([]error, 0)
	err := r.ParseForm()
	if err != nil {
		errs = append(errs, err)
		return
	}
	// Parse Form
	errs = append(errs, this.Caller.ParseForm("caller", r)...)
	this.ReasonForCalling = r.FormValue(prefix + ".ReasonForCalling")

	// Process Actions
	args := strings.Split(r.FormValue(prefix + "._action"), ".")
	switch args[0]{

	}
	return
}

func (this *Call)Validate() (errs []error) {
	errs = this.Caller.Validate()
	return
}