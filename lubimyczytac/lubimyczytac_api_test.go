package lubimyczytac

import (
	"testing"
)

func TestSzukajLubimyCzytacDlaKonkretnegoTytułu(t *testing.T) {
	books := SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+nietoperz&main_search=1")
	if len(books) != 1 {
		t.Errorf("found %d books; wanted 1", len(books))
	}
}

func TestSzukajLubimyCzytacDlaFragmentuTytułu(t *testing.T) {
	books := SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+niet&main_search=1")
	if len(books) <= 1 {
		t.Errorf("found %d books; wanted more than 1", len(books))
	}
}
