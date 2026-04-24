package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	fmt.Println()
	fmt.Println("A stock market scraper!! : ")
	fmt.Println()

	ticker := []string{
		"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA", "CRM", "PYPL", "UBER", "LYFT", "IBM", "ORCL",
	}

	stocks := []Stock{}

	// Initialize colly with a User-Agent to avoid being blocked
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	// Add a delay to be respectful and avoid rate limiting
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*finance.yahoo.com*",
		RandomDelay: 2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went Wrong: ", err)
	})

	// Updated selector for the main quote header
	c.OnHTML("section[data-testid='quote-hdr']", func(e *colly.HTMLElement) {
		stock := Stock{}

		// Updated selectors based on current Yahoo Finance layout
		stock.company = e.ChildText("h1")
		stock.price = e.ChildText("[data-testid='qsp-price']")
		stock.change = e.ChildText("[data-testid='qsp-price-change-percent']")

		if stock.company != "" {
			fmt.Printf("\nCompany: %s | Price: %s | Change: %s\n\n", stock.company, stock.price, stock.change)
			stocks = append(stocks, stock)
		}
	})

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	c.Wait()

	if len(stocks) == 0 {
		log.Println("No stocks were scraped. Yahoo Finance might be blocking the request or selectors might have changed.")
		return
	}

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"company", "price", "change"}
	writer.Write(headers)

	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)
	}
	fmt.Println("Scraping complete. Results saved to stocks.csv")
}
