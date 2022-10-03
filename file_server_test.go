package main

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// With Directory
func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")
	fileServer := http.FileServer(dir)
	// File Server read Pattern
	// /resources/file(index.html)

	mux := http.NewServeMux()
	// 404 not Found
	// StripPrefix
	// Delete /static
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:1417",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// File Server With Golang Embed
//
//go:embed resources
var resources embed.FS

func TestFileServerWithEmbed(t *testing.T) {

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/akbar/", http.StripPrefix("/akbar", fileServer))

	server := http.Server{
		Addr:    "localhost:1417",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
