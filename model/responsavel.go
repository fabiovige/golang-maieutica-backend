package model

import "time"

type Responsavel struct {
	ID             int       `json:"id" db:"id"`
	Nome           string    `json:"nome" db:"nome"`
	CPF            string    `json:"cpf" db:"cpf"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Telefone       string    `json:"telefone" db:"telefone"`
	Email          string    `json:"email" db:"email"`
	Endereco       string    `json:"endereco" db:"endereco"`
	Numero         string    `json:"numero" db:"numero"`
	Complemento    string    `json:"complemento" db:"complemento"`
	Bairro         string    `json:"bairro" db:"bairro"`
	Cidade         string    `json:"cidade" db:"cidade"`
	Estado         string    `json:"estado" db:"estado"`
	CEP            string    `json:"cep" db:"cep"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
