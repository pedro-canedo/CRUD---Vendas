package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"vendas/internal/database"
	"vendas/internal/domain"
)

type SeedData struct {
	Users []struct {
		Nome  string      `json:"nome"`
		Email string      `json:"email"`
		Senha string      `json:"senha"`
		Role  domain.Role `json:"role"`
	} `json:"users"`
	Clients []struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Telefone string `json:"telefone"`
		Endereco string `json:"endereco"`
		CPF      string `json:"cpf"`
	} `json:"clients"`
	Products []struct {
		Nome       string  `json:"nome"`
		Descricao  string  `json:"descricao"`
		Preco      float64 `json:"preco"`
		Quantidade int     `json:"quantidade"`
		ImagemURL  string  `json:"imagem_url"`
	} `json:"products"`
	Sales []struct {
		ClienteID  string  `json:"cliente_id"`
		VendedorID string  `json:"vendedor_id"`
		DataVenda  string  `json:"data_venda"`
		ValorTotal float64 `json:"valor_total"`
		Items      []struct {
			ProdutoID     string  `json:"produto_id"`
			Quantidade    int     `json:"quantidade"`
			PrecoUnitario float64 `json:"preco_unitario"`
		} `json:"items"`
	} `json:"sales"`
}

func LoadSeedData() (*SeedData, error) {
	// Carrega os dados dos arquivos JSON
	usersData, err := ioutil.ReadFile("usersCreate.json")
	if err != nil {
		return nil, err
	}

	clientsData, err := ioutil.ReadFile("clientsCreate.json")
	if err != nil {
		return nil, err
	}

	productsData, err := ioutil.ReadFile("productsCreate.json")
	if err != nil {
		return nil, err
	}

	salesData, err := ioutil.ReadFile("salesCreate.json")
	if err != nil {
		return nil, err
	}

	var seedData SeedData

	// Parse dos dados JSON
	if err := json.Unmarshal(usersData, &seedData); err != nil {
		return nil, fmt.Errorf("erro ao parsear usuários: %v", err)
	}
	if err := json.Unmarshal(clientsData, &seedData); err != nil {
		return nil, fmt.Errorf("erro ao parsear clientes: %v", err)
	}
	if err := json.Unmarshal(productsData, &seedData); err != nil {
		return nil, fmt.Errorf("erro ao parsear produtos: %v", err)
	}
	if err := json.Unmarshal(salesData, &seedData); err != nil {
		return nil, fmt.Errorf("erro ao parsear vendas: %v", err)
	}

	return &seedData, nil
}

func createAdminUser() error {
	// Verifica se o usuário admin já existe
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email = ?", "admin@vendas.com").Scan(&count)
	if err != nil {
		return fmt.Errorf("erro ao verificar usuário admin: %v", err)
	}

	if count > 0 {
		log.Println("Usuário admin já existe")
		return nil
	}

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar hash da senha: %v", err)
	}

	// Cria o usuário admin
	_, err = database.DB.Exec(
		"INSERT INTO usuarios (id, nome, email, senha, role, ativo, data_criacao) VALUES (?, ?, ?, ?, ?, ?, ?)",
		uuid.New().String(),
		"Administrador",
		"admin@vendas.com",
		string(hashedPassword),
		domain.RoleAdmin,
		true,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("erro ao criar usuário admin: %v", err)
	}

	log.Println("Usuário admin criado com sucesso")
	return nil
}

