package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"./lubimyczytac"

	"github.com/julienschmidt/httprouter"
)

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
	json.NewEncoder(w).Encode(books)
}

func homeLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "API-golang")
}

func main() {
	router := httprouter.New()
	router.GET("/", homeLink)
	router.GET("/books", searchBook)
	log.Fatal(http.ListenAndServe(":8080", router))
}
