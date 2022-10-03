package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Layout
func templateLayout(w http.ResponseWriter, r *http.Request) {
	var (
		header = "./templates/header.html"
		footer = "./templates/footer.html"
		layout = "./templates/layout.html"
		data   = map[string]interface{}{
			"Name":  "Akbar",
			"Title": "This is Layout",
		}
	)
	t := template.Must(template.ParseFiles(header, layout, footer))
	t.ExecuteTemplate(w, "layout.html", data)
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	templateLayout(recorder, request)
	resp := recorder.Result()
	bodyResp, _ := io.ReadAll(resp.Body)

	fmt.Println(string(bodyResp))
}

// Define name
func templateLayoutWithDefineName(w http.ResponseWriter, r *http.Request) {
	var (
		header = "./templates/define_header.html"
		footer = "./templates/define_footer.html"
		layout = "./templates/define_layout.html"
		data   = map[string]interface{}{"Name": "Akbar", "Title": "Template Define"}
	)
	t := template.Must(template.ParseFiles(header, footer, layout))

	//t := template.Must(template.ParseFiles(header, footer, layout))
	t.ExecuteTemplate(w, "layout", data)
}

func TestTempplateWithDefineName(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	templateLayoutWithDefineName(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
