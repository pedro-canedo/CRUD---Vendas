package service

import (
	"errors"
	"time"
	"vendas/internal/domain"
	"vendas/internal/repository"
)

type ProdutoService struct {
	repo repository.ProdutoRepository
}

func NewProdutoService(repo repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) GetAll() ([]domain.Produto, error) {
	return s.repo.GetAll()
}

func (s *ProdutoService) GetByID(id string) (*domain.Produto, error) {
	return s.repo.GetByID(id)
}

func (s *ProdutoService) Create(produto *domain.Produto) error {
	if produto.Nome == "" {
		return errors.New("nome do produto é obrigatório")
	}
	if produto.Preco <= 0 {
		return errors.New("preço do produto deve ser maior que zero")
	}
	if produto.Quantidade < 0 {
		return errors.New("quantidade do produto não pode ser negativa")
	}

	return s.repo.Create(produto)
}

func (s *ProdutoService) Update(produto *domain.Produto) error {
	if produto.ID == "" {
		return errors.New("id do produto é obrigatório")
	}
	if produto.Nome == "" {
		return errors.New("nome do produto é obrigatório")
	}
	if produto.Preco <= 0 {
		return errors.New("preço do produto deve ser maior que zero")
	}
	if produto.Quantidade < 0 {
		return errors.New("quantidade do produto não pode ser negativa")
	}

	return s.repo.Update(produto)
}

func (s *ProdutoService) Delete(id string) error {
	if id == "" {
		return errors.New("id do produto é obrigatório")
	}

	return s.repo.Delete(id)
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

func (s *ProdutoService) DeleteProduto(id string) error {
	return s.repo.Delete(id)
}
