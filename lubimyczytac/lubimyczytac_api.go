package lubimyczytac

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gocolly/colly"
)

// LubimyCzytacBook struktura danych do wykorzystania przez API
type LubimyCzytacBook struct {
	Author string `json:"Author"`
	Title  string `json:"Title"`
	Image  string `json:"Image"`
	Book   string `json:"Book"`
}

func newLubimyCzytacBook() *LubimyCzytacBook {
	return &LubimyCzytacBook{
		Author: "",
		Title:  "",
		Image:  "",
		Book:   "",
	}
}

func (book LubimyCzytacBook) String() string {
	return fmt.Sprintf("author: %s title: %s image: %s book: %s", book.Author, book.Title, book.Image, book.Book)
}

// SzukajLubimyCzytac API dla lubimyczytac.pl do wyszukiwania ksiazek
func SzukajLubimyCzytac(url string) []*LubimyCzytacBook {
	books := []*LubimyCzytacBook{}

	c := colly.NewCollector(
		colly.AllowedDomains("lubimyczytac.pl"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(e.Request.URL.String(), "/szukaj/") && len(e.Text) > 0 && strings.Contains(link, "ksiazka") && !strings.Contains(e.Text, "Dodaj książkę") && !strings.Contains(e.Text, "Kup książkę") && !strings.Contains(e.Text, "Sprawdź cenę") {
			fmt.Printf("Book link found: %q -> %s\n", e.Text, link)

			book := newLubimyCzytacBook()
			book.Title = strings.TrimFunc(e.Text, func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			})
			book.Book = e.Request.AbsoluteURL(link)
			// fmt.Printf("Struct: %s\n", book)
			books = append(books, book)

			c.Visit(e.Request.AbsoluteURL(link))
		}
		if strings.Contains(e.Request.URL.String(), "/ksiazka/") && len(e.Text) > 0 && strings.Contains(link, "autor") && strings.Contains(e.Attr("class"), "link-name") {
			fmt.Printf("Author link found: %q -> %s\n", e.Text, link)

			for _, book := range books {
				// fmt.Printf("Check %s with book %s\n", e.Request.URL.String(), book)
				if book.Book == e.Request.URL.String() {
					book.Author = strings.TrimFunc(e.Text, func(r rune) bool {
						return !unicode.IsLetter(r) && !unicode.IsNumber(r)
					})
					// fmt.Printf("Struct: %s\n", book)
				}
			}
		}
	})

	c.OnHTML("img[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "img-fluid book-cover d-lg-none" {
			for _, book := range books {
				if book.Book == e.Request.URL.String() {
					fmt.Printf("Image link found: %q -> %s\n", book.Title, e.Attr("src"))
					book.Image = e.Attr("src")
				}
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	c.Wait()
	fmt.Printf("List of structs: %s\n", books)

	return books
}
