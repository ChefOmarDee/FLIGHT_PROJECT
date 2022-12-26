package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//"strings"

	"github.com/gocolly/colly"
)

func askQuestions(source *string, destination *string) {
	fmt.Println("Which city and state would you like to fly from?")
	fmt.Scanln(source)
	fmt.Println("Which city and state would you like to fly to?")
	fmt.Scanln(destination)
}
func main() {
	var source []string
	//var data string=""
	var so string = ""
	var de string = ""
	var flag int = 1
	var count int = 0
	askQuestions(&so, &de)
	fmt.Printf(so)
	fmt.Printf(de)
	c := colly.NewCollector(
		colly.AllowedDomains("bing.com", "www.bing.com"),
	)
	exp := colly.NewCollector(
		colly.AllowedDomains("bing.com", "www.bing.com", "www.kayak.com", "kayak.com", "www.delta.com", "delta.com", "expedia.com", "www.expedia.com"),
	)

	c.OnHTML("div.b_focusTextLarge", func(h *colly.HTMLElement) {
		source = append(source, h.Text)
		count++
		flag = 0
	})
	if flag == 1 {
		c.OnHTML("div.b_focusTextMedium", func(h *colly.HTMLElement) {
			source = append(source, h.Text)
			count++

		})
	}
	exp.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		fmt.Println("giggity")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Print("we are in\n")
		flag = 1
	})
	so = "https://www.bing.com/search?q=" + so + " IATA code"
	de = "https://www.bing.com/search?q=" + de + " IATA code"
	c.Visit(so)
	c.Visit(de)
	print(count)
	if count == 2 {
		fmt.Printf("%v", source)
		url := "https://priceline-com-provider.p.rapidapi.com/v2/flight/departures?sid=iSiX639&departure_date=2022-12-25&adults=1&origin_airport_code=" + source[0] + "&destination_airport_code=" + source[1] + "&results_per_page=1"

		req, _ := http.NewRequest("GET", url, nil)
		//had to use priceline API as colly nor any other webscraper could scrape directly from html page immediately after the javascript loads the page.
		req.Header.Add("X-RapidAPI-Key", "47d048e09dmsh82922bd4aa60f6ep15bd6bjsnf22dbc12cd4b")
		req.Header.Add("X-RapidAPI-Host", "priceline-com-provider.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		flightData := strings.Split(string(body), ",")
		for _, v := range flightData {
			//fmt.Println(v)
			if strings.Contains(v, "fare") {
				fmt.Print(v)

			}
		}
		print(len(flightData))
		// fmt.Println(flightData)
		//fmt.Println(string(body))
	}
}

//link mianly used for flight search
//https://www.expedia.com/Flights-Search?trip=oneway&leg1=from:FLL,to:MCI,departure:4/5/2023TANYT&passengers=children:0,adults:1,seniors:0,infantinlap:N&mode=search&options=sortby:price&paandi=true