func main() {
	// Inicializa o banco de dados
	if err := database.InitDB(); err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}

	// Cria o usuário admin se não existir
	if err := createAdminUser(); err != nil {
		log.Printf("Erro ao criar usuário admin: %v", err)
	}

	// Carrega os dados de seed
	seedData, err := LoadSeedData()
	if err != nil {
		log.Fatalf("Erro ao carregar dados de seed: %v", err)
	}

	// Mapa para armazenar os IDs gerados
	userIDs := make(map[string]string)
	productIDs := make(map[string]string)

	// Insere usuários (incluindo vendedores)
	for _, user := range seedData.Users {
		// Verifica se o usuário já existe
		var count int
		err := database.DB.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email = ?", user.Email).Scan(&count)
		if err != nil {
			log.Fatalf("Erro ao verificar usuário existente: %v", err)
		}

		if count > 0 {
			log.Printf("Usuário %s já existe, pulando...", user.Email)
			continue
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Erro ao gerar hash da senha: %v", err)
		}

		usuario := &domain.Usuario{
			ID:          uuid.New().String(),
			Nome:        user.Nome,
			Email:       user.Email,
			Senha:       string(hashedPassword),
			Role:        user.Role,
			Ativo:       true,
			DataCriacao: time.Now(),
		}

		_, err = database.DB.Exec(
			"INSERT INTO usuarios (id, nome, email, senha, role, ativo, data_criacao) VALUES (?, ?, ?, ?, ?, ?, ?)",
			usuario.ID, usuario.Nome, usuario.Email, usuario.Senha, usuario.Role, usuario.Ativo, usuario.DataCriacao,
		)
		if err != nil {
			log.Fatalf("Erro ao inserir usuário: %v", err)
		}

		userIDs[user.Email] = usuario.ID
	}

	// Insere clientes como usuários
	for _, client := range seedData.Clients {
		// Verifica se o cliente já existe como usuário
		var count int
		err := database.DB.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email = ?", client.Email).Scan(&count)
		if err != nil {
			log.Fatalf("Erro ao verificar cliente existente: %v", err)
		}

		if count > 0 {
			log.Printf("Cliente %s já existe como usuário, pulando...", client.Email)
			continue
		}

		// Gera uma senha aleatória para o cliente
		senha := uuid.New().String()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Erro ao gerar hash da senha: %v", err)
		}

		usuario := &domain.Usuario{
			ID:          uuid.New().String(),
			Nome:        client.Nome,
			Email:       client.Email,
			Senha:       string(hashedPassword),
			Role:        domain.RoleCliente,
			Ativo:       true,
			DataCriacao: time.Now(),
		}

		_, err = database.DB.Exec(
			"INSERT INTO usuarios (id, nome, email, senha, role, ativo, data_criacao) VALUES (?, ?, ?, ?, ?, ?, ?)",
			usuario.ID, usuario.Nome, usuario.Email, usuario.Senha, usuario.Role, usuario.Ativo, usuario.DataCriacao,
		)
		if err != nil {
			log.Fatalf("Erro ao inserir cliente: %v", err)
		}

		userIDs[client.Email] = usuario.ID
	}

	// Insere produtos
	for _, product := range seedData.Products {
		produto := &domain.Produto{
			ID:          uuid.New().String(),
			Nome:        product.Nome,
			Descricao:   product.Descricao,
			Preco:       product.Preco,
			Quantidade:  product.Quantidade,
			ImagemURL:   product.ImagemURL,
			DataCriacao: time.Now(),
		}

		_, err = database.DB.Exec(
			"INSERT INTO produtos (id, nome, descricao, preco, quantidade, imagem_url, data_criacao) VALUES (?, ?, ?, ?, ?, ?, ?)",
			produto.ID, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.ImagemURL, produto.DataCriacao,
		)
		if err != nil {
			log.Fatalf("Erro ao inserir produto: %v", err)
		}

		productIDs[product.Nome] = produto.ID
	}

	// Lista de vendedores disponíveis
	var vendedores []string
	for _, id := range userIDs {
		var role domain.Role
		err := database.DB.QueryRow("SELECT role FROM usuarios WHERE id = ?", id).Scan(&role)
		if err != nil {
			log.Fatalf("Erro ao verificar role do usuário: %v", err)
		}
		if role == domain.RoleVendedor {
			vendedores = append(vendedores, id)
		}
	}

	// Lista de clientes disponíveis
	var clientes []string
	for _, id := range userIDs {
		var role domain.Role
		err := database.DB.QueryRow("SELECT role FROM usuarios WHERE id = ?", id).Scan(&role)
		if err != nil {
			log.Fatalf("Erro ao verificar role do usuário: %v", err)
		}
		if role == domain.RoleCliente {
			clientes = append(clientes, id)
		}
	}

	// Insere vendas usando os IDs reais
	for _, sale := range seedData.Sales {
		// Seleciona um cliente aleatório
		clienteID := clientes[0] // Por enquanto, usa o primeiro cliente
		// Seleciona um vendedor aleatório
		vendedorID := vendedores[0] // Por enquanto, usa o primeiro vendedor

		dataVenda, err := time.Parse(time.RFC3339, sale.DataVenda)
		if err != nil {
			log.Fatalf("Erro ao parsear data da venda: %v", err)
		}

		venda := &domain.Venda{
			ID:          uuid.New().String(),
			ClienteID:   clienteID,
			VendedorID:  vendedorID,
			DataVenda:   dataVenda,
			ValorTotal:  sale.ValorTotal,
			DataCriacao: time.Now(),
		}

		// Insere a venda
		_, err = database.DB.Exec(
			"INSERT INTO vendas (id, cliente_id, vendedor_id, data_venda, valor_total, data_criacao) VALUES (?, ?, ?, ?, ?, ?)",
			venda.ID, venda.ClienteID, venda.VendedorID, venda.DataVenda, venda.ValorTotal, venda.DataCriacao,
		)
		if err != nil {
			log.Fatalf("Erro ao inserir venda: %v", err)
		}

		// Insere os itens da venda
		for _, item := range sale.Items {
			itemVenda := domain.ItemVenda{
				ID:            uuid.New().String(),
				VendaID:       venda.ID,
				ProdutoID:     productIDs[item.ProdutoID],
				Quantidade:    item.Quantidade,
				PrecoUnitario: item.PrecoUnitario,
			}

			_, err = database.DB.Exec(
				"INSERT INTO itens_venda (id, venda_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?, ?)",
				itemVenda.ID, itemVenda.VendaID, itemVenda.ProdutoID, itemVenda.Quantidade, itemVenda.PrecoUnitario,
			)
			if err != nil {
				log.Fatalf("Erro ao inserir item de venda: %v", err)
			}
		}
	}

	log.Println("Dados iniciais inseridos com sucesso!")
}
