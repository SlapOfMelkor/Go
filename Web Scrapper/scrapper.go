package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	secim := flag.String("site", "", "Site secimi. \n Ornegin: -site 1 ya da 2 \n 1 - The Hacker News \n 2 - Rust Devblogs & Updates")
	tarihF := flag.Bool("date", false, "Tarih bilgileri.\n Ornegin: -site 1 ya da 2 -date\n")
	icerikF := flag.Bool("desc", false, "Aciklamayi goster.\n Ornegin: -site 1 ya da 2 -desc\n")
	flag.Parse()
	///
	if *secim == "" {
		log.Fatal("Parametre eksik. -h ile mevcut secenekleri gorebilirsiniz")
	}
	sites := strings.Split(*secim, ",")
	for _, site := range sites {
		switch site {
		case "1":
			hackernews("https://thehackernews.com/", *tarihF, *icerikF)
		case "2":
			rustnews("https://rust.facepunch.com/news", *tarihF, *icerikF)

		default:
			log.Fatalf("Hatali Giris. Desteklenen degerler: 1 veya 2")
		}
	}
}

func hackernews(scrapeURL string, tarihF bool, icerikF bool) {

	c := colly.NewCollector(colly.AllowedDomains("thehackernews.com"))

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Hata: %s\n", err.Error())
	})

	c.OnHTML("div.body-post", func(h *colly.HTMLElement) {
		secim := h.DOM

		date := secim.Find("span.h-datetime").Text()
		title := secim.Find("h2.home-title").Text()
		description := secim.Find("div.home-desc").Text()
		trimmedDate := date[3:]

		if !tarihF && !icerikF {
			fmt.Printf("Haber Basligi: %s\n\n Tarih: %s\n\n Aciklama: %s\n\n\n\n", title, trimmedDate, description)

		} else if tarihF && !icerikF {

			fmt.Printf("Haber Basligi: %s\n\n Aciklama: %s \n\n", title, description)

		} else if !tarihF && icerikF {
			fmt.Printf("Haber Basligi: %s\n\n Tarih: %s\n\n \n\n", title, trimmedDate)

		} else {
			fmt.Printf("Haber Basligi: %s\n\n \n\n", title)

		}

	})

	c.Visit(scrapeURL)
}

func rustnews(scrapeURL string, tarihF bool, icerikF bool) {
	c := colly.NewCollector(colly.AllowedDomains("https://rust.facepunch.com/", "rust.facepunch.com"))

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error: %s\n", err.Error())
	})

	c.OnHTML("div.blog-post", func(h *colly.HTMLElement) {
		selection := h.DOM

		title := selection.Find("a h1").Text()

		description := selection.Find("p").Text()

		date := selection.Find("div.tag").Text()

		if !tarihF && !icerikF {
			fmt.Printf("Haber Basligi: %s\n\n Tarih: %s\n\n Aciklama: %s\n\n\n\n", title, date, description)

		} else if tarihF && !icerikF {

			fmt.Printf("Haber Basligi: %s\n\n Aciklama: %s \n\n", title, description)

		} else if !tarihF && icerikF {
			fmt.Printf("Haber Basligi: %s\n\n Tarih: %s\n\n \n\n", title, date)

		} else {
			fmt.Printf("Haber Basligi: %s\n\n \n\n", title)

		}

	})

	c.Visit(scrapeURL)
}
