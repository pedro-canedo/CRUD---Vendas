package repository

import (
	"database/sql"
	"errors"
	"vendas/internal/domain"
)

type ClienteRepository struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) Create(cliente *domain.Cliente) error {
	query := `
        INSERT INTO clientes (id, nome, email, telefone, endereco, cpf, usuario_id, data_criacao)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
	_, err := r.db.Exec(query,
		cliente.ID,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
		cliente.Endereco,
		cliente.CPF,
		cliente.UsuarioID,
		cliente.DataCriacao,
	)
	return err
}

func (r *ClienteRepository) GetByID(id string) (*domain.Cliente, error) {
	var cliente domain.Cliente
	query := `
        SELECT id, nome, email, telefone, endereco, cpf, usuario_id, data_criacao
        FROM clientes
        WHERE id = $1
    `
	err := r.db.QueryRow(query, id).Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Email,
		&cliente.Telefone,
		&cliente.Endereco,
		&cliente.CPF,
		&cliente.UsuarioID,
		&cliente.DataCriacao,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("cliente não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *ClienteRepository) GetByCPF(cpf string) (*domain.Cliente, error) {
	var cliente domain.Cliente
	query := `
        SELECT id, nome, email, telefone, endereco, cpf, usuario_id, data_criacao
        FROM clientes
        WHERE cpf = $1
    `
	err := r.db.QueryRow(query, cpf).Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Email,
		&cliente.Telefone,
		&cliente.Endereco,
		&cliente.CPF,
		&cliente.UsuarioID,
		&cliente.DataCriacao,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("cliente não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *ClienteRepository) GetAll() ([]domain.Cliente, error) {
	query := `
        SELECT id, nome, email, telefone, endereco, cpf, usuario_id, data_criacao
        FROM clientes
    `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []domain.Cliente
	for rows.Next() {
		var cliente domain.Cliente
		err := rows.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Email,
			&cliente.Telefone,
			&cliente.Endereco,
			&cliente.CPF,
			&cliente.UsuarioID,
			&cliente.DataCriacao,
		)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}

func (r *ClienteRepository) Update(cliente *domain.Cliente) error {
	query := `
        UPDATE clientes
        SET nome = $1, email = $2, telefone = $3, endereco = $4, cpf = $5, usuario_id = $6
        WHERE id = $7
    `
	result, err := r.db.Exec(query,
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
		cliente.Endereco,
		cliente.CPF,
		cliente.UsuarioID,
		cliente.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cliente não encontrado")
	}

	return nil
}

func (r *ClienteRepository) Delete(id string) error {
	query := `DELETE FROM clientes WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("cliente não encontrado")
	}

	return nil
}
