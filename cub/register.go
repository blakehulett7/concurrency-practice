package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var UrlSigningKey = []byte("83c71d87-36dd-4c08-866e-048e2e3cd611")

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

	email := r.FormValue("email")
	password := r.FormValue("password")
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(w, "Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := User{
		Email:      email,
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

	url := fmt.Sprintf("http://localhost:1000/activate?email=%s", email)
	signer := hmac.New(sha256.New, UrlSigningKey)

	signer.Write([]byte(url))
	signedUrl := signer.Sum(nil)
	encoded := hex.EncodeToString(signedUrl)

	msg := fmt.Sprintf("Please click the following link to activate your account http://localhost:1000/activate?hash=%s", encoded)

	app.SendEmail(email, "Activate your account", msg)
}
