package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

// HTML Template String
func SimpleTemplate(w http.ResponseWriter, r *http.Request) {
	var (
		textTemplate        = "<html><body>{{.}}</body></html>"
		nameTemplate        = "SIMPLE"
		dinamisDataTemplate = "ini data dinamis nya"
	)
	t := template.New(nameTemplate)
	template := template.Must(t.Parse(textTemplate))

	template.ExecuteTemplate(w, nameTemplate, dinamisDataTemplate)
}
func TestSimpleHTML(t *testing.T) {
	var (
		method = http.MethodGet
		target = "localhost:1417"
	)
	request := httptest.NewRequest(method, target, nil)
	recorder := httptest.NewRecorder()

	SimpleTemplate(recorder, request)

	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "<html><body>ini data dinamis nya</body></html>", string(respBody))
	// server := http.Server{
	// 	Addr:    "localhost:1417",
	// 	Handler: http.HandlerFunc(SimpleTemplate),
	// }
	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }
}

// Template dari File
func FileTemplate(w http.ResponseWriter, r *http.Request) {
	var (
		nameTemplate = "template.html"
		textTemplate = "./resources/template.html"
		dataTemplate = "ini data nya"
	)
	t := template.Must(template.New(nameTemplate).ParseFiles(textTemplate))

	t.ExecuteTemplate(w, nameTemplate, dataTemplate)
}

func TestTemplateWithFile(t *testing.T) {
	var (
		method = http.MethodGet
		target = "localhost:1417"
	)
	request := httptest.NewRequest(method, target, nil)
	recorder := httptest.NewRecorder()

	FileTemplate(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))
	//assert.Equal(t, templateOK, string(respBody))
}

// Directory Template
func DirTemplate(w http.ResponseWriter, r *http.Request) {
	var (
		nameTemplate    = "template.html"
		patternTemplate = "./templates/*.html"
		dataTemplate    = "ini lagi data nya"
	)
	t := template.Must(template.ParseGlob(patternTemplate))

	t.ExecuteTemplate(w, nameTemplate, dataTemplate)
}

func TestTemplateWithDir(t *testing.T) {
	var (
		method = http.MethodGet
		target = "localhost:1417"
	)
	request := httptest.NewRequest(method, target, nil)
	recorder := httptest.NewRecorder()

	DirTemplate(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))
}

// Embed Template

func TemplateWithEmbed(w http.ResponseWriter, r *http.Request) {
	var (
		nameTemplate    = "template.html"
		patternTemplate = "templates/*.html"
		dataTemplate    = "ini data embed"
	)
	//t := template.Must(template.New(nameTemplate).ParseFS(templatess, patternTemplate))
	t, err := template.ParseFS(templates, patternTemplate)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, nameTemplate, dataTemplate)
}

func TestTemplateWithEmbed(t *testing.T) {
	var (
		method = http.MethodGet
		target = "localhost:1417"
	)
	request := httptest.NewRequest(method, target, nil)
	recorder := httptest.NewRecorder()

	TemplateWithEmbed(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	fmt.Print(string(respBody))
}
