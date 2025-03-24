package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"vendas/internal/domain"
	"vendas/internal/dto"
)

type ClienteHandler struct {
	clienteService domain.ClienteService
}

func NewClienteHandler(service domain.ClienteService) *ClienteHandler {
	return &ClienteHandler{
		clienteService: service,
	}
}

func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var dto dto.CreateClienteDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se já existe um cliente com o mesmo CPF
	existingCliente, _ := h.clienteService.GetClienteByCPF(dto.CPF)
	if existingCliente != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "CPF já cadastrado"})
		return
	}

	// TODO: Obter o ID do usuário autenticado
	usuarioID := c.GetString("usuario_id")
	if usuarioID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
		return
	}

	cliente := &domain.Cliente{
		ID:          uuid.New().String(),
		Nome:        dto.Nome,
		Email:       dto.Email,
		Telefone:    dto.Telefone,
		Endereco:    dto.Endereco,
		CPF:         dto.CPF,
		UsuarioID:   usuarioID,
		DataCriacao: time.Now(),
	}

	if err := h.clienteService.CreateCliente(cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func (h *ClienteHandler) GetCliente(c *gin.Context) {
	id := c.Param("id")
	cliente, err := h.clienteService.GetCliente(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (h *ClienteHandler) ListClientes(c *gin.Context) {
	clientes, err := h.clienteService.ListClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clientes)
}

func (h *ClienteHandler) UpdateCliente(c *gin.Context) {
	id := c.Param("id")
	var dto dto.UpdateClienteDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente, err := h.clienteService.GetCliente(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cliente não encontrado"})
		return
	}

	// Atualiza apenas os campos fornecidos
	if dto.Nome != "" {
		cliente.Nome = dto.Nome
	}
	if dto.Email != "" {
		cliente.Email = dto.Email
	}
	if dto.Telefone != "" {
		cliente.Telefone = dto.Telefone
	}
	if dto.Endereco != "" {
		cliente.Endereco = dto.Endereco
	}
	if dto.CPF != "" {
		cliente.CPF = dto.CPF
	}

	if err := h.clienteService.UpdateCliente(cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func (h *ClienteHandler) DeleteCliente(c *gin.Context) {
	id := c.Param("id")
	if err := h.clienteService.DeleteCliente(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cliente deletado com sucesso"})
}
