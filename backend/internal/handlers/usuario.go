package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"vendas/internal/domain"
	"vendas/internal/dto"
	"vendas/internal/utils"
)

type UsuarioHandler struct {
	usuarioService domain.UsuarioService
	jwtSecretKey   string
}

func NewUsuarioHandler(service domain.UsuarioService, jwtSecretKey string) *UsuarioHandler {
	return &UsuarioHandler{
		usuarioService: service,
		jwtSecretKey:   jwtSecretKey,
	}
}

func (h *UsuarioHandler) CreateUsuario(c *gin.Context) {
	var dto dto.CreateUsuarioDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se já existe um usuário com o mesmo email
	existingUser, _ := h.usuarioService.GetUsuarioByEmail(dto.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email já cadastrado"})
		return
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao processar senha"})
		return
	}

	usuario := &domain.Usuario{
		ID:          uuid.New().String(),
		Nome:        dto.Nome,
		Email:       dto.Email,
		Senha:       string(hashedPassword),
		Role:        dto.Role,
		Ativo:       true,
		DataCriacao: time.Now(),
	}

	if err := h.usuarioService.CreateUsuario(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

func (h *UsuarioHandler) Login(c *gin.Context) {
	var dto dto.LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, err := h.usuarioService.Autenticar(dto.Email, dto.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}

	// Gera o token JWT
	token, err := utils.GenerateToken(usuario.ID, string(usuario.Role), h.jwtSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login realizado com sucesso",
		"token":   token,
		"usuario": usuario,
	})
}

func (h *UsuarioHandler) GetUsuario(c *gin.Context) {
	id := c.Param("id")
	usuario, err := h.usuarioService.GetUsuario(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) ListUsuarios(c *gin.Context) {
	usuarios, err := h.usuarioService.ListUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (h *UsuarioHandler) UpdateUsuario(c *gin.Context) {
	id := c.Param("id")
	var dto dto.UpdateUsuarioDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, err := h.usuarioService.GetUsuario(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	// Atualiza apenas os campos fornecidos
	if dto.Nome != "" {
		usuario.Nome = dto.Nome
	}
	if dto.Email != "" {
		usuario.Email = dto.Email
	}
	if dto.Senha != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Senha), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao processar senha"})
			return
		}
		usuario.Senha = string(hashedPassword)
	}
	if dto.Role != "" {
		usuario.Role = dto.Role
	}
	if dto.Ativo != nil {
		usuario.Ativo = *dto.Ativo
	}

	if err := h.usuarioService.UpdateUsuario(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (h *UsuarioHandler) DeleteUsuario(c *gin.Context) {
	id := c.Param("id")
	if err := h.usuarioService.DeleteUsuario(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "usuário deletado com sucesso"})
}

func (h *UsuarioHandler) GetUsuarioAtual(c *gin.Context) {
	userID := c.GetString("usuario_id")
	usuario, err := h.usuarioService.GetUsuario(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
