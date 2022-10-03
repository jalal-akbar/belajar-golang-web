package main

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Logic
	// Before Execute Handler
	fmt.Println("Before Execute")
	// Init
	middleware.Handler.ServeHTTP(w, r)
	// Logic
	// After Execute Handler
	fmt.Println("After Execute")
}

func TestMiddleware(t *testing.T) {
	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success Execute Middleware")
		fmt.Fprintln(w, "This is Middleware")
	})
	// Set Up Middlerware
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

// Error Handler
type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Logic Before Execute
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success Execute Middleware")
		fmt.Fprint(w, "Hello Middlware")
	})
	// Error Handler
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("UPS")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
