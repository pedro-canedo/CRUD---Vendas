package main

import (
	"net/http"
	"strconv"
	"vendas/internal/domain"
	"vendas/internal/repository"
	"vendas/internal/service"

	_ "vendas/docs" // Importa os docs do Swagger

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Sistema de Vendas API
// @version 1.0
// @description API para gerenciamento de vendas e produtos
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Inicializa o Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Inicializa repositórios e serviços
	produtoRepo := repository.NewProdutoRepository()
	produtoService := service.NewProdutoService(produtoRepo)

	// Grupo de rotas API v1
	v1 := e.Group("/api/v1")

	// Rotas de Produtos
	produtos := v1.Group("/produtos")
	produtos.GET("", getProdutos(produtoService))
	produtos.GET("/:id", getProduto(produtoService))
	produtos.POST("", createProduto(produtoService))
	produtos.PUT("/:id", updateProduto(produtoService))
	produtos.DELETE("/:id", deleteProduto(produtoService))

	// Rotas de Vendas
	vendas := v1.Group("/vendas")
	vendas.GET("", getVendas)
	vendas.POST("", createVenda)

	// Documentação Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Inicia o servidor
	e.Logger.Fatal(e.Start(":8080"))
}

// @Summary Lista todos os produtos
// @Description Retorna a lista completa de produtos
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {array} domain.Produto
// @Router /produtos [get]
func getProdutos(service *service.ProdutoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		produtos, err := service.ListProdutos()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, produtos)
	}
}

// @Summary Obtém um produto específico
// @Description Retorna um produto pelo ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do Produto"
// @Success 200 {object} domain.Produto
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [get]
func getProduto(service *service.ProdutoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		}
		produto, err := service.GetProduto(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
		}
		return c.JSON(http.StatusOK, produto)
	}
}

// @Summary Cria um novo produto
// @Description Cria um novo produto no sistema
// @Tags produtos
// @Accept json
// @Produce json
// @Param produto body domain.CreateProdutoDTO true "Produto a ser criado"
// @Success 201 {object} domain.Produto
// @Failure 400 {object} map[string]string
// @Router /produtos [post]
func createProduto(service *service.ProdutoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := new(domain.CreateProdutoDTO)
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		produto := &domain.Produto{
			Nome:       dto.Nome,
			Descricao:  dto.Descricao,
			Preco:      dto.Preco,
			Quantidade: dto.Quantidade,
		}

		if err := service.CreateProduto(produto); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, produto)
	}
}

// @Summary Atualiza um produto
// @Description Atualiza um produto existente
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do Produto"
// @Param produto body domain.UpdateProdutoDTO true "Dados do produto a serem atualizados"
// @Success 200 {object} domain.Produto
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [put]
func updateProduto(service *service.ProdutoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		}

		dto := new(domain.UpdateProdutoDTO)
		if err := c.Bind(dto); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		produto := &domain.Produto{
			ID:         id,
			Nome:       dto.Nome,
			Descricao:  dto.Descricao,
			Preco:      dto.Preco,
			Quantidade: dto.Quantidade,
		}

		if err := service.UpdateProduto(produto); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, produto)
	}
}

// @Summary Remove um produto
// @Description Remove um produto do sistema
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do Produto"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [delete]
func deleteProduto(service *service.ProdutoService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		}
		if err := service.DeleteProduto(id); err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
		}
		return c.NoContent(http.StatusNoContent)
	}
}

// Handlers temporários para vendas
func getVendas(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Lista de Vendas"})
}

func createVenda(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{"message": "Venda criada"})
}
