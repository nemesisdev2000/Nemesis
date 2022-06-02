package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/nemesisdev2000/Nemesis/TeamServer/ClientServices"
	"github.com/nemesisdev2000/Nemesis/TeamServer/HandleListeners"
)

func main() {

	bindhost := os.Args[1] + ":" + os.Args[2]
	router := gin.Default()
	router.POST("/login", ClientServices.Login)
	router.POST("/listen", HandleListeners.HandleListener)
	router.GET("/showListeners", HandleListeners.ShowListeners)
	router.POST("/stopListener", HandleListeners.StopListener)

	router.Run(bindhost)
}
