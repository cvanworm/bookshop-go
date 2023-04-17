package handlers

import (
	"fmt"

	"github.com/andey-robins/bookshop-go/db"
	"github.com/andey-robins/bookshop-go/validate"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	Id             int     `json:"id"`
	Name           string  `json:"name" validate:"required"`
	ShippingAddr   string  `json:"shippingAddr" validate:"required"`
	AccountBalance float32 `json:"accountBalance"`
}

type UpCustomer struct {
	Id           int    `json:"id" validate:"required"`
	ShippingAddr string `json:"shippingAddr" validate:"required"`
}

type CustomerId struct {
	Id int `json:"id" validate:"required"`
}

func CreateCustomer(c *gin.Context) {
	var json Customer
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Validate(json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.CreateCustomer(json.Name, json.ShippingAddr)
	if err != nil {
		fmt.Println("FAILED")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func UpdateCustomerAddress(c *gin.Context) {
	var json UpCustomer
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Validate(json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateCustomerAddress(json.Id, json.ShippingAddr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func GetCustomerBalance(c *gin.Context) {
	var json CustomerId
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Validate(json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance, err := db.CustomerBalance(json.Id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"balance": balance})
}
