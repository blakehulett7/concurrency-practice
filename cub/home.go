package main

import (
	"net/http"
)

func (app *Bridge) Home(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: false,
	}
	Render(w, "home.html", data)
}
