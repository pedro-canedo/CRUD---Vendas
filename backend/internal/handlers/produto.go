package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"vendas/internal/domain"
	"vendas/internal/dto"
)

type ProdutoHandler struct {
	produtoService domain.ProdutoService
}

func NewProdutoHandler(service domain.ProdutoService) *ProdutoHandler {
	return &ProdutoHandler{
		produtoService: service,
	}
}

func (h *ProdutoHandler) CreateProduto(c *gin.Context) {
	var dto dto.CreateProdutoDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	produto := &domain.Produto{
		ID:          uuid.New().String(),
		Nome:        dto.Nome,
		Descricao:   dto.Descricao,
		Preco:       dto.Preco,
		Quantidade:  dto.Quantidade,
		ImagemURL:   dto.ImagemURL,
		DataCriacao: time.Now(),
	}

	if err := h.produtoService.CreateProduto(produto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, produto)
}

func (h *ProdutoHandler) GetProduto(c *gin.Context) {
	id := c.Param("id")
	produto, err := h.produtoService.GetProduto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, produto)
}

func (h *ProdutoHandler) ListProdutos(c *gin.Context) {
	produtos, err := h.produtoService.ListProdutos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produtos)
}

func (h *ProdutoHandler) UpdateProduto(c *gin.Context) {
	id := c.Param("id")
	var dto dto.UpdateProdutoDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	produto, err := h.produtoService.GetProduto(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	}

	// Atualiza apenas os campos fornecidos
	if dto.Nome != "" {
		produto.Nome = dto.Nome
	}
	if dto.Descricao != "" {
		produto.Descricao = dto.Descricao
	}
	if dto.Preco > 0 {
		produto.Preco = dto.Preco
	}
	if dto.Quantidade >= 0 {
		produto.Quantidade = dto.Quantidade
	}
	if dto.ImagemURL != "" {
		produto.ImagemURL = dto.ImagemURL
	}

	if err := h.produtoService.UpdateProduto(produto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produto)
}

func (h *ProdutoHandler) DeleteProduto(c *gin.Context) {
	id := c.Param("id")
	if err := h.produtoService.DeleteProduto(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "produto deletado com sucesso"})
}
