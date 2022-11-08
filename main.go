package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json: "id"`
	Title  string `json: "title"`
	Author string `json: "author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J.K Rowling"},
	{ID: "2", Title: "The Lord Of The Rings", Author: "J.R.R Tolkein"},
	{ID: "3", Title: "The Wizard Of Oz", Author: "L. Frank Baum"},
}

func main() {
	r := gin.New()
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.Run()
}