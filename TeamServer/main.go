package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nemesisdev2000/Nemesis/TeamServer/ClientServices"
)

func main() {
	router := gin.Default()
	router.POST("/login", ClientServices.Login)

	router.Run("localhost:8000")
}
