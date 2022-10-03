package main

import (
	"fmt"
	"net/http"
)

// Request Adalah struct yang merepresentasikan http.Request yg dikirim oleh Web Browser
func Request() http.Handler {
	mux := http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Host)
		fmt.Fprintln(w, r.Proto)
	})
	return &mux
}
