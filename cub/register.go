package main

import "net/http"

func (app *Bridge) Register(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: false,
	}
	Render(w, "register.html", data)
}

func (app *Bridge) PostRegister(w http.ResponseWriter, r *http.Request) {
	// Do some stuff

	app.WaitGroup.Add(1)
	go SendEmail("to", "sub", "body", app.EmailErrChannel, app.WaitGroup)
}
