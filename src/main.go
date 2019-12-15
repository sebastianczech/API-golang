package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/sebastianczech/API-golang/imdb"
	"github.com/sebastianczech/API-golang/info"
	"github.com/sebastianczech/API-golang/lubimyczytac"
)

func searchFilm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	find := queryValues.Get("find")

	baseUrl, err := url.Parse("https://www.imdb.com")
	if err != nil {
		fmt.Printf("Error URL: %s", err.Error())
		return
	}
	baseUrl.Path += "find"
	params := url.Values{}
	params.Add("q", find)
	baseUrl.RawQuery = params.Encode()
	url := baseUrl.String()
	fmt.Printf("Encoded URL is %q\n", url)

	films := imdb.FindImdbFilm(url)
	fmt.Printf("API /films/%s: %s\n", find, films)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(films)
}

func searchBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	find := queryValues.Get("find")

	// url := fmt.Sprintf("http://lubimyczytac.pl/szukaj/ksiazki?phrase=%s&main_search=1", find)
	baseUrl, err := url.Parse("http://lubimyczytac.pl")
	if err != nil {
		fmt.Printf("Error URL: %s", err.Error())
		return
	}
	baseUrl.Path += "szukaj/ksiazki"
	params := url.Values{}
	params.Add("phrase", find)
	params.Add("main_search", "1")
	baseUrl.RawQuery = params.Encode()
	url := baseUrl.String()
	fmt.Printf("Encoded URL is %q\n", url)

	books := lubimyczytac.SzukajLubimyCzytac(url)
	fmt.Printf("API /books/%s: %s\n", find, books)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(books)
}

func metricsLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	b, _ := json.Marshal(mem)
	fmt.Fprintf(w, string(b))
}

func homeLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	infoApi := info.NewInfoApi()
	infoApi.Name = "API-golang"
	infoApi.Version = "1.0.0"
	b, _ := json.Marshal(infoApi)
	fmt.Fprintf(w, string(b))
}

func main() {
	router := httprouter.New()
	router.GET("/", homeLink)
	router.GET("/metrics", metricsLink)
	router.GET("/books", searchBook)
	router.GET("/films", searchFilm)
	log.Fatal(http.ListenAndServe(":8080", router))
}
