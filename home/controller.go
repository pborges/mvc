package home

import (
	"net/http"
	"github.com/pborges/log"
	"github.com/pborges/mvc/view"
	"github.com/pborges/mvc/common"
)

const TemplatePrefix string = "tmpl/home/"

func RegisterController() (c *Controller) {
	log.WithField("controller", "home").Info("register controller")
	c = new(Controller)
	return c
}

type Controller struct {
	common.Controller
}

func (this *Controller)Index(w http.ResponseWriter, r *http.Request) {
	model := view.NewViewModel()
	view.Layout(model, TemplatePrefix + "index.tmpl.html")(w, r)
}
