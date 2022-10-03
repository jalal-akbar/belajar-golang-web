package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Request Header
func requestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Add("Content-Type", "application/json")

	requestHeader(recorder, request)
	resp := recorder.Result()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

// Response Header
func responseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-Key", "RAHASIA")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	responseHeader(recorder, request)
	resp := recorder.Result()
	respHeader := recorder.Header().Get("X-API-Key")
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
	assert.Equal(t, "OK", string(body))
	assert.Equal(t, "RAHASIA", respHeader)

}
