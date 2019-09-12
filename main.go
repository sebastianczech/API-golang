package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./lubimyczytac"

	"github.com/julienschmidt/httprouter"
)

func searchBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	find := queryValues.Get("find")
	url := fmt.Sprintf("http://lubimyczytac.pl/szukaj/ksiazki?phrase=%s&main_search=1", find)
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
