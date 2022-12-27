package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type flightInfo struct {
	flightPrice   int
	flightTime    string
	flightAirline string
}

func main() {
	var source []string
	var so string = "San+Francisco,CA"
	var de string = "Atlanta,GA"
	var flag int = 1
	var count int = 0
	var tempIndex int
	var tempStrings string
	var bookedFlightInfo flightInfo
	c := colly.NewCollector(
		colly.AllowedDomains("bing.com", "www.bing.com"),
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("we are in\n")
		flag = 1
	})
	so = "https://www.bing.com/search?q=" + so + "+IATA+code"
	de = "https://www.bing.com/search?q=" + de + "+IATA+code"
	fmt.Printf(so)
	fmt.Printf(de)
	c.Visit(so)
	c.Visit(de)
	print(count)
	if count == 2 {
		fmt.Printf("%v", source)
		url := "https://priceline-com-provider.p.rapidapi.com/v2/flight/departures?sid=iSiX639&departure_date=2023-02-25&adults=1&origin_airport_code=" + source[0] + "&destination_airport_code=" + source[1] + "&results_per_page=1"

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-RapidAPI-Key", "47d048e09dmsh82922bd4aa60f6ep15bd6bjsnf22dbc12cd4b")
		req.Header.Add("X-RapidAPI-Host", "priceline-com-provider.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		// fmt.Println(res)
		flightData := strings.Split(string(body), ",")
		for _, v := range flightData {
			if strings.Contains(v, "baseline_total_fare_per_ticket") {
				var tempIndex int = strings.Index(v, ":") + 1
				var tempStrings string = v[tempIndex:(len(v))]
				TEMP, _ := strconv.ParseFloat(tempStrings, 64)
				bookedFlightInfo.flightPrice = int(TEMP)
				fmt.Print("\n", bookedFlightInfo.flightPrice, "\n")
			} else if strings.Contains(v, "\"marketing_airline\"") {
				tempIndex = strings.Index(v, ":") + 2
				tempStrings = v[tempIndex : (len(v))-1]
				bookedFlightInfo.flightAirline = tempStrings
				fmt.Print("\n", bookedFlightInfo.flightAirline, "\n")
				//fmt.Println("\n", v)
			} else if strings.Contains(v, "\"duration\"") {
				tempIndex = strings.Index(v, ":") + 2
				tempStrings = v[tempIndex : (len(v))-1]
				bookedFlightInfo.flightTime = tempStrings
				fmt.Print("\n", bookedFlightInfo.flightTime, "\n")
			}
		}
		print(len(flightData))
	}
	fmt.Printf("%v", source)
	//print(price)
}
