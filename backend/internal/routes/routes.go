package routes

import (
	"github.com/gin-gonic/gin"

	"vendas/internal/handlers"
	"vendas/internal/middleware"
)

func SetupRoutes(r *gin.Engine, handlers *handlers.Handlers, jwtSecretKey string) {
	// Rotas públicas
	public := r.Group("/api")
	{
		public.POST("/login", handlers.Usuario.Login)
	}

	// Rotas protegidas
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware(jwtSecretKey))
	{
		// Rotas de usuário
		usuarios := authorized.Group("/usuarios")
		{
			// Rota para obter dados do usuário atual (acessível para qualquer usuário autenticado)
			usuarios.GET("/me", handlers.Usuario.GetUsuarioAtual)

			// Rotas que requerem role admin
			adminUsuarios := usuarios.Group("")
			adminUsuarios.Use(middleware.RequireRole("admin"))
			{
				adminUsuarios.POST("/", handlers.Usuario.CreateUsuario)
				adminUsuarios.GET("/", handlers.Usuario.ListUsuarios)
				adminUsuarios.GET("/:id", handlers.Usuario.GetUsuario)
				adminUsuarios.PUT("/:id", handlers.Usuario.UpdateUsuario)
				adminUsuarios.DELETE("/:id", handlers.Usuario.DeleteUsuario)
			}
		}

		// Rotas de cliente
		clientes := authorized.Group("/clientes")
		clientes.Use(middleware.RequireRole("admin", "vendedor"))
		{
			clientes.POST("/", handlers.Cliente.CreateCliente)
			clientes.GET("/", handlers.Cliente.ListClientes)
			clientes.GET("/:id", handlers.Cliente.GetCliente)
			clientes.PUT("/:id", handlers.Cliente.UpdateCliente)
			clientes.DELETE("/:id", handlers.Cliente.DeleteCliente)
		}

		// Rotas de produto
		produtos := authorized.Group("/produtos")
		produtos.Use(middleware.RequireRole("admin", "vendedor"))
		{
			produtos.POST("/", handlers.Produto.CreateProduto)
			produtos.GET("/", handlers.Produto.ListProdutos)
			produtos.GET("/:id", handlers.Produto.GetProduto)
			produtos.PUT("/:id", handlers.Produto.UpdateProduto)
			produtos.DELETE("/:id", handlers.Produto.DeleteProduto)
		}
	}
}
