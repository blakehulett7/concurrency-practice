package main

import "net/http"

func (app *Bridge) Login(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: false,
	}
	Render(w, "login.html", data)
}
