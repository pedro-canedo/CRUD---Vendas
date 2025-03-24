package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "vendas.db")
	if err != nil {
		return err
	}

	// Cria as tabelas
	if err := createTables(); err != nil {
		return err
	}

	return nil
}

func createTables() error {
	// Cria a tabela de usuários
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS usuarios (
			id TEXT PRIMARY KEY,
			nome TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			senha TEXT NOT NULL,
			role TEXT NOT NULL,
			ativo BOOLEAN NOT NULL DEFAULT true,
			data_criacao DATETIME NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Cria a tabela de produtos
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS produtos (
			id TEXT PRIMARY KEY,
			nome TEXT NOT NULL,
			descricao TEXT,
			preco REAL NOT NULL,
			quantidade INTEGER NOT NULL,
			imagem_url TEXT,
			data_criacao DATETIME NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Cria a tabela de vendas
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS vendas (
			id TEXT PRIMARY KEY,
			cliente_id TEXT NOT NULL,
			vendedor_id TEXT NOT NULL,
			data_venda DATETIME NOT NULL,
			valor_total REAL NOT NULL,
			data_criacao DATETIME NOT NULL,
			FOREIGN KEY (cliente_id) REFERENCES usuarios(id),
			FOREIGN KEY (vendedor_id) REFERENCES usuarios(id)
		)
	`)
	if err != nil {
		return err
	}

	// Cria a tabela de itens de venda
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS itens_venda (
			id TEXT PRIMARY KEY,
			venda_id TEXT NOT NULL,
			produto_id TEXT NOT NULL,
			quantidade INTEGER NOT NULL,
			preco_unitario REAL NOT NULL,
			FOREIGN KEY (venda_id) REFERENCES vendas(id),
			FOREIGN KEY (produto_id) REFERENCES produtos(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
