package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./lubimyczytac"

	"github.com/gorilla/mux"
)

func searchBook(w http.ResponseWriter, r *http.Request) {
	find := mux.Vars(r)["find"]
	url := fmt.Sprintf("http://lubimyczytac.pl/szukaj/ksiazki?phrase=%s&main_search=1", find)
	books := lubimyczytac.SzukajLubimyCzytac(url)
	fmt.Printf("API /books/%s: %s\n", find, books)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "API-golang")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/books/{find}", searchBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
