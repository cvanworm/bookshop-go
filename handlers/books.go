package handlers

import (
	"github.com/andey-robins/bookshop-go/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Author string  `json:"author" validate:"required"`
	Price  float32 `json:"price" validate:"price,required"`
}

func CreateBook(c *gin.Context) {
	var json Book
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(json)
	if err != nil {
		// log out this error
		// return a bad request and a helpful error message
		// if you wished, you could concat the validation error into this
		// message to help point your consumer in the right direction.
		c.JSON(400, gin.H{"error": "failed to validate struct: " + err.Error()})
		return
	}

	_, err = db.CreateBook(json.Title, json.Author, json.Price)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func GetPrice(c *gin.Context) {
	var json Book
	if err := c.BindJSON(&json); err != nil {
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
