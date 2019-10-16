package imdb

import (
	"fmt"

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

	// ...

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	c.Wait()
	fmt.Printf("List of structs: %s\n", films)

	return films
}
