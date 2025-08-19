package main

import (
	"fmt"
	"net/http"
)

func (app *Bridge) Activate(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")

	email, err := UnhashEmail(hash)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Bad hash")
		return
	}

	fmt.Fprintf(w, "success, %s", email)
}
