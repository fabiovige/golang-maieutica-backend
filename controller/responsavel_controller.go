package controller

import (
	"net/http"
	"time"

	"github.com/fabiovige/golang-maieutica-backend/model"
	"github.com/gin-gonic/gin"
)

type responsavelController struct {
	//@TODO adicionar usecase
}

// inicializa a struct responsavelController
func NewResponsavelController() responsavelController {
	return responsavelController{}
}

func (r *responsavelController) GetResponsaveis(ctx *gin.Context) {
	responsaveis := []model.Responsavel{
		{
			ID:             1,
			Nome:           "João da Silva",
			CPF:            "123.456.789-00",
			DataNascimento: time.Date(1980, 1, 15, 0, 0, 0, 0, time.UTC),
			Telefone:       "(11) 98765-4321",
			Email:          "joao.silva@example.com",
			Endereco:       "Rua das Flores",
			Numero:         "123",
			Complemento:    "Apto 45",
			Bairro:         "Jardim Primavera",
			Cidade:         "São Paulo",
			Estado:         "SP",
			CEP:            "01234-567",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			ID:             2,
			Nome:           "Maria Oliveira",
			CPF:            "987.654.321-00",
			DataNascimento: time.Date(1975, 5, 20, 0, 0, 0, 0, time.UTC),
			Telefone:       "(21) 99876-5432",
			Email:          "maria.oliveira@example.com",
			Endereco:       "Av. Central",
			Numero:         "456",
			Complemento:    "Casa 10",
			Bairro:         "Centro",
			Cidade:         "Rio de Janeiro",
			Estado:         "RJ",
			CEP:            "21000-000",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	ctx.JSON(http.StatusOK, responsaveis)
}
