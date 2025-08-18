package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *Bridge) Login(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Flash:         "",
		Warning:       "",
		Error:         "",
		Authenticated: false,
	}
	Render(w, "login.html", data)
}

func (app *Bridge) PostLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	user := User{}
	result := app.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		fmt.Fprintf(w, "email not found")
		return
	}

	password := r.FormValue("password")
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Fprintf(w, "incorrect password")
		return
	}

	session_token_cookie := fmt.Sprintf("token=%dIsAuthenticated", user.Id)
	w.Header().Set("Set-Cookie", session_token_cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
