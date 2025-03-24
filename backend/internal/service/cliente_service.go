package service

import (
	"errors"
	"time"
	"vendas/internal/domain"
	"vendas/internal/repository"
)

type ClienteService struct {
	repo *repository.ClienteRepository
}

func NewClienteService(repo *repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) CreateCliente(cliente *domain.Cliente) error {
	if cliente.Nome == "" {
		return errors.New("nome do cliente é obrigatório")
	}
	if cliente.Email == "" {
		return errors.New("email do cliente é obrigatório")
	}
	if cliente.CPF == "" {
		return errors.New("CPF do cliente é obrigatório")
	}

	// Define a data de criação automaticamente
	cliente.DataCriacao = time.Now()

	// O ID será definido pelo repositório
	return s.repo.Create(cliente)
}

func (s *ClienteService) GetCliente(id string) (*domain.Cliente, error) {
	return s.repo.GetByID(id)
}

func (s *ClienteService) GetClienteByCPF(cpf string) (*domain.Cliente, error) {
	return s.repo.GetByCPF(cpf)
}

func (s *ClienteService) ListClientes() ([]domain.Cliente, error) {
	return s.repo.GetAll()
}

func (s *ClienteService) UpdateCliente(cliente *domain.Cliente) error {
	if cliente.ID == "" {
		return errors.New("id do cliente é obrigatório")
	}
	if cliente.Nome == "" {
		return errors.New("nome do cliente é obrigatório")
	}
	if cliente.Email == "" {
		return errors.New("email do cliente é obrigatório")
	}
	if cliente.CPF == "" {
		return errors.New("CPF do cliente é obrigatório")
	}

	// Busca o cliente existente para manter a data de criação original
	clienteExistente, err := s.repo.GetByID(cliente.ID)
	if err != nil {
		return err
	}

	// Mantém a data de criação original
	cliente.DataCriacao = clienteExistente.DataCriacao

	return s.repo.Update(cliente)
}

func (s *ClienteService) DeleteCliente(id string) error {
	return s.repo.Delete(id)
}
