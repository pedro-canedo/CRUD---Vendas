package dto

type CreateClienteDTO struct {
	Nome     string `json:"nome" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Telefone string `json:"telefone" binding:"required"`
	Endereco string `json:"endereco" binding:"required"`
	CPF      string `json:"cpf" binding:"required"`
}

type UpdateClienteDTO struct {
	Nome     string `json:"nome"`
	Email    string `json:"email" binding:"omitempty,email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
	CPF      string `json:"cpf"`
}
