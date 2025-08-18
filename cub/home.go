package main

import (
	"html/template"
	"net/http"
)

func (app *Bridge) Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("templates/*.gohtml"))

	t.ExecuteTemplate(w, "home.page.gohtml", nil)
}
