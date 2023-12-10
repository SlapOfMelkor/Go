package main

//
import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type NewsItem struct {
	Title       string
	Date        string
	Description string
}

func rustnews() string {
	var scrapeURL string = "https://rust.facepunch.com/news"
	var newsData string
	c := colly.NewCollector(colly.AllowedDomains("rust.facepunch.com"))
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error: %s\n", err.Error())
	})
	c.OnHTML("div.blog-post", func(h *colly.HTMLElement) {
		selection := h.DOM

		title := selection.Find("a h1").Text()
		description := selection.Find("p").Text()
		date := selection.Find("div.tag").Text()
		newsData += fmt.Sprintf("Haber Basligi: %s\n\n Tarih: %s\n\n Aciklama: %s\n\n\n\n", title, date, description)
	})
	c.Visit(scrapeURL)
	return newsData
}
