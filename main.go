package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	fmt.Print("Lets make a stock market scraper!!")

	ticker := []string{
		"AAPL",
		"MSFT",
		"GOOGL",
		"AMZN",
		"TSLA",
		"META",
		"NVDA",
		"NFLX",
		"INTC",
		"AMD",
		"BABA",
		"ORCL",
		"IBM",
		"ADBE",
		"CRM",
		"PYPL",
		"UBER",
		"LYFT",
		"SHOP",
		"SQ",
	}

	stocks := []Stock{}

	c := colly.NewCollector() // innitialize a instance of colly

	c.OnRequest(func(r *colly.Request) { // handling req
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { // handling error
		log.Println("Something went Wrong: ", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) { // handling main logic

		stock := Stock{}

		stock.company = e.ChildText("h1")
		stock.price = e.ChildText("fin-streamer[date-field='regularMarketPrice']")
		stock.change = e.ChildText("fin-streamer[date-field='regularMarketChangePercent']")

		stocks = append(stocks, stock)

	})

}
