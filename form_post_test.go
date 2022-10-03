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

// PostForm Like Query Param
// PostFrom Send Param in Body not URL
func FormPost(w http.ResponseWriter, r *http.Request) {
	// Parsing first Before exec PostFrom
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := r.PostForm.Get("first_name")
	lastname := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	assert := assert.New(t)

	requestBody := strings.NewReader("first_name=jalal&last_name=akbar")
	request := httptest.NewRequest(http.MethodPost, "localhost:1417/", requestBody)
	// "application/x-www-form-urlencoded" standard to exec FormPost
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))
	assert.NotNil(string(respBody))
}
