package mvc

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)
// GetIdFromRequest only exists to facilitate an easy change if I decide to move away from gorilla mux
func GetIdFromRequest(w http.ResponseWriter, r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	id, err = strconv.Atoi(vars["id"])
	return
}