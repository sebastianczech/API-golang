package lubimyczytac

import (
	"testing"
)

func TestSzukajLubimyCzytacDlaKonkretnegoTytułu(t *testing.T) {
	SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek+nietoperz&main_search=1")
}

func TestSzukajLubimyCzytacDlaFragmentuTytułu(t *testing.T) {
	SzukajLubimyCzytac("http://lubimyczytac.pl/szukaj/ksiazki?phrase=cz%C5%82owiek&main_search=1")
}
