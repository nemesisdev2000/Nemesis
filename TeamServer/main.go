package main

import (
	"github.com/gin-gonic/gin"

	"github.com/nemesisdev2000/Nemesis/TeamServer/ClientServices"
	"github.com/nemesisdev2000/Nemesis/TeamServer/HandleListeners"
)

func main() {
	router := gin.Default()
	router.POST("/login", ClientServices.Login)
	router.POST("/listen", HandleListeners.HandleListener)

	router.Run("localhost:8000")
}
