package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	// Create Query Param
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name not found")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "localhost:1417/?name=akbar", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respBody))

	assert := assert.New(t)
	assert.Equal(200, resp.StatusCode)
	assert.Equal("Hello akbar", string(respBody))
}

func TestResponseCodeFail(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "localhost:1417/?name=", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	resp := recorder.Result()
	respBody, _ := io.ReadAll(resp.Body)

	assert := assert.New(t)
	assert.Equal(400, resp.StatusCode)
	assert.Equal("name not found", string(respBody))
}
