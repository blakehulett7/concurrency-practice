package main

import "net/http"

func (app *Bridge) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Set-Cookie", "token=; Max-Age=0;")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
