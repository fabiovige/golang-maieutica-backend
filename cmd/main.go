package main

import (
	"github.com/fabiovige/golang-maieutica-backend/controller"
	"github.com/fabiovige/golang-maieutica-backend/db"
	"github.com/fabiovige/golang-maieutica-backend/repository"
	"github.com/fabiovige/golang-maieutica-backend/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// conexao com o banco
	dbConn, err := db.ConnDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	responsavelRepository := repository.NewResponsavelRepository(dbConn)

	// camada usecase
	responsavelUseCase := usecase.NewResponsavelUseCase(responsavelRepository)

	// camadas de controller
	ResponsavelController := controller.NewResponsavelController(responsavelUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/responsaveis", ResponsavelController.GetResponsaveis)

	server.Run(":8000")
}
