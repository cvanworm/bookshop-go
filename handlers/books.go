package handlers

import (
	"github.com/andey-robins/bookshop-go/db"
	"github.com/andey-robins/bookshop-go/validate"
	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Author string  `json:"author" validate:"required"`
	Price  float32 `json:"price" validate:"required"`
}

type BookPrice struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func CreateBook(c *gin.Context) {
	var json Book
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Validate(json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.CreateBook(json.Title, json.Author, json.Price)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func GetPrice(c *gin.Context) {
	var json BookPrice
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Validate(json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bid, err := db.GetBookId(json.Title, json.Author)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	price, err := db.GetBookPrice(bid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"price": price})
}
