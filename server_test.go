package main

import (
	"net/http"
	"testing"
)

// Test Server
func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:1417",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Test Handler
func TestHandler(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:1417",
		Handler: Handler(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Test ServeMux
func TestServeMux(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:1417",
		Handler: ServeMux(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Test Request
func TestRequest(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:1417",
		Handler: Request(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
