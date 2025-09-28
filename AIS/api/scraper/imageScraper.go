package scraper

import (
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

// ScrapeImageUrlForShipName
// https://www.zenrows.com/blog/web-scraping-golang#get-page-html/**
func ScrapeImageUrlForShipName(shipName string) string {
	query := shipName + " ship"

	url := "https://suche.aol.de/aol/image;_ylt=Awr.pOy_etloKkwgBoE8CmVH;_ylu=Y29sbwNpcjIEcG9zAzEEdnRpZAMEc2VjA3BpdnM-?q=" + url.QueryEscape(query) + "&v_t=aolde-homePage50"

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; CollyBot/1.0; +https://github.com/gocolly/colly)"),
		colly.AllowedDomains("suche.aol.de", "www.duckduckgo.com"),
	)

	result := ""

	c.OnHTML("#resitem-0 a", func(e *colly.HTMLElement) {
		result = e.Attr("href")

	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
