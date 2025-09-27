package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

// ScrapeImageUrlForShipName
// https://www.zenrows.com/blog/web-scraping-golang#get-page-html/**
func ScrapeImageUrlForShipName(shipName string) string {
	c := colly.NewCollector(
		colly.AllowedDomains("https://images.google.com/"))

	c.Visit("https://images.google.com/")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	return ""
}
