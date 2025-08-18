package main

import (
	"net/http"
)

func (app *Bridge) Home(w http.ResponseWriter, r *http.Request) {
	is_authenticated := false
	_, err := r.Cookie("token")
	if err == nil {
		is_authenticated = true
	}

	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: is_authenticated,
	}
	Render(w, "home.html", data)
}
