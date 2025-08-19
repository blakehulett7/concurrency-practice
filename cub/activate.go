package main

import (
	"fmt"
	"net/http"
)

func (app *Bridge) Activate(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	email, err := UnhashEmail(hash)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Bad hash")
		return
	}

	user := User{}
	result := app.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		panic("should never get here... correct hash but no corresponding email")
	}

	user.UserActive = 1
	result = app.DB.Save(&user)
	if result.Error != nil {
		panic("should never get here... found user but could not update")
	}

	session_token_cookie := fmt.Sprintf("token=%dIsAuthenticated", user.Id)
	w.Header().Set("Set-Cookie", session_token_cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
