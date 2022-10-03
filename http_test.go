package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHTTP(t *testing.T) {
	// Arrange
	request := httptest.NewRequest(http.MethodGet, "localhost:1417", nil)
	recorder := httptest.NewRecorder()
	// Act
	// ServeHTTP
	Handler().ServeHTTP(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)
	// Assert
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Hello", string(body))
}
