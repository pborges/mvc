package view

import "html/template"

var templateCache map[[32]byte]*template.Template
var CacheTemplates = true

func init() {
	templateCache = make(map[[32]byte]*template.Template)
}
