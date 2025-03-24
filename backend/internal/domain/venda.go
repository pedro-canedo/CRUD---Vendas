package domain

import "time"

// ItemVenda representa um item individual em uma venda
type ItemVenda struct {
	ID            string   `json:"id"`
	VendaID       string   `json:"venda_id"`
	ProdutoID     string   `json:"produto_id"`
	Quantidade    int      `json:"quantidade"`
	PrecoUnitario float64  `json:"preco_unitario"`
	Produto       *Produto `json:"produto"`
}

// Venda representa uma transação de venda
type Venda struct {
	ID          string      `json:"id"`
	ClienteID   string      `json:"cliente_id"`
	VendedorID  string      `json:"vendedor_id"`
	DataVenda   time.Time   `json:"data_venda"`
	ValorTotal  float64     `json:"valor_total"`
	DataCriacao time.Time   `json:"data_criacao"`
	Items       []ItemVenda `json:"items"`
	Cliente     *Usuario    `json:"cliente"`
	Vendedor    *Usuario    `json:"vendedor"`
}

// VendaRepository define as operações que podem ser realizadas com vendas
type VendaRepository interface {
	Create(venda *Venda) error
	GetByID(id string) (*Venda, error)
	GetAll() ([]Venda, error)
	Update(venda *Venda) error
	Delete(id string) error
}

// VendaService define a lógica de negócio relacionada a vendas
type VendaService interface {
	CreateVenda(venda *Venda) error
	GetVenda(id string) (*Venda, error)
	ListVendas() ([]Venda, error)
	UpdateVenda(venda *Venda) error
	DeleteVenda(id string) error
}
