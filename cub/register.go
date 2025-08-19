package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(w, "Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := User{
		Email:      r.FormValue("email"),
		FirstName:  r.FormValue("first-name"),
		LastName:   r.FormValue("last-name"),
		Password:   string(hashed_password),
		UserActive: 0,
		IsAdmin:    0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		PlanId:     4,
	}
	result := app.DB.Create(&user)
	if result.Error != nil {
		fmt.Fprintf(w, "Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.SendEmail("to", "sub", "body")
}
