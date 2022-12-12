package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var source []string
var flag int = 1

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("bing.com", "www.bing.com"),
	)
	c.OnHTML("div.b_focusTextLarge", func(h *colly.HTMLElement) {
		source = append(source, h.Text)
	})
	if flag == 0 {
		c.OnHTML("div.b_focusTextMedium", func(h *colly.HTMLElement) {
			source = append(source, h.Text)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		fmt.Print("we are in\n")
		flag = 1
	})
	c.Visit("https://www.bing.com/search?q=Atlanta,GA iata code")
	c.Visit("https://www.bing.com/search?q=Valdosta,GA iata code")
	fmt.Printf("%v", source)
}
