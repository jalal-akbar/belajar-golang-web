package main

import (
	"fmt"
	"net/http"
)

// Handler adalah penerima HTTP Request yang masuk Ke Server
func Handler() http.Handler {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}
	return handler
}
