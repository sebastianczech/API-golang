package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestSearchBook(t *testing.T) {
	router := httprouter.New()
	router.GET("/books", searchBook)

	req, _ := http.NewRequest("GET", "/books?find=człowiek+nieto", nil)
	rep := httptest.NewRecorder()

	router.ServeHTTP(rep, req)

	resp := rep.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(string(body))

	if status := rep.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}

func TestHomeLink(t *testing.T) {
	router := httprouter.New()
	router.GET("/", homeLink)

	req, _ := http.NewRequest("GET", "/", nil)
	rep := httptest.NewRecorder()

	router.ServeHTTP(rep, req)

	resp := rep.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(string(body))

	if status := rep.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
