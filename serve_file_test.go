package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "/resources/notFound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:1417",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Serve File With Embed
//
//go:embed resources/ok.html
var resourcesOK string

//
//go:embed resources/notFound.html
var resourcesNotFound string

func ServeFileWithEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		fmt.Fprint(w, resourcesOK)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}
}

func TestServeFileWithEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:1417",
		Handler: http.HandlerFunc(ServeFileWithEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
