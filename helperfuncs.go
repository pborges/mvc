package mvc

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"time"
	"fmt"
)
// GetIdFromRequest only exists to facilitate an easy change if I decide to move away from gorilla mux
func GetIdFromRequest(w http.ResponseWriter, r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	id, err = strconv.Atoi(vars["id"])
	return
}

func PrefixForm(prefix string, m interface{}) (o *Form) {
	o = new(Form)
	o.Prefix = prefix
	o.Model = m
	return
}

func FormatDate(d *time.Time) (string) {
	return d.Format("01-02-2006 03:04 PM")
}

func FormatMoney(d float64) (string) {
	return fmt.Sprintf("$%0.2f", d)
}