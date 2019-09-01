package lubimyczytac

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type lubimyCzytacBook struct {
	author string
	title  string
	image  string
	book   string
}

func newLubimyCzytacBook() *lubimyCzytacBook {
	return &lubimyCzytacBook{
		author: "",
		title:  "",
		image:  "",
		book:   "",
	}
}

func (book lubimyCzytacBook) String() string {
	return fmt.Sprintf("Author: %s, title: %s, image: %s, book: %s", book.author, book.title, book.image, book.book)
}

// SzukajLubimyCzytac API dla lubimyczytac.pl do wyszukiwania ksiazek
func SzukajLubimyCzytac(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("lubimyczytac.pl"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(e.Request.URL.String(), "/szukaj/") && len(e.Text) > 0 && strings.Contains(link, "ksiazka") && !strings.Contains(e.Text, "Dodaj książkę") && !strings.Contains(e.Text, "Kup książkę") {
			fmt.Printf("Book link found: %q -> %s\n", e.Text, link)
			c.Visit(e.Request.AbsoluteURL(link))

			book := newLubimyCzytacBook()
			book.title = e.Text
			book.book = link
			fmt.Printf("Struct: %s\n", book)
		}
		if strings.Contains(e.Request.URL.String(), "/ksiazka/") && len(e.Text) > 0 && strings.Contains(link, "autor") && strings.Contains(e.Attr("itemprop"), "name") {
			fmt.Printf("Author link found: %q -> %s\n", e.Text, link)

			book := newLubimyCzytacBook()
			book.author = e.Text
			fmt.Printf("Struct: %s\n", book)

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
}
