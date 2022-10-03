package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Query Parameter digunukan untuk mengirim data dari client ke server
// Query Parameter di simpan dalam URL
// Dari URL ini kita dapat mengambil data query parameter yang dikirim dari client dengan method Query() yang return map

// Single Query Parameter
func SingleQueryParameter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, "Hello Who Are You?")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}
func TestSingleQueryParameter(t *testing.T) {
	// Arrange
	request := httptest.NewRequest(http.MethodGet, "localhost:1417/?name=Sonia", nil) // ?name=Sonia is Parameter
	recorder := httptest.NewRecorder()
	// Act
	SingleQueryParameter(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello Sonia", string(body))
}

// Multiple Query Parameter
func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("first_name")
	lastname := r.URL.Query().Get("last_name")

	if firstname != "" {
		fmt.Fprintf(w, "%s \n", firstname)
	} else {
		fmt.Fprintln(w, "first name required")
	}
	if lastname != "" {
		fmt.Fprintf(w, "%s \n", lastname)
	} else {
		fmt.Fprintln(w, "last name required")
	}

}
func TestMultipleQueryParameter(t *testing.T) {
	// Arrange
	request := httptest.NewRequest(http.MethodGet, "localhost:1417/?first_name=Jalal&last_name=Akbar", nil)
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	recorder := httptest.NewRecorder()
	// Act
	MultipleQueryParameter(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	// Assert
	fmt.Println(string(body))
	fmt.Println(first_name)
	fmt.Println(last_name)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Jalal", first_name)
	assert.Equal(t, "Akbar", last_name)
}

// Multiple Value Query Paramater
func MultipleValuesQueryParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	// fmt.Fprintln(w, name[0]+name[1])
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleValuesQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/?name=jalal&name=akbar", nil)
	request.URL.Query().Get("name")
	recorder := httptest.NewRecorder()

	MultipleValuesQueryParameter(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
	assert.Equal(t, "jalal akbar", string(body))
}
