package home

import (
	"net/http"
	"github.com/pborges/mvc/common"
	log "github.com/Sirupsen/logrus"
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
	model := common.CreateViewModel()
	this.ViewLayout(model, TemplatePrefix + "index.tmpl.html")(w, r)
}
