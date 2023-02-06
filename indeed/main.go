package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Menentukan selector elemen yang ingin diambil
	c.OnHTML("#resultsCol .jobsearch-SerpJobCard", func(e *colly.HTMLElement) {
		title := e.ChildText(".title>a")
		company := e.ChildText(".sjcl .company")
		location := e.ChildText(".sjcl .location")

		fmt.Printf("Judul: %s\nPerusahaan: %s\nLokasi: %s\n\n", strings.TrimSpace(title), strings.TrimSpace(company), strings.TrimSpace(location))
	})

	c.Visit("https://www.indeed.com/jobs?q=web+developer&l=New+York%2C+NY")
}
