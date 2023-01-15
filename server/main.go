package main

import (
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

var tickets = []ticket{
	{
		SourceLocation:      "Vero Beach, FL",
		DestinationLocation: "Kansas City, MO",
		SourceDate:          "12/12/2022",
		PassengerCount:      1,
	},
}

func main() {
	scraper()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/tickets", GetTicket)
	router.POST("/tickets", AddTicket)
	router.Run("localhost:8080")
}

func GetTicket(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tickets)
}
func AddTicket(c *gin.Context) {
	var newTicket ticket

	if err := c.BindJSON(&newTicket); err != nil {
		return
	}
	tickets = append(tickets, newTicket)
	c.IndentedJSON(http.StatusCreated, newTicket)
}
