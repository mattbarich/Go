package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name    string `json:"name"`
	Price   string `json:"price"`
	ProdURL string `json:"prod_url"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	c.OnHTML("div.product-grid.grid.grid-cols-2.sm\\:grid-cols-3.lg\\:grid-cols-4.gap-4.mb-8 div.product-item.flex.flex-col.items-center.rounded-lg", func(e *colly.HTMLElement) {
		item := item{
			Name:    e.ChildText("span.product-name"),
			Price:   e.ChildText("span.product-price.text-slate-600"),
			ProdURL: e.ChildAttr("img", "src"),
		}

		fmt.Printf("Name: %s\nPrice: %s\nProduct URL: %s\n\n", item.Name, item.Price, item.ProdURL)
	})

	c.Visit("https://www.scrapingcourse.com/pagination")
}
