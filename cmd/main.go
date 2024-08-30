package main

import (
	"github.com/fabiovige/golang-maieutica-backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ResponsavelController := controller.NewResponsavelController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/responsaveis", ResponsavelController.GetResponsaveis)

	server.Run(":8000")
}
