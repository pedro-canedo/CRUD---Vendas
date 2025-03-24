package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"vendas/docs"
	"vendas/internal/database"
	"vendas/internal/domain"
	"vendas/internal/repository"
	"vendas/internal/service"
	"vendas/internal/web"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sistema de Vendas API
// @version 1.0
// @description API para gerenciamento de vendas e produtos
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @contact.name Pedro
// @contact.email pedro@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @tag.name produtos
// @tag.description Operações relacionadas a produtos
// @tag.name vendas
// @tag.description Operações relacionadas a vendas
func main() {
	// Inicializa o banco de dados
	if err := database.InitDB(); err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}

	// Carrega produtos iniciais
	if err := loadInitialProducts(); err != nil {
		log.Printf("Erro ao carregar produtos iniciais: %v", err)
	}

	// Inicializa os repositories
	produtoRepo := repository.NewProdutoRepository(database.DB)
	vendaRepo := repository.NewVendaRepository(database.DB)

	// Inicializa os services
	produtoService := service.NewProdutoService(produtoRepo)
	vendaService := service.NewVendaService(vendaRepo, produtoRepo)

	// Inicializa o router
	router := gin.Default()

	// Configura o Swagger
	docs.SwaggerInfo.Title = "API de Vendas"
	docs.SwaggerInfo.Description = "API para gerenciamento de vendas e produtos"
	docs.SwaggerInfo.Version = "1.0"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configura as rotas
	web.SetupRoutes(router, produtoService, vendaService)

	// Inicia o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func loadInitialProducts() error {
	file, err := os.ReadFile("productsCreate.json")
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo de produtos: %v", err)
	}

	var produtos []domain.Produto
	if err := json.Unmarshal(file, &produtos); err != nil {
		return fmt.Errorf("erro ao decodificar produtos: %v", err)
	}

	produtoRepo := repository.NewProdutoRepository(database.DB)
	for _, produto := range produtos {
		produto.DataCriacao = time.Now()
		if err := produtoRepo.Create(&produto); err != nil {
			return fmt.Errorf("erro ao criar produto %s: %v", produto.Nome, err)
		}
	}

	return nil
}
