package main

import (
	"fmt"
	"net/http"
)

// Serve Mux adalah Implementasi Handler yang bisa mendukung multiple endpoint
func ServeMux() http.Handler {
	mux := http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Who are you?")
	})
	mux.HandleFunc("/sonia", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Sonia")
	})
	mux.HandleFunc("/jazil", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello jazil")
	})
	// URL Pattern
	mux.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello ..??")
	})
	mux.HandleFunc("/hello/sonia/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Sonia")
	})
	return &mux
}
