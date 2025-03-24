package repository

import (
	"database/sql"
	"vendas/internal/domain"
	"vendas/internal/utils"
)

type ProdutoRepository interface {
	Create(produto *domain.Produto) error
	GetByID(id string) (*domain.Produto, error)
	GetAll() ([]domain.Produto, error)
	Update(produto *domain.Produto) error
	Delete(id string) error
}

type ProdutoRepositoryImpl struct {
	db *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepositoryImpl {
	return &ProdutoRepositoryImpl{db: db}
}

func (r *ProdutoRepositoryImpl) Create(produto *domain.Produto) error {
	// Gera UUID para o produto
	produto.ID = utils.GenerateUUID()

	query := `INSERT INTO produtos (id, nome, descricao, preco, quantidade, data_criacao) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, produto.ID, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.DataCriacao)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProdutoRepositoryImpl) GetByID(id string) (*domain.Produto, error) {
	produto := &domain.Produto{}
	query := `SELECT id, nome, descricao, preco, quantidade, data_criacao FROM produtos WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade, &produto.DataCriacao)
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (r *ProdutoRepositoryImpl) GetAll() ([]domain.Produto, error) {
	query := `SELECT id, nome, descricao, preco, quantidade, data_criacao FROM produtos`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []domain.Produto
	for rows.Next() {
		var produto domain.Produto
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade, &produto.DataCriacao)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}
	return produtos, nil
}

func (r *ProdutoRepositoryImpl) Update(produto *domain.Produto) error {
	query := `UPDATE produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?`
	_, err := r.db.Exec(query, produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.ID)
	return err
}

func (r *ProdutoRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM produtos WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
