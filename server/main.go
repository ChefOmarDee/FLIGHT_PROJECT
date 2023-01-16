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
	PassengerCount      int    `json:"passengercount"`
}
type bookedTicket struct {
	SourceLocation      string `json:"sourcelocation"`
	DestinationLocation string `json:"destinationlocation"`
	SourceDate          string `json:"sourcedate"`
	PassengerCount      int    `json:"passengercount"`
	FlightPrice         int    `json:"flightprice"`
	FlightTime          string `json:"flighttime"`
	FlightAirline       string `json:"flightairline"`
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
	userBookedTicket := scraper(newTicket.SourceLocation, newTicket.DestinationLocation, newTicket.SourceDate, newTicket.PassengerCount)
	flightBooking.SourceLocation = newTicket.SourceLocation
	flightBooking.FlightTime = userBookedTicket.flightTime
	flightBooking.DestinationLocation = newTicket.DestinationLocation
	flightBooking.FlightPrice = userBookedTicket.flightPrice
	flightBooking.PassengerCount = newTicket.PassengerCount
	flightBooking.SourceDate = newTicket.SourceDate
	flightBooking.FlightAirline = userBookedTicket.flightAirline
	print("\nyyyyyyyyyyyyyyyyy")
	fmt.Print(flightBooking)
	print("\nyyyyyyyyyyyyyyyyy")
}
