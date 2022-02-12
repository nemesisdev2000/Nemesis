package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type test struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var tests = []test{
	{ID: "1", Title: "Test1"},
	{ID: "2", Title: "Test2"},
}

func main() {
	router := gin.Default()
	router.GET("/tests", getData)
	router.GET("/tests/:id", getItemByID)
	router.POST("/tests", addData)

	router.Run("localhost:8000")
	fmt.Println("Running")
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tests)
}

func addData(c *gin.Context) {
	var newItem test
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	tests = append(tests, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range tests {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
