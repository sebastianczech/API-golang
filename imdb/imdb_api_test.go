package imdb

import (
	"testing"
)

func TestFindImdbFilmDlaNieznanegoTytulu(t *testing.T) {
	films := FindImdbFilm("https://www.imdb.com/find?q=nieznany%20tytul")
	if len(films) != 0 {
		t.Errorf("found %d films; wanted 0", len(films))
	}
}

func TestFindImdbFilmDlaKonkretnegoTytulu(t *testing.T) {
	films := FindImdbFilm("https://www.imdb.com/find?q=zasada%20przyjemnosc")
	if len(films) != 1 {
		t.Errorf("found %d films; wanted 1", len(films))
	}
}
