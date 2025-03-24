package repository

import (
	"database/sql"
	"errors"
	"vendas/internal/domain"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) Create(usuario *domain.Usuario) error {
	query := `
        INSERT INTO usuarios (id, nome, email, senha, role, ativo, data_criacao)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.db.Exec(query,
		usuario.ID,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuario.Role,
		usuario.Ativo,
		usuario.DataCriacao,
	)
	return err
}

func (r *UsuarioRepository) GetByID(id string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	query := `
        SELECT id, nome, email, senha, role, ativo, data_criacao
        FROM usuarios
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.Senha,
		&usuario.Role,
		&usuario.Ativo,
		&usuario.DataCriacao,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("usuário não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) GetByEmail(email string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	query := `
        SELECT id, nome, email, senha, role, ativo, data_criacao
        FROM usuarios
        WHERE email = $1
    `
	err := r.db.QueryRow(query, email).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.Senha,
		&usuario.Role,
		&usuario.Ativo,
		&usuario.DataCriacao,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("usuário não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) GetAll() ([]domain.Usuario, error) {
	query := `
        SELECT id, nome, email, senha, role, ativo, data_criacao
        FROM usuarios
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []domain.Usuario
	for rows.Next() {
		var usuario domain.Usuario
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.Senha,
			&usuario.Role,
			&usuario.Ativo,
			&usuario.DataCriacao,
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (r *UsuarioRepository) Update(usuario *domain.Usuario) error {
	query := `
        UPDATE usuarios
        SET nome = $1, email = $2, senha = $3, role = $4, ativo = $5
        WHERE id = $6
    `
	result, err := r.db.Exec(query,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
		usuario.Role,
		usuario.Ativo,
		usuario.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("usuário não encontrado")
	}

	return nil
}

func (r *UsuarioRepository) Delete(id string) error {
	query := `DELETE FROM usuarios WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("usuário não encontrado")
	}

	return nil
}
