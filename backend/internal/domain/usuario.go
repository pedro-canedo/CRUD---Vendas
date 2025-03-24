package domain

import "time"

// Role representa o papel do usuário no sistema
type Role string

const (
	RoleAdmin    Role = "admin"
	RoleVendedor Role = "vendedor"
	RoleCliente  Role = "cliente"
)

// Usuario representa um usuário do sistema
type Usuario struct {
	ID          string    `json:"id"`
	Nome        string    `json:"nome"`
	Email       string    `json:"email"`
	Senha       string    `json:"-"` // O "-" impede que a senha seja enviada no JSON
	Role        Role      `json:"role"`
	Ativo       bool      `json:"ativo"`
	DataCriacao time.Time `json:"data_criacao"`
}

// UsuarioRepository define as operações que podem ser realizadas com usuários
type UsuarioRepository interface {
	Create(usuario *Usuario) error
	GetByID(id string) (*Usuario, error)
	GetByEmail(email string) (*Usuario, error)
	GetAll() ([]Usuario, error)
	Update(usuario *Usuario) error
	Delete(id string) error
}

// UsuarioService define a lógica de negócio relacionada a usuários
type UsuarioService interface {
	CreateUsuario(usuario *Usuario) error
	GetUsuario(id string) (*Usuario, error)
	GetUsuarioByEmail(email string) (*Usuario, error)
	ListUsuarios() ([]Usuario, error)
	UpdateUsuario(usuario *Usuario) error
	DeleteUsuario(id string) error
	Autenticar(email, senha string) (*Usuario, error)
}
