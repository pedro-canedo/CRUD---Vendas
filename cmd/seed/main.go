package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
	"vendas/internal/database"
	"vendas/internal/utils"
)

type ProdutoInput struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func main() {
	// Inicializa o banco de dados
	database.InitDB()

	// Lê o arquivo JSON
	data, err := ioutil.ReadFile("productsCreate.json")
	if err != nil {
		log.Fatal("Erro ao ler arquivo JSON:", err)
	}

	// Decodifica o JSON
	var produtos []ProdutoInput
	if err := json.Unmarshal(data, &produtos); err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	// Insere os produtos no banco
	for _, p := range produtos {
		id := utils.GenerateUUID()
		_, err := database.DB.Exec(`
			INSERT INTO produtos (id, nome, descricao, preco, quantidade, data_criacao)
			VALUES (?, ?, ?, ?, ?, ?)
		`, id, p.Nome, p.Descricao, p.Preco, p.Quantidade, time.Now())

		if err != nil {
			log.Printf("Erro ao inserir produto %s: %v", p.Nome, err)
		} else {
			log.Printf("Produto %s inserido com sucesso (ID: %s)", p.Nome, id)
		}
	}

	log.Println("Processo de seed concluído!")
}
