package domain

// CreateProdutoDTO representa os dados necessários para criar um produto
type CreateProdutoDTO struct {
	Nome       string  `json:"nome" validate:"required"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco" validate:"required,gt=0"`
	Quantidade int     `json:"quantidade" validate:"required,gte=0"`
}

// UpdateProdutoDTO representa os dados necessários para atualizar um produto
type UpdateProdutoDTO struct {
	Nome       string  `json:"nome" validate:"required"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco" validate:"required,gt=0"`
	Quantidade int     `json:"quantidade" validate:"required,gte=0"`
}

type CreateItemVendaDTO struct {
	ProdutoID  string `json:"produto_id" validate:"required"`
	Quantidade int    `json:"quantidade" validate:"required,gt=0"`
}

type CreateVendaDTO struct {
	Cliente  string               `json:"cliente" validate:"required"`
	Itens    []CreateItemVendaDTO `json:"itens" validate:"required,dive"`
	Desconto float64              `json:"desconto,omitempty" validate:"gte=0,lte=100"`
}

type UpdateVendaDTO struct {
	Cliente string      `json:"cliente" validate:"required"`
	Itens   []ItemVenda `json:"itens" validate:"required,dive"`
}
