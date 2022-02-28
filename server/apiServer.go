package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nemesisdev2000/Nemesis/server/C2Server"
)

type test struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var tests = []test{
	{ID: "1", Title: "Test1"},
	{ID: "2", Title: "Test2"},
}

type listenerProfile struct {
	Type string `json:"type"`
	Port string `json:"port"`
}

func main() {
	router := gin.Default()
	router.GET("/tests/:id", getItemByID)
	router.GET("/stoplistener/:id", stopListener)
	router.POST("/tests", addData)
	router.POST("/listen", startListener)

	router.Run("localhost:8000")
	fmt.Println("Running")
}

func startListener(c *gin.Context) {
	var lp listenerProfile
	if err := c.BindJSON(&lp); err != nil {
		return
	}

	port := lp.Port
	ltype := lp.Type

	l := C2Server.Listen(port, ltype)
	c.IndentedJSON(http.StatusOK, lp)
	c.Data(http.StatusOK, "Content-Type: text/html", []byte(l.RemoteAddr().String()))
}

func stopListener(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("ID : ", id)
	C2Server.Stop(id)
	c.Data(http.StatusOK, "Content-Type: text/html", []byte("Stopped listener"))
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
