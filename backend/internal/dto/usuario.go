package dto

import "vendas/internal/domain"

type CreateUsuarioDTO struct {
	Nome  string      `json:"nome" binding:"required"`
	Email string      `json:"email" binding:"required,email"`
	Senha string      `json:"senha" binding:"required,min=6"`
	Role  domain.Role `json:"role" binding:"required,oneof=admin vendedor cliente"`
}

type UpdateUsuarioDTO struct {
	Nome  string      `json:"nome"`
	Email string      `json:"email" binding:"omitempty,email"`
	Senha string      `json:"senha" binding:"omitempty,min=6"`
	Role  domain.Role `json:"role" binding:"omitempty,oneof=admin vendedor cliente"`
	Ativo *bool       `json:"ativo"`
}

type LoginDTO struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}
