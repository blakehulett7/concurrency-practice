package main

import (
	"fmt"
	"net/http"
)

func RecoveryMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("Recovered from panic...")
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}
}
