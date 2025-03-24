package repository

import (
	"errors"
	"sync"
	"vendas/internal/domain"
)

type ProdutoRepository struct {
	produtos map[int64]*domain.Produto
	mu       sync.RWMutex
	nextID   int64
}

func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{
		produtos: make(map[int64]*domain.Produto),
		nextID:   1,
	}
}

func (r *ProdutoRepository) Create(produto *domain.Produto) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	produto.ID = r.nextID
	r.produtos[produto.ID] = produto
	r.nextID++
	return nil
}

func (r *ProdutoRepository) GetByID(id int64) (*domain.Produto, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if produto, exists := r.produtos[id]; exists {
		return produto, nil
	}
	return nil, errors.New("produto não encontrado")
}

func (r *ProdutoRepository) GetAll() ([]domain.Produto, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	produtos := make([]domain.Produto, 0, len(r.produtos))
	for _, produto := range r.produtos {
		produtos = append(produtos, *produto)
	}
	return produtos, nil
}

func (r *ProdutoRepository) Update(produto *domain.Produto) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.produtos[produto.ID]; !exists {
		return errors.New("produto não encontrado")
	}
	r.produtos[produto.ID] = produto
	return nil
}

func (r *ProdutoRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.produtos[id]; !exists {
		return errors.New("produto não encontrado")
	}
	delete(r.produtos, id)
	return nil
}
