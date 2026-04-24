# Marketix

A simple, lightweight stock market scraper built with Go and the [Colly](https://github.com/gocolly/colly) framework.

## Features
- Fetches real-time stock prices and percentage changes from Yahoo Finance.
- Uses anti-blocking measures (User-Agent rotation and rate limiting).
- Saves results into a clean `stocks.csv` file.

## Prerequisites
- [Go](https://go.dev/doc/install) 1.21 or higher.

## Installation
Clone the repository and install dependencies:
```bash
go mod tidy
```

## Usage
Run the scraper:
```bash
go run main.go
```
The data will be saved to `stocks.csv` in the project root.

## Configuration
To scrape different stocks, simply update the `ticker` slice in `main.go`:
```go
ticker := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA"}
```
