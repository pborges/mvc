package mvc

import (
	"html/template"
)

var templateCache map[[32]byte]*template.Template
var CacheTemplates = true
var Config *Configuration

func init() {
	templateCache = make(map[[32]byte]*template.Template)
	Config = newConfiguration()
	Config.SiteName = "mvc site"
}
