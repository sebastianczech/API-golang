package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/sebastianczech/API-golang/imdb"
	"github.com/sebastianczech/API-golang/lubimyczytac"
)

func TestSearchFilm(t *testing.T) {
	router := httprouter.New()
	router.GET("/films", searchFilm)

	req, _ := http.NewRequest("GET", "/films?find=Titanic", nil)
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

	data := []*imdb.IdmbFilm{}
	json.Unmarshal([]byte(body), &data)
	fmt.Printf("Films: %s", data)

	if len(data) < 1 {
		t.Error("Not found any film")
	}

	if len(data[0].Title) == 0 {
		t.Error("Empty title")
	}
	if len(data[0].Image) == 0 {
		t.Error("Empty image URL")
	}
	if len(data[0].Website) == 0 {
		t.Error("Empty film URL")
	}
}

func TestSearchBook(t *testing.T) {
	router := httprouter.New()
	router.GET("/books", searchBook)

	req, _ := http.NewRequest("GET", "/books?find=człowiek nietoperz", nil)
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

	data := []*lubimyczytac.LubimyCzytacBook{}
	json.Unmarshal([]byte(body), &data)
	fmt.Printf("Books: %s", data)

	if len(data) < 1 {
		t.Error("Not found any book")
	}

	if len(data[0].Title) == 0 {
		t.Error("Empty title")
	}
	if len(data[0].Author) == 0 {
		t.Error("Empty author")
	}
	if len(data[0].Image) == 0 {
		t.Error("Empty image URL")
	}
	if len(data[0].Website) == 0 {
		t.Error("Empty book URL")
	}
}

func TestSearchBookWithWrongUrl(t *testing.T) {
	router := httprouter.New()
	router.GET("/books", searchBook)

	req, _ := http.NewRequest("GET", "/books?findWrong=człowiek nietoperz", nil)
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

	data := []*lubimyczytac.LubimyCzytacBook{}
	json.Unmarshal([]byte(body), &data)
	fmt.Printf("Books: %s\n", data)

	if len(data) > 0 {
		t.Error("Not found any book")
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

	fmt.Printf("Info: %s\n", string(body))
}

func TestMetricsLink(t *testing.T) {
	router := httprouter.New()
	router.GET("/metrics", metricsLink)

	req, _ := http.NewRequest("GET", "/metrics", nil)
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

	fmt.Printf("Metrics: %s\n", string(body))
}
