package domain

import "time"

// Cliente representa um cliente do sistema
type Cliente struct {
	ID          string    `json:"id"`
	Nome        string    `json:"nome"`
	Email       string    `json:"email"`
	Telefone    string    `json:"telefone"`
	Endereco    string    `json:"endereco"`
	CPF         string    `json:"cpf"`
	UsuarioID   string    `json:"usuario_id"`
	DataCriacao time.Time `json:"data_criacao"`
}

// ClienteRepository define as operações que podem ser realizadas com clientes
type ClienteRepository interface {
	Create(cliente *Cliente) error
	GetByID(id string) (*Cliente, error)
	GetByCPF(cpf string) (*Cliente, error)
	GetAll() ([]Cliente, error)
	Update(cliente *Cliente) error
	Delete(id string) error
}

// ClienteService define a lógica de negócio relacionada a clientes
type ClienteService interface {
	CreateCliente(cliente *Cliente) error
	GetCliente(id string) (*Cliente, error)
	GetClienteByCPF(cpf string) (*Cliente, error)
	ListClientes() ([]Cliente, error)
	UpdateCliente(cliente *Cliente) error
	DeleteCliente(id string) error
}
