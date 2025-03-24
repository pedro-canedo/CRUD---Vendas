# Criando Novas Rotas na API de Vendas

Este documento descreve o passo a passo para criar uma nova rota na API de Vendas, seguindo a arquitetura existente.

## Estrutura do Projeto

```
vendas/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── domain/          # Definições de domínio e DTOs
│   ├── repository/      # Camada de acesso a dados
│   ├── service/         # Lógica de negócio
│   └── web/            # Handlers e rotas
└── docs/               # Documentação
```

## Passo a Passo

### 1. Definir o Domínio

Primeiro, crie ou atualize as estruturas no pacote `domain`:

```go
// internal/domain/seu_dominio.go

type SeuDominio struct {
    ID          int64     `json:"id"`
    Nome        string    `json:"nome"`
    // ... outros campos
}

// DTO para criação
type CreateSeuDominioDTO struct {
    Nome string `json:"nome" validate:"required"`
    // ... outros campos
}

// DTO para atualização
type UpdateSeuDominioDTO struct {
    Nome string `json:"nome" validate:"required"`
    // ... outros campos
}
```

### 2. Criar o Repositório

Crie o repositório no pacote `repository`:

```go
// internal/repository/seu_repositorio.go

type SeuRepositorio interface {
    Create(dominio *domain.SeuDominio) error
    GetByID(id int64) (*domain.SeuDominio, error)
    GetAll() ([]domain.SeuDominio, error)
    Update(dominio *domain.SeuDominio) error
    Delete(id int64) error
}

type SeuRepositorioImpl struct {
    db *sql.DB
}

func NewSeuRepositorio(db *sql.DB) *SeuRepositorioImpl {
    return &SeuRepositorioImpl{db: db}
}

// Implemente os métodos da interface
```

### 3. Criar o Serviço

Crie o serviço no pacote `service`:

```go
// internal/service/seu_servico.go

type SeuServico struct {
    repo repository.SeuRepositorio
}

func NewSeuServico(repo repository.SeuRepositorio) *SeuServico {
    return &SeuServico{repo: repo}
}

// Implemente os métodos do serviço
```

### 4. Criar o Handler

Adicione o handler no arquivo `internal/web/routes.go`:

```go
// @Summary Descrição da operação
// @Description Descrição detalhada da operação
// @Tags seu-tag
// @Accept json
// @Produce json
// @Param seu_parametro body domain.CreateSeuDominioDTO true "Descrição do parâmetro"
// @Success 201 {object} domain.SeuDominio
// @Failure 400 {object} map[string]string
// @Router /seu-endpoint [post]
func seuHandler(service *service.SeuServico) gin.HandlerFunc {
    return func(c *gin.Context) {
        var dto domain.CreateSeuDominioDTO
        if err := c.ShouldBindJSON(&dto); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Lógica do handler
        if err := service.Create(&dto); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, dto)
    }
}
```

### 5. Registrar a Rota

Adicione a rota no método `SetupRoutes` em `internal/web/routes.go`:

```go
func SetupRoutes(router *gin.Engine, seuServico *service.SeuServico) {
    api := router.Group("/api/v1")
    {
        seuGrupo := api.Group("/seu-endpoint")
        {
            seuGrupo.POST("", seuHandler(seuServico))
            seuGrupo.GET("/:id", getSeuDominio(seuServico))
            // ... outras rotas
        }
    }
}
```

### 6. Inicializar no Main

Atualize o `cmd/main.go` para incluir o novo serviço:

```go
func main() {
    // ... código existente ...

    seuRepo := repository.NewSeuRepositorio(database.DB)
    seuServico := service.NewSeuServico(seuRepo)

    // ... código existente ...

    web.SetupRoutes(router, seuServico)
}
```

### 7. Documentação Swagger

As anotações Swagger devem ser adicionadas em todos os handlers:

```go
// @Summary Título da operação
// @Description Descrição detalhada
// @Tags seu-tag
// @Accept json
// @Produce json
// @Param id path int true "ID do recurso"
// @Success 200 {object} domain.SeuDominio
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /seu-endpoint/{id} [get]
```

### 8. Exemplo Completo

Aqui está um exemplo completo de como criar uma rota para gerenciar categorias:

```go
// internal/domain/categoria.go
type Categoria struct {
    ID   int64  `json:"id"`
    Nome string `json:"nome"`
}

type CreateCategoriaDTO struct {
    Nome string `json:"nome" validate:"required"`
}

// internal/repository/categoria_repository.go
type CategoriaRepository interface {
    Create(categoria *domain.Categoria) error
    GetByID(id int64) (*domain.Categoria, error)
    GetAll() ([]domain.Categoria, error)
    Update(categoria *domain.Categoria) error
    Delete(id int64) error
}

// internal/service/categoria_service.go
type CategoriaService struct {
    repo repository.CategoriaRepository
}

// internal/web/routes.go
// @Summary Cria uma nova categoria
// @Description Cria uma nova categoria com o nome fornecido
// @Tags categorias
// @Accept json
// @Produce json
// @Param categoria body domain.CreateCategoriaDTO true "Dados da categoria"
// @Success 201 {object} domain.Categoria
// @Failure 400 {object} map[string]string
// @Router /categorias [post]
func createCategoria(service *service.CategoriaService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var dto domain.CreateCategoriaDTO
        if err := c.ShouldBindJSON(&dto); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        categoria := &domain.Categoria{
            Nome: dto.Nome,
        }

        if err := service.Create(categoria); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, categoria)
    }
}
```

## Boas Práticas

1. **Validação**: Use tags `validate` nos DTOs para validação de dados
2. **Tratamento de Erros**: Retorne erros apropriados com mensagens claras
3. **Documentação**: Mantenha a documentação Swagger atualizada
4. **Testes**: Crie testes unitários para cada camada
5. **Consistência**: Siga o mesmo padrão de nomenclatura e estrutura
6. **Segurança**: Implemente autenticação e autorização quando necessário
7. **Logging**: Adicione logs apropriados para monitoramento
