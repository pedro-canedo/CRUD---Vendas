package domain

import "time"

// ItemVenda representa um item individual em uma venda
type ItemVenda struct {
	ID            string  `json:"id"`
	ProdutoID     string  `json:"produto_id"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float64 `json:"preco_unitario"`
	Subtotal      float64 `json:"subtotal"`
}

// Venda representa uma transação de venda
type Venda struct {
	ID        string      `json:"id"`
	DataVenda time.Time   `json:"data_venda"`
	Itens     []ItemVenda `json:"itens"`
	Total     float64     `json:"total"`
	Cliente   string      `json:"cliente"`
	Desconto  float64     `json:"desconto,omitempty"`
}

// VendaRepository define as operações que podem ser realizadas com vendas
type VendaRepository interface {
	Create(venda *Venda) error
	GetByID(id string) (*Venda, error)
	GetAll() ([]Venda, error)
	Update(venda *Venda) error
	Delete(id string) error
	GetVendasPorCliente(cliente string) ([]Venda, error)
	GetVendasPorPeriodo(inicio, fim int64) ([]Venda, error)
}

// VendaService define a lógica de negócio relacionada a vendas
type VendaService interface {
	CreateVenda(venda *Venda) error
	GetVenda(id string) (*Venda, error)
	ListVendas() ([]Venda, error)
	UpdateVenda(venda *Venda) error
	DeleteVenda(id string) error
	CalcularTotal(venda *Venda) error
}
