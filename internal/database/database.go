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
	// Cria a tabela de produtos
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS produtos (
			id TEXT PRIMARY KEY,
			nome TEXT NOT NULL,
			descricao TEXT,
			preco REAL NOT NULL,
			quantidade INTEGER NOT NULL,
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
			cliente TEXT NOT NULL,
			data_venda DATETIME NOT NULL,
			total REAL NOT NULL
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
			subtotal REAL NOT NULL,
			FOREIGN KEY (venda_id) REFERENCES vendas(id),
			FOREIGN KEY (produto_id) REFERENCES produtos(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
