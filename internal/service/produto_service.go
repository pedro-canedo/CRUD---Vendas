package service

import (
	"errors"
	"time"
	"vendas/internal/domain"
)

type ProdutoService struct {
	repo domain.ProdutoRepository
}

func NewProdutoService(repo domain.ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) CreateProduto(produto *domain.Produto) error {
	if produto.Nome == "" {
		return errors.New("nome do produto é obrigatório")
	}
	if produto.Preco <= 0 {
		return errors.New("preço do produto deve ser maior que zero")
	}
	if produto.Quantidade < 0 {
		return errors.New("quantidade não pode ser negativa")
	}

	// Define a data de criação automaticamente
	produto.DataCriacao = time.Now()

	// O ID será definido pelo repositório
	return s.repo.Create(produto)
}

func (s *ProdutoService) GetProduto(id int64) (*domain.Produto, error) {
	return s.repo.GetByID(id)
}

func (s *ProdutoService) ListProdutos() ([]domain.Produto, error) {
	return s.repo.GetAll()
}

func (s *ProdutoService) UpdateProduto(produto *domain.Produto) error {
	if produto.Nome == "" {
		return errors.New("nome do produto é obrigatório")
	}
	if produto.Preco <= 0 {
		return errors.New("preço do produto deve ser maior que zero")
	}
	if produto.Quantidade < 0 {
		return errors.New("quantidade não pode ser negativa")
	}

	// Busca o produto existente para manter a data de criação original
	produtoExistente, err := s.repo.GetByID(produto.ID)
	if err != nil {
		return err
	}

	// Mantém a data de criação original
	produto.DataCriacao = produtoExistente.DataCriacao

	return s.repo.Update(produto)
}

func (s *ProdutoService) DeleteProduto(id int64) error {
	return s.repo.Delete(id)
}
