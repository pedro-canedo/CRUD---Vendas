package handlers

import (
	"vendas/internal/domain"
)

type Handlers struct {
	Usuario *UsuarioHandler
	Cliente *ClienteHandler
	Produto *ProdutoHandler
}

func NewHandlers(
	usuarioService domain.UsuarioService,
	clienteService domain.ClienteService,
	produtoService domain.ProdutoService,
	jwtSecretKey string,
) *Handlers {
	return &Handlers{
		Usuario: NewUsuarioHandler(usuarioService, jwtSecretKey),
		Cliente: NewClienteHandler(clienteService),
		Produto: NewProdutoHandler(produtoService),
	}
}
