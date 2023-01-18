package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ticket struct {
	SourceLocation      string `json:"sourcelocation"`
	DestinationLocation string `json:"destinationlocation"`
	SourceDate          string `json:"sourcedate"`
}
type bookedTicket struct {
	SourceLocation      string `json:"sourcelocation"`
	DestinationLocation string `json:"destinationlocation"`
	SourceDate          string `json:"sourcedate"`
	FlightPrice         int    `json:"flightprice"`
	FlightTime          string `json:"flighttime"`
	FlightAirline       string `json:"flightairline"`
	SourceIata          string `json:"sourceiata"`
	DestIata            string `json:"destiata"`
}

// var tickets = bookedTicket{}
var flightBooking bookedTicket
var userBookedTicket flightInfo

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/gettickets", GetTicket)
	router.POST("/tickets", AddTicket)
	router.Run("localhost:8080")
}

func GetTicket(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, flightBooking)
}
func AddTicket(c *gin.Context) {
	var newTicket ticket

	if err := c.BindJSON(&newTicket); err != nil {
		return
	}
	print("\nyyyyyyyyyyyyyyyyy")

	fmt.Printf("%T", newTicket)
	print("\nyyyyyyyyyyyyyyyyy")

	c.IndentedJSON(http.StatusCreated, newTicket)
	userBookedTicket := scraper(newTicket.SourceLocation, newTicket.DestinationLocation, newTicket.SourceDate)
	flightBooking.SourceLocation = newTicket.SourceLocation
	flightBooking.FlightTime = userBookedTicket.flightTime
	flightBooking.DestinationLocation = newTicket.DestinationLocation
	flightBooking.FlightPrice = userBookedTicket.flightPrice
	flightBooking.SourceDate = newTicket.SourceDate
	flightBooking.FlightAirline = userBookedTicket.flightAirline
	flightBooking.SourceIata = userBookedTicket.sourceIata
	flightBooking.DestIata = userBookedTicket.destIata
	print("\nyyyyyyyyyyyyyyyyy")
	fmt.Print(flightBooking)
	print("\nyyyyyyyyyyyyyyyyy")
}
