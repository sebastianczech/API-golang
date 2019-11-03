package lubimyczytac

import (
	"testing"
)

func TestSzukajLubimyCzytacDlaKonkretnegoTytułu(t *testing.T) {
	books := SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+nietoperz&main_search=1")
	if len(books) != 1 {
		t.Errorf("found %d books; wanted 1", len(books))
	}

	if len(books[0].Title) == 0 {
		t.Error("Empty title")
	}
	if len(books[0].Author) == 0 {
		t.Error("Empty author")
	}
	if len(books[0].Image) == 0 {
		t.Error("Empty image URL")
	}
	if len(books[0].Book) == 0 {
		t.Error("Empty book URL")
	}
}

func TestSzukajLubimyCzytacDlaFragmentuTytułu(t *testing.T) {
	books := SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+nie&main_search=1")
	if len(books) <= 1 {
		t.Errorf("found %d books; wanted more than 1", len(books))
	}
}
