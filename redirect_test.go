package main

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Redirect")
}

func redirectFrom(w http.ResponseWriter, r *http.Request) {
	// logic
	http.Redirect(w, r, "/redirect-youtube", http.StatusTemporaryRedirect)
}

func redirectYoutube(w http.ResponseWriter, r *http.Request) {
	// logic
	http.Redirect(w, r, "https://www.youtube.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-to", redirectTo)
	mux.HandleFunc("/redirect-from", redirectFrom)
	mux.HandleFunc("/redirect-youtube", redirectYoutube)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	// req := httptest.NewRequest("GET", "localhost:1417", nil)
	// rec := httptest.NewRecorder()

	//redirectFrom(rec, req)
	// resp := rec.Result()
	// body, _ := io.ReadAll(resp.Body)

	// fmt.Println(string(body))
}
