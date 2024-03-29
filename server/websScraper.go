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
	sourceIata    string
	destIata      string
}

func scraper(sourceLocation string, destinationLocation string, sourceDate string) flightInfo {
	var source []string
	sourceLocation = strings.ReplaceAll(sourceLocation, " ", "+")
	destinationLocation = strings.ReplaceAll(destinationLocation, " ", "+")
	index := len(sourceLocation) - 2
	tempTimeSo := sourceLocation[:index] + ",+" + sourceLocation[index:]
	index = len(destinationLocation) - 2
	tempTimeDe := destinationLocation[:index] + ",+" + destinationLocation[index:]
	index = len(sourceLocation) - 2
	var so string = sourceLocation[:index] + "," + sourceLocation[index:]
	index = len(destinationLocation) - 2
	var de string = destinationLocation[:index] + "," + destinationLocation[index:]
	var flag int = 1
	var count int = 0
	var tempIndex int
	var tempStrings string
	var bookedFlightInfo flightInfo
	d := colly.NewCollector(
		colly.AllowedDomains("www.travelmath.com", "travelmath.com"),
	)
	d.OnHTML("h3", func(h *colly.HTMLElement) {
		bookedFlightInfo.flightTime = h.Text
		print(h.Text)
	})
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
	so = "https://www.bing.com/search?q=what+is+" + so + "+airport+code"
	de = "https://www.bing.com/search?q=what+is+" + de + "+airport+code"
	fmt.Printf(so)
	fmt.Printf(de)
	c.Visit(so)
	c.Visit(de)
	d.Visit("https://www.travelmath.com/flying-time/from/" + tempTimeSo + "/to/" + tempTimeDe)
	print(count)
	if count == 2 {
		url := "https://priceline-com-provider.p.rapidapi.com/v2/flight/departures?sid=iSiX639&departure_date=" + sourceDate + "&adults=1&origin_airport_code=" + source[0] + "&destination_airport_code=" + source[1] + "&results_per_page=1"

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-RapidAPI-Key", "47d048e09dmsh82922bd4aa60f6ep15bd6bjsnf22dbc12cd4b")
		req.Header.Add("X-RapidAPI-Host", "priceline-com-provider.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		flightData := strings.Split(string(body), ",")
		for _, v := range flightData {
			if strings.Contains(v, "baseline_total_fare_per_ticket") {
				var tempIndex int = strings.Index(v, ":") + 1
				var tempStrings string = v[tempIndex:(len(v))]
				TEMP, _ := strconv.ParseFloat(tempStrings, 64)
				bookedFlightInfo.flightPrice = int(TEMP)
			} else if strings.Contains(v, "\"marketing_airline\"") {
				tempIndex = strings.Index(v, ":") + 2
				tempStrings = v[tempIndex : (len(v))-1]
				bookedFlightInfo.flightAirline = tempStrings
				bookedFlightInfo.sourceIata = source[0]
				bookedFlightInfo.destIata = source[1]
			}
		}
	}

	fmt.Printf("%v", source)
	fmt.Print("\n", bookedFlightInfo.flightAirline, "\n")
	fmt.Print("\n", bookedFlightInfo.flightPrice, "\n")
	fmt.Print("\n", bookedFlightInfo.flightPrice, "\n")
	fmt.Print("\n", bookedFlightInfo.flightTime, "\n")
	fmt.Print("\n", bookedFlightInfo.flightTime, "\n")
	return (bookedFlightInfo)
}
