# Sistema de Vendas em Go

Este é um projeto de estudo em Go que implementa um sistema simples de CRUD (Create, Read, Update, Delete) para registro de vendas.

## Estrutura do Projeto

```
vendas/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação
├── internal/
│   ├── domain/          # Entidades e regras de negócio
│   ├── repository/      # Camada de acesso a dados
│   └── service/         # Lógica de negócio
├── pkg/
│   └── utils/           # Utilitários compartilhados
└── go.mod              # Arquivo de dependências
```

## Módulos

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

## Como Executar

1. Certifique-se de ter o Go instalado
2. Clone o repositório
3. Execute `go mod tidy` para baixar as dependências
4. Execute `go run cmd/main.go` para iniciar o servidor

## Funcionalidades

- Cadastro de produtos
- Registro de vendas
- Cálculo de totais
- Relatórios básicos
