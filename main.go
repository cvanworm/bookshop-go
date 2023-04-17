package main

import (
	"log"
	"os"

	"github.com/andey-robins/bookshop-go/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	defer file.Close()

	router := gin.Default()

	router.POST("/books/new", handlers.CreateBook)
	router.GET("/books/price", handlers.GetPrice)

	router.POST("/customers/new", handlers.CreateCustomer)
	router.PUT("/customers/updateAddress", handlers.UpdateCustomerAddress)
	router.GET("/customers/balance", handlers.GetCustomerBalance)

	router.POST("/orders/new", handlers.CreateOrder)
	router.GET("/orders/shipped", handlers.GetShipmentStatus)
	router.PUT("/orders/ship", handlers.ShipOrder)
	router.GET("/orders/status", handlers.GetOrderStatus)

	router.Run(":8080")
}
