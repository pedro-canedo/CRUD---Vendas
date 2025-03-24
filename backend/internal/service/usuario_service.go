package service

import (
	"errors"
	"time"
	"vendas/internal/domain"
	"vendas/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) CreateUsuario(usuario *domain.Usuario) error {
	if usuario.Nome == "" {
		return errors.New("nome do usuário é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("email do usuário é obrigatório")
	}
	if usuario.Senha == "" {
		return errors.New("senha do usuário é obrigatória")
	}
	if usuario.Role == "" {
		return errors.New("role do usuário é obrigatória")
	}

	// Define a data de criação automaticamente
	usuario.DataCriacao = time.Now()

	// O ID será definido pelo repositório
	return s.repo.Create(usuario)
}

func (s *UsuarioService) GetUsuario(id string) (*domain.Usuario, error) {
	return s.repo.GetByID(id)
}

func (s *UsuarioService) GetUsuarioByEmail(email string) (*domain.Usuario, error) {
	return s.repo.GetByEmail(email)
}

func (s *UsuarioService) ListUsuarios() ([]domain.Usuario, error) {
	return s.repo.GetAll()
}

func (s *UsuarioService) UpdateUsuario(usuario *domain.Usuario) error {
	if usuario.ID == "" {
		return errors.New("id do usuário é obrigatório")
	}
	if usuario.Nome == "" {
		return errors.New("nome do usuário é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("email do usuário é obrigatório")
	}
	if usuario.Role == "" {
		return errors.New("role do usuário é obrigatória")
	}

	// Busca o usuário existente para manter a data de criação original
	usuarioExistente, err := s.repo.GetByID(usuario.ID)
	if err != nil {
		return err
	}

	// Mantém a data de criação original
	usuario.DataCriacao = usuarioExistente.DataCriacao

	return s.repo.Update(usuario)
}

func (s *UsuarioService) DeleteUsuario(id string) error {
	return s.repo.Delete(id)
}

func (s *UsuarioService) Autenticar(email, senha string) (*domain.Usuario, error) {
	usuario, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !usuario.Ativo {
		return nil, errors.New("usuário inativo")
	}

	// Verifica a senha usando bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	return usuario, nil
}
