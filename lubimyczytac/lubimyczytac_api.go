package lubimyczytac

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// SzukajLubimyCzytac API dla lubimyczytac.pl do wyszukiwania ksiazek
func SzukajLubimyCzytac(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("lubimyczytac.pl"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(e.Request.URL.String(), "/szukaj/") && len(e.Text) > 0 && strings.Contains(link, "ksiazka") && !strings.Contains(e.Text, "Dodaj książkę") && !strings.Contains(e.Text, "Kup książkę") {
			fmt.Printf("Book link found: %q -> %s from URL: %s\n", e.Text, link, e.Request.URL)
			c.Visit(e.Request.AbsoluteURL(link))
		}
		if strings.Contains(e.Request.URL.String(), "/ksiazka/") && len(e.Text) > 0 && strings.Contains(link, "autor") && strings.Contains(e.Attr("itemprop"), "name") {
			fmt.Printf("Author link found: %q -> %s from URL: %s\n", e.Text, link, e.Request.URL)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
