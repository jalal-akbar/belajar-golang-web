package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{"Name": ""})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// Comparator
func TemplateDataActionComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Value": 80,
	})
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()
	TemplateDataActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// Range
func TemplateActitonRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))
	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Hobbies": []string{"Gaming", "Reading", "Cooking"},
	})
}

func TestTemplatesActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	TemplateActitonRange(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// With
func templateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))
	t.ExecuteTemplate(w, "with.html", map[string]interface{}{
		"Name": "Akbar",
		"Address": map[string]interface{}{
			"Street": "jln yos sudarso",
			"City":   "Kota Bima",
		},
	})
}
func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	templateWith(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
