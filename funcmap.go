package mvc

import (
	"html/template"
)

var FuncMap template.FuncMap = template.FuncMap{
	"prefix":PrefixForm,
	"formatdate":FormatDate,
}