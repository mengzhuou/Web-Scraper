package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type ScrapedContent struct {
	novelName string
	novelURL string
}

func main() {
	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatalln("Fail", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"novelName", "novelURL"}
	writer.Write(headers)

	c := colly.NewCollector()
	c.Visit("https://books.toscrape.com")
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		scrapedContent := ScrapedContent{}

		scrapedContent.novelName = e.ChildText(".price_color")
		scrapedContent.novelURL = e.ChildAttr(".image_container img", "alt")
		row := []string{scrapedContent.novelName, scrapedContent.novelURL}

		writer.Write(row)
	})
	webURL := "https://books.toscrape.com/"
	c.Visit(webURL)
}
