package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title                string `json:"title"`
	Number_of_times_read int    `json:"number_of_times_read"`
	Type_of_book         string `json:"type_of_book"`
	Author               string `json:"author"`
}

var books = []Book{
	{Title: "test", Number_of_times_read: 1, Type_of_book: "audio", Author: "Phil"},
	{Title: "test2", Number_of_times_read: 4, Type_of_book: "book", Author: "Phil2"},
}

var port string = ":8080"

func returnAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addNewBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, gin.H{"message": "book created"})

}

func patchBook(c *gin.Context) {
	title := c.Param("title")
	for _, v := range books {
		if strings.EqualFold(title, v.Title) {
			v.Number_of_times_read++
			msg := fmt.Sprintf("Updating book and adding one to count for  %s", v.Title)
			c.JSON(http.StatusAccepted, gin.H{"message": msg})
		} else {
			msg := fmt.Sprintf("Book %s not found", v.Title)
			c.JSON(http.StatusNotFound, gin.H{"message": msg})
		}

	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})

}

func main() {
	r := gin.Default()
	r.GET("/api/health", healthCheck)
	r.GET("/api/books", returnAllBooks)
	r.POST("/api/book", addNewBook)
	r.PATCH("/api/book/:title", patchBook)
	r.Run(port)

}
