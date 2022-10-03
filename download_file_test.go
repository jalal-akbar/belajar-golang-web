package main

import (
	"fmt"
	"net/http"
	"testing"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	// logic
	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}
	// Render File
	//http.ServeFile(w, r, "./resources/"+file)

	// Without Render / Render Paksa
	// \"\"
	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./resources/"+file)

}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(downloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
