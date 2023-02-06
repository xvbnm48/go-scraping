package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Course stores information about a coursera course
func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36"

	// Menentukan selector elemen yang ingin diambil
	c.OnHTML("#resultsCol .jobsearch-SerpJobCard", func(e *colly.HTMLElement) {
		title := e.ChildText(".title>a")
		company := e.ChildText(".sjcl .company")
		location := e.ChildText(".sjcl .location")

		fmt.Printf("Judul: %s\nPerusahaan: %s\nLokasi: %s\n\n", strings.TrimSpace(title), strings.TrimSpace(company), strings.TrimSpace(location))
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.indeed.com/jobs?q=web+developer&l=New+York%2C+NY")
}
