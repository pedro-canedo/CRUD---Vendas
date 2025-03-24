package domain

import "time"

// Produto representa um item que pode ser vendido
type Produto struct {
	ID          int64     `json:"id"`
	Nome        string    `json:"nome"`
	Descricao   string    `json:"descricao"`
	Preco       float64   `json:"preco"`
	Quantidade  int       `json:"quantidade"`
	DataCriacao time.Time `json:"data_criacao"`
}

// ProdutoRepository define as operações que podem ser realizadas com produtos
type ProdutoRepository interface {
	Create(produto *Produto) error
	GetByID(id int64) (*Produto, error)
	GetAll() ([]Produto, error)
	Update(produto *Produto) error
	Delete(id int64) error
}

// ProdutoService define a lógica de negócio relacionada a produtos
type ProdutoService interface {
	CreateProduto(produto *Produto) error
	GetProduto(id int64) (*Produto, error)
	ListProdutos() ([]Produto, error)
	UpdateProduto(produto *Produto) error
	DeleteProduto(id int64) error
}
