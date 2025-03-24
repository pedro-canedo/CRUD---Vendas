package service

import (
	"errors"
	"fmt"
	"time"
	"vendas/internal/domain"
	"vendas/internal/repository"
)

type VendaService struct {
	vendaRepo   repository.VendaRepository
	produtoRepo repository.ProdutoRepository
}

func NewVendaService(vendaRepo repository.VendaRepository, produtoRepo repository.ProdutoRepository) *VendaService {
	return &VendaService{
		vendaRepo:   vendaRepo,
		produtoRepo: produtoRepo,
	}
}

func (s *VendaService) GetAll() ([]domain.Venda, error) {
	return s.vendaRepo.GetAll()
}

func (s *VendaService) GetByID(id string) (*domain.Venda, error) {
	return s.vendaRepo.GetByID(id)
}

func (s *VendaService) Create(venda *domain.Venda) error {
	if venda.ClienteID == "" {
		return errors.New("cliente é obrigatório")
	}
	if len(venda.Items) == 0 {
		return errors.New("venda deve ter pelo menos um item")
	}

	// Define a data da venda como o momento atual
	venda.DataVenda = time.Now()

	var subtotal float64
	// Validar disponibilidade de estoque e calcular valores
	for i := range venda.Items {
		if venda.Items[i].ProdutoID == "" {
			return errors.New("id do produto é obrigatório")
		}
		if venda.Items[i].Quantidade <= 0 {
			return errors.New("quantidade deve ser maior que zero")
		}

		produto, err := s.produtoRepo.GetByID(venda.Items[i].ProdutoID)
		if err != nil {
			return fmt.Errorf("erro ao buscar produto %s: %v", venda.Items[i].ProdutoID, err)
		}

		if produto == nil {
			return fmt.Errorf("produto %s não encontrado", venda.Items[i].ProdutoID)
		}

		// Validar estoque disponível
		if produto.Quantidade < venda.Items[i].Quantidade {
			return fmt.Errorf("estoque insuficiente para o produto %s. Disponível: %d, Solicitado: %d",
				produto.Nome, produto.Quantidade, venda.Items[i].Quantidade)
		}

		// Define o preço unitário
		venda.Items[i].PrecoUnitario = produto.Preco
		subtotal += float64(venda.Items[i].Quantidade) * venda.Items[i].PrecoUnitario
	}

	// Calcula o total final
	venda.ValorTotal = subtotal

	// Cria a venda em uma transação
	return s.vendaRepo.Create(venda)
}

func (s *VendaService) Update(venda *domain.Venda) error {
	if venda.ID == "" {
		return errors.New("id da venda é obrigatório")
	}
	if venda.ClienteID == "" {
		return errors.New("cliente é obrigatório")
	}
	if len(venda.Items) == 0 {
		return errors.New("venda deve ter pelo menos um item")
	}

	var total float64
	for _, item := range venda.Items {
		if item.ProdutoID == "" {
			return errors.New("id do produto é obrigatório")
		}
		if item.Quantidade <= 0 {
			return errors.New("quantidade deve ser maior que zero")
		}

		produto, err := s.produtoRepo.GetByID(item.ProdutoID)
		if err != nil {
			return err
		}

		if produto.Quantidade < item.Quantidade {
			return errors.New("quantidade insuficiente em estoque")
		}

		item.PrecoUnitario = produto.Preco
		total += float64(item.Quantidade) * item.PrecoUnitario
	}

	venda.ValorTotal = total
	return s.vendaRepo.Update(venda)
}

func (s *VendaService) Delete(id string) error {
	if id == "" {
		return errors.New("id da venda é obrigatório")
	}

	return s.vendaRepo.Delete(id)
}

func (s *VendaService) GetVendasPorCliente(cliente string) ([]domain.Venda, error) {
	if cliente == "" {
		return nil, errors.New("cliente é obrigatório")
	}

	return s.vendaRepo.GetVendasPorCliente(cliente)
}

func (s *VendaService) GetVendasPorPeriodo(inicio, fim int64) ([]domain.Venda, error) {
	if inicio == 0 || fim == 0 {
		return nil, errors.New("período é obrigatório")
	}
	if inicio > fim {
		return nil, errors.New("data inicial deve ser menor que a data final")
	}

	return s.vendaRepo.GetVendasPorPeriodo(inicio, fim)
}
