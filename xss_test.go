package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Auto Esacape

func templateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post", map[string]interface{}{
		"Title":   "Golang Auto Escape",
		"<p>Body": "Belajar Golang XSS</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateAutoEscape(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Disable Auto Escape
func templateAutoEscapeDisable(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post", map[string]interface{}{
		"Title": "Disable Auto Escape",
		// Disable
		"Body": template.HTML("<p>Belajar Golang XSS</p>"),
	})
}

func TestTemplateAutoEscapeDisable(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateAutoEscapeDisable(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// XSS
func templateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post", map[string]interface{}{
		"Name": "XSS",
		"Body": template.HTML(r.URL.Query().Get("body")),
	})
}

func TestXss(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateXSS(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
