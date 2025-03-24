package web

import (
	"net/http"
	"strconv"
	"vendas/internal/domain"
	"vendas/internal/service"

	"github.com/gin-gonic/gin"
)

// @Summary Lista todos os produtos
// @Description Retorna uma lista de todos os produtos cadastrados
// @Tags produtos
// @Accept json
// @Produce json
// @Success 200 {array} domain.Produto
// @Router /produtos [get]
func getProdutos(service *service.ProdutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		produtos, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, produtos)
	}
}

// @Summary Obtém um produto por ID
// @Description Retorna um produto específico pelo seu ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Success 200 {object} domain.Produto
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [get]
func getProduto(service *service.ProdutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		produto, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, produto)
	}
}

// @Summary Cria um novo produto
// @Description Cria um novo produto com os dados fornecidos
// @Tags produtos
// @Accept json
// @Produce json
// @Param produto body domain.Produto true "Dados do produto"
// @Success 201 {object} domain.Produto
// @Failure 400 {object} map[string]string
// @Router /produtos [post]
func createProduto(service *service.ProdutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produto domain.Produto
		if err := c.ShouldBindJSON(&produto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.Create(&produto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, produto)
	}
}

// @Summary Atualiza um produto
// @Description Atualiza um produto existente com os dados fornecidos
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Param produto body domain.Produto true "Dados do produto"
// @Success 200 {object} domain.Produto
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [put]
func updateProduto(service *service.ProdutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		var produto domain.Produto
		if err := c.ShouldBindJSON(&produto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		produto.ID = id
		if err := service.Update(&produto); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, produto)
	}
}

// @Summary Remove um produto
// @Description Remove um produto pelo seu ID
// @Tags produtos
// @Accept json
// @Produce json
// @Param id path string true "ID do produto"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /produtos/{id} [delete]
func deleteProduto(service *service.ProdutoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		if err := service.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

// @Summary Lista todas as vendas
// @Description Retorna uma lista de todas as vendas cadastradas
// @Tags vendas
// @Accept json
// @Produce json
// @Success 200 {array} domain.Venda
// @Router /vendas [get]
func getVendas(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		vendas, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, vendas)
	}
}

// @Summary Obtém uma venda por ID
// @Description Retorna uma venda específica pelo seu ID
// @Tags vendas
// @Accept json
// @Produce json
// @Param id path string true "ID da venda"
// @Success 200 {object} domain.Venda
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /vendas/{id} [get]
func getVenda(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		venda, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, venda)
	}
}

// @Summary Cria uma nova venda
// @Description Cria uma nova venda com os dados fornecidos
// @Tags vendas
// @Accept json
// @Produce json
// @Param venda body domain.CreateVendaDTO true "Dados da venda"
// @Success 201 {object} domain.Venda
// @Failure 400 {object} map[string]string
// @Router /vendas [post]
func createVenda(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto domain.CreateVendaDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Converte os itens do DTO para a entidade ItemVenda
		itens := make([]domain.ItemVenda, len(dto.Itens))
		for i, itemDTO := range dto.Itens {
			itens[i] = domain.ItemVenda{
				ProdutoID:  itemDTO.ProdutoID,
				Quantidade: itemDTO.Quantidade,
			}
		}

		// Cria a entidade Venda
		venda := &domain.Venda{
			Cliente:  dto.Cliente,
			Itens:    itens,
			Desconto: dto.Desconto,
		}

		if err := service.Create(venda); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, venda)
	}
}

// @Summary Atualiza uma venda
// @Description Atualiza uma venda existente com os dados fornecidos
// @Tags vendas
// @Accept json
// @Produce json
// @Param id path string true "ID da venda"
// @Param venda body domain.Venda true "Dados da venda"
// @Success 200 {object} domain.Venda
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /vendas/{id} [put]
func updateVenda(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		var venda domain.Venda
		if err := c.ShouldBindJSON(&venda); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		venda.ID = id
		if err := service.Update(&venda); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, venda)
	}
}

// @Summary Remove uma venda
// @Description Remove uma venda pelo seu ID
// @Tags vendas
// @Accept json
// @Produce json
// @Param id path string true "ID da venda"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /vendas/{id} [delete]
func deleteVenda(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}

		if err := service.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

// @Summary Lista vendas por cliente
// @Description Retorna uma lista de vendas filtrada por cliente
// @Tags vendas
// @Accept json
// @Produce json
// @Param cliente path string true "Nome do cliente"
// @Success 200 {array} domain.Venda
// @Failure 400 {object} map[string]string
// @Router /vendas/cliente/{cliente} [get]
func getVendasPorCliente(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		cliente := c.Param("cliente")
		vendas, err := service.GetVendasPorCliente(cliente)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, vendas)
	}
}

// @Summary Lista vendas por período
// @Description Retorna uma lista de vendas filtrada por período
// @Tags vendas
// @Accept json
// @Produce json
// @Param inicio path int true "Data inicial (timestamp)"
// @Param fim path int true "Data final (timestamp)"
// @Success 200 {array} domain.Venda
// @Failure 400 {object} map[string]string
// @Router /vendas/periodo/{inicio}/{fim} [get]
func getVendasPorPeriodo(service *service.VendaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		inicio, err := strconv.ParseInt(c.Param("inicio"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data inicial inválida"})
			return
		}

		fim, err := strconv.ParseInt(c.Param("fim"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "data final inválida"})
			return
		}

		vendas, err := service.GetVendasPorPeriodo(inicio, fim)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, vendas)
	}
}

func SetupRoutes(router *gin.Engine, produtoService *service.ProdutoService, vendaService *service.VendaService) {
	api := router.Group("/api/v1")
	{
		// Rotas de produtos
		produtos := api.Group("/produtos")
		{
			produtos.GET("", getProdutos(produtoService))
			produtos.GET("/:id", getProduto(produtoService))
			produtos.POST("", createProduto(produtoService))
			produtos.PUT("/:id", updateProduto(produtoService))
			produtos.DELETE("/:id", deleteProduto(produtoService))
		}

		// Rotas de vendas
		vendas := api.Group("/vendas")
		{
			vendas.GET("", getVendas(vendaService))
			vendas.GET("/:id", getVenda(vendaService))
			vendas.POST("", createVenda(vendaService))
			vendas.PUT("/:id", updateVenda(vendaService))
			vendas.DELETE("/:id", deleteVenda(vendaService))
			vendas.GET("/cliente/:cliente", getVendasPorCliente(vendaService))
			vendas.GET("/periodo/:inicio/:fim", getVendasPorPeriodo(vendaService))
		}
	}
}
