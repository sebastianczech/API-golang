package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func searchBook(w http.ResponseWriter, r *http.Request) {
	searchBy := mux.Vars(r)["searchBy"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API-golang")
}

func main() {
	//lubimyczytac.SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+nietoperz&main_search=1")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
