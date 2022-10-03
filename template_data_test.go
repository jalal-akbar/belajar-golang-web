package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Template With Map
func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Ini Title",
		"Name":  "Akbar",
	})
}

func TestTemplateDataWithMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Template With Struct
type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	var (
		nameTemplate    = "name.gohtml"
		patternTemplate = "./templates/*.gohtml"
	)
	//template.ParseFiles(filenames)
	temp := template.Must(template.ParseGlob(patternTemplate))

	temp.ExecuteTemplate(w, nameTemplate, Page{
		Title: "This Title Template With Struct",
		Name:  "Akbar",
	})
}

func TestTemplateDataWithStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
