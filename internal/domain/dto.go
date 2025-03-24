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
