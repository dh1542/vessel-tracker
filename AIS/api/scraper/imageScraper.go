package scraper

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

// ScrapeImageUrlForShipName
// https://www.zenrows.com/blog/web-scraping-golang#get-page-html/**
func ScrapeImageUrlForShipName(shipName string) string {
	query := shipName + " ship"
	url := "https://duckduckgo.com/?q=" + url.QueryEscape(query) + "&iax=images&ia=images"
	fmt.Println(url)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; CollyBot/1.0; +https://github.com/gocolly/colly)"),
		colly.AllowedDomains("duckduckgo.com", "www.duckduckgo.com"),
	)

	c.OnHTML("body[div]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	c.Visit(url)
	return ""
}
