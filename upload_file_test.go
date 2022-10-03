package main

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func uploadForm(w http.ResponseWriter, r *http.Request) {
	var (
		name = "upload_form.html"
	)
	myTemplates.ExecuteTemplate(w, name, nil)
}

func TestUploadFormServer(t *testing.T) {
	var (
		pattern                  = "/upload-form"
		handler http.HandlerFunc = uploadForm
	)
	mux := http.NewServeMux()
	mux.HandleFunc(pattern, handler)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func upload(w http.ResponseWriter, r *http.Request) {
	// Ambil Form File
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	// Membuat File Destinasi
	destination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// Masukkan File UPload ke Direktori Destinasi
	_, err = io.Copy(destination, file)
	if err != nil {
		panic(err)
	}
	// Ambil Form Name
	name := r.PostFormValue("name")
	// Execute Template
	myTemplates.ExecuteTemplate(w, "upload_success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadSuccessServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadForm)
	mux.HandleFunc("/upload", upload)
	// static
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
