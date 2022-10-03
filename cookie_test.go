package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	// Init Cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-JMA-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	// Set Cookie
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-JMA-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}
func TestServerCookie(t *testing.T) {
	// Handling
	mux := new(http.ServeMux)
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:1417",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1417/?name=akbar", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)
	// Set on Recorder
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		//fmt.Printf("Name:%s \nValue:%s", cookie.Name, cookie.Value)
		assert.Equal(t, "X-JMA-Name", cookie.Name)
		assert.Equal(t, "akbar", cookie.Value)
		assert.Equal(t, "/", cookie.Path)
	}
}

func TestGetCookies(t *testing.T) {
	// cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-JMA-Name"
	cookie.Value = "akbar"
	cookie.Path = "/"

	request := httptest.NewRequest(http.MethodGet, "localhost:1417/", nil)
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)
	resp := recorder.Result()
	body, _ := io.ReadAll(resp.Body)

	//fmt.Print(string(body))
	assert.Equal(t, "Hello akbar", string(body))

}
