package view

import (
	"html/template"
	"time"
)

var FuncMap template.FuncMap = template.FuncMap{
	"prefix":Prefix,
	"formatdate":FormatDate,
}

func Prefix(prefix string, m interface{}) (o *ViewModel) {
	o = NewViewModel()
	o.Prefix = prefix
	o.Model = m
	return
}

func FormatDate(d *time.Time) (string) {
	return d.Format("01-02-2006 03:04 PM")
}