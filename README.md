# Sistema de Vendas em Go

Este é um projeto de estudo em Go que implementa um sistema simples de CRUD (Create, Read, Update, Delete) para registro de vendas.

## Estrutura do Projeto

```
vendas/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── domain/          # Entidades e regras de negócio
│   │   ├── produto.go   # Definição de produtos
│   │   ├── venda.go     # Definição de vendas
│   │   └── dto.go       # Objetos de transferência de dados
│   ├── repository/      # Camada de acesso a dados
│   └── service/         # Lógica de negócio
├── pkg/
│   └── utils/           # Utilitários compartilhados
└── go.mod              # Arquivo de dependências
```

## Modelos de Dados

### Produto

- ID: Identificador único
- Nome: Nome do produto
- Descrição: Descrição detalhada
- Preço: Valor unitário
- Quantidade: Estoque disponível
- Data de Criação: Data de cadastro

### Venda

- ID: Identificador único
- Data de Venda: Data da transação
- Itens: Lista de produtos vendidos
  - ID do Item
  - ID do Produto
  - Quantidade
  - Preço Unitário
  - Subtotal
- Total: Valor total da venda
- Cliente: Nome do cliente

## Funcionalidades

### Gestão de Produtos

- Cadastro de novos produtos
- Consulta de produtos por ID
- Listagem de todos os produtos
- Atualização de informações
- Remoção de produtos

### Gestão de Vendas

- Registro de novas vendas
- Consulta de vendas por ID
- Listagem de todas as vendas
- Atualização de vendas
- Remoção de vendas
- Cálculo automático de totais e subtotais

## Como Executar

1. Certifique-se de ter o Go instalado (versão 1.16 ou superior)
2. Clone o repositório
3. Execute `go mod tidy` para baixar as dependências
4. Execute `go run cmd/main.go` para iniciar o servidor

## Arquitetura

O projeto segue uma arquitetura limpa (Clean Architecture) com as seguintes camadas:

### Domain

- Contém as estruturas de dados principais
- Define as interfaces e tipos básicos
- Implementa as regras de negócio básicas

### Repository

- Responsável pelo acesso a dados
- Implementa as operações CRUD
- Abstrai a camada de persistência

### Service

- Implementa a lógica de negócio
- Coordena as operações entre diferentes partes do sistema
- Realiza cálculos e validações

### Utils

- Funções auxiliares
- Constantes
- Helpers compartilhados

## Tecnologias Utilizadas

- Go 1.16+
- Gin (Framework Web)
- GORM (ORM)
- PostgreSQL (Banco de Dados)
