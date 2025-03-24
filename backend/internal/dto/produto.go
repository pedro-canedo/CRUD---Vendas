package dto

type CreateProdutoDTO struct {
	Nome       string  `json:"nome" binding:"required"`
	Descricao  string  `json:"descricao" binding:"required"`
	Preco      float64 `json:"preco" binding:"required,gt=0"`
	Quantidade int     `json:"quantidade" binding:"required,gte=0"`
	ImagemURL  string  `json:"imagem_url"`
}

type UpdateProdutoDTO struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco" binding:"omitempty,gt=0"`
	Quantidade int     `json:"quantidade" binding:"omitempty,gte=0"`
	ImagemURL  string  `json:"imagem_url"`
}
