package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

// Template Function
type MyPgae struct {
	Name string
}

func (m MyPgae) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + m.Name
}
func templateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/mypage.html"))

	t.ExecuteTemplate(w, "mypage", MyPgae{
		Name: "Akbar",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()

	templateFunction(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Global Function
func templateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("glob").ParseFiles("./templates/global_func.html"))

	t.ExecuteTemplate(w, "glob", map[string]interface{}{
		"Name": "Jalaluddin Muh Akbar",
	})
}

func TestTemplateGlob(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateFunctionGlobal(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Menambah Global Function

func templateAddFunction(w http.ResponseWriter, r *http.Request) {
	parse := `{{upper .Name}}`
	data := MyPgae{Name: "akbar"}
	t := template.New("addFunc")
	// adding func
	t.Funcs(map[string]interface{}{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
	})
	t = template.Must(t.Parse(parse))

	t.ExecuteTemplate(w, "addFunc", data)
}

func TestTemplateAddFunc(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateAddFunction(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Pipelines
func templateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("pipelines")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(upper string) string {
			return strings.ToUpper(upper)
		},
	})
	var parse = `{{sayHello .Name | upper}}`
	t = template.Must(t.Parse(parse))

	t.ExecuteTemplate(w, "pipelines", MyPgae{Name: "Akbar"})
}

func TestTemplatePipelines(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:1417", nil)
	rec := httptest.NewRecorder()

	templateFunctionPipelines(rec, req)
	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
