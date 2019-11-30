package imdb

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gocolly/colly"
)

// IdmbFilm struktura danych do wykorzystania przez API
type IdmbFilm struct {
	Title string `json:"Title"`
	Tv    string `json:"Tv"`
	Image string `json:"Image"`
	Film  string `json:"Film"`
}

func newIdmbFilm() *IdmbFilm {
	return &IdmbFilm{
		Title: "",
		Tv:    "",
		Image: "",
		Film:  "",
	}
}

func (film IdmbFilm) String() string {
	return fmt.Sprintf("tv: %s title: %s image: %s film: %s", film.Tv, film.Title, film.Image, film.Film)
}

// FindImdbFilm API dla imdb.com do wyszukiwania filmow
func FindImdbFilm(url string) []*IdmbFilm {
	films := []*IdmbFilm{}

	c := colly.NewCollector(
		colly.AllowedDomains("www.imdb.com"),
	)

	c.OnHTML("td[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "result_text" {
			link := e.ChildAttr("a[href]", "href")
			fmt.Printf("Film link found: %s -> %s\n", e.ChildText("a[href]"), link)

			film := newIdmbFilm()
			film.Film = e.Request.AbsoluteURL(link)
			films = append(films, film)

			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnHTML("h1[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "long" {
			fmt.Printf("Title found: %s\n", e.Text)

			for _, film := range films {
				if film.Film == e.Request.URL.String() {
					film.Title = strings.TrimFunc(e.Text, func(r rune) bool {
						return !unicode.IsLetter(r) && !unicode.IsNumber(r)
					})
				}
			}
		}
	})

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		itemprop := e.Attr("class")
		if itemprop == "poster" {
			link := e.ChildAttr("img[src]", "src")
			fmt.Printf("Image link found: %s -> %s\n", e.ChildText("img[src]"), link)

			for _, film := range films {
				if film.Film == e.Request.URL.String() {
					film.Image = e.Request.AbsoluteURL(link)
				}
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	c.Wait()
	fmt.Printf("List of structs: %s\n", films)

	return films
}
