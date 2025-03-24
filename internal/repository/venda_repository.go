package repository

import (
	"database/sql"
	"fmt"
	"vendas/internal/domain"
	"vendas/internal/utils"
)

type VendaRepository interface {
	Create(venda *domain.Venda) error
	GetByID(id string) (*domain.Venda, error)
	GetAll() ([]domain.Venda, error)
	Update(venda *domain.Venda) error
	Delete(id string) error
	GetVendasPorCliente(cliente string) ([]domain.Venda, error)
	GetVendasPorPeriodo(inicio, fim int64) ([]domain.Venda, error)
}

type VendaRepositoryImpl struct {
	db *sql.DB
}

func NewVendaRepository(db *sql.DB) *VendaRepositoryImpl {
	return &VendaRepositoryImpl{db: db}
}

func (r *VendaRepositoryImpl) Create(venda *domain.Venda) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Gera UUID para a venda
	venda.ID = utils.GenerateUUID()

	// Insere a venda
	query := `INSERT INTO vendas (id, cliente, data_venda, total) VALUES (?, ?, ?, ?)`
	_, err = tx.Exec(query, venda.ID, venda.Cliente, venda.DataVenda, venda.Total)
	if err != nil {
		return err
	}

	// Insere os itens da venda
	for i := range venda.Itens {
		// Gera UUID para o item
		venda.Itens[i].ID = utils.GenerateUUID()

		// Insere o item
		query = `INSERT INTO itens_venda (id, venda_id, produto_id, quantidade, preco_unitario, subtotal) 
			VALUES (?, ?, ?, ?, ?, ?)`
		_, err = tx.Exec(query,
			venda.Itens[i].ID,
			venda.ID,
			venda.Itens[i].ProdutoID,
			venda.Itens[i].Quantidade,
			venda.Itens[i].PrecoUnitario,
			venda.Itens[i].Subtotal)
		if err != nil {
			return err
		}

		// Atualiza o estoque do produto
		query = `UPDATE produtos 
			SET quantidade = quantidade - ? 
			WHERE id = ? AND quantidade >= ?`
		result, err := tx.Exec(query,
			venda.Itens[i].Quantidade,
			venda.Itens[i].ProdutoID,
			venda.Itens[i].Quantidade)
		if err != nil {
			return err
		}

		// Verifica se o produto foi atualizado
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			return fmt.Errorf("estoque insuficiente para o produto %s", venda.Itens[i].ProdutoID)
		}
	}

	return tx.Commit()
}

func (r *VendaRepositoryImpl) GetByID(id string) (*domain.Venda, error) {
	venda := &domain.Venda{}
	query := `SELECT id, cliente, data_venda, total FROM vendas WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&venda.ID, &venda.Cliente, &venda.DataVenda, &venda.Total)
	if err != nil {
		return nil, err
	}

	query = `SELECT id, produto_id, quantidade, preco_unitario, subtotal FROM itens_venda WHERE venda_id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.ItemVenda
		err := rows.Scan(&item.ID, &item.ProdutoID, &item.Quantidade, &item.PrecoUnitario, &item.Subtotal)
		if err != nil {
			return nil, err
		}
		venda.Itens = append(venda.Itens, item)
	}

	return venda, nil
}

func (r *VendaRepositoryImpl) GetAll() ([]domain.Venda, error) {
	query := `SELECT id, cliente, data_venda, total FROM vendas`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		err := rows.Scan(&venda.ID, &venda.Cliente, &venda.DataVenda, &venda.Total)
		if err != nil {
			return nil, err
		}

		query = `SELECT id, produto_id, quantidade, preco_unitario, subtotal FROM itens_venda WHERE venda_id = ?`
		itemRows, err := r.db.Query(query, venda.ID)
		if err != nil {
			return nil, err
		}

		for itemRows.Next() {
			var item domain.ItemVenda
			err := itemRows.Scan(&item.ID, &item.ProdutoID, &item.Quantidade, &item.PrecoUnitario, &item.Subtotal)
			if err != nil {
				itemRows.Close()
				return nil, err
			}
			venda.Itens = append(venda.Itens, item)
		}
		itemRows.Close()

		vendas = append(vendas, venda)
	}
	return vendas, nil
}

func (r *VendaRepositoryImpl) Update(venda *domain.Venda) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Restaurar estoque dos itens antigos
	query := `SELECT produto_id, quantidade FROM itens_venda WHERE venda_id = ?`
	rows, err := tx.Query(query, venda.ID)
	if err != nil {
		return err
	}

	for rows.Next() {
		var produtoID, quantidade int64
		err := rows.Scan(&produtoID, &quantidade)
		if err != nil {
			rows.Close()
			return err
		}

		query = `UPDATE produtos SET quantidade = quantidade + ? WHERE id = ?`
		_, err = tx.Exec(query, quantidade, produtoID)
		if err != nil {
			rows.Close()
			return err
		}
	}
	rows.Close()

	// Atualizar venda
	query = `UPDATE vendas SET cliente = ?, data_venda = ?, total = ? WHERE id = ?`
	_, err = tx.Exec(query, venda.Cliente, venda.DataVenda, venda.Total, venda.ID)
	if err != nil {
		return err
	}

	// Remover itens antigos
	query = `DELETE FROM itens_venda WHERE venda_id = ?`
	_, err = tx.Exec(query, venda.ID)
	if err != nil {
		return err
	}

	// Inserir novos itens
	for _, item := range venda.Itens {
		query = `INSERT INTO itens_venda (venda_id, produto_id, quantidade, preco_unitario, subtotal) VALUES (?, ?, ?, ?, ?)`
		_, err = tx.Exec(query, venda.ID, item.ProdutoID, item.Quantidade, item.PrecoUnitario, item.Subtotal)
		if err != nil {
			return err
		}

		query = `UPDATE produtos SET quantidade = quantidade - ? WHERE id = ?`
		_, err = tx.Exec(query, item.Quantidade, item.ProdutoID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *VendaRepositoryImpl) Delete(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Restaurar estoque
	query := `SELECT produto_id, quantidade FROM itens_venda WHERE venda_id = ?`
	rows, err := tx.Query(query, id)
	if err != nil {
		return err
	}

	for rows.Next() {
		var produtoID string
		var quantidade int
		err := rows.Scan(&produtoID, &quantidade)
		if err != nil {
			rows.Close()
			return err
		}

		query = `UPDATE produtos SET quantidade = quantidade + ? WHERE id = ?`
		_, err = tx.Exec(query, quantidade, produtoID)
		if err != nil {
			rows.Close()
			return err
		}
	}
	rows.Close()

	// Remover itens
	query = `DELETE FROM itens_venda WHERE venda_id = ?`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	// Remover venda
	query = `DELETE FROM vendas WHERE id = ?`
	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Métodos adicionais específicos para vendas

func (r *VendaRepositoryImpl) GetVendasPorCliente(cliente string) ([]domain.Venda, error) {
	query := `SELECT id, cliente, data_venda, total FROM vendas WHERE cliente = ?`
	rows, err := r.db.Query(query, cliente)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		err := rows.Scan(&venda.ID, &venda.Cliente, &venda.DataVenda, &venda.Total)
		if err != nil {
			return nil, err
		}

		query = `SELECT produto_id, quantidade, preco_unitario FROM itens_venda WHERE venda_id = ?`
		itemRows, err := r.db.Query(query, venda.ID)
		if err != nil {
			return nil, err
		}

		for itemRows.Next() {
			var item domain.ItemVenda
			err := itemRows.Scan(&item.ProdutoID, &item.Quantidade, &item.PrecoUnitario)
			if err != nil {
				itemRows.Close()
				return nil, err
			}
			venda.Itens = append(venda.Itens, item)
		}
		itemRows.Close()

		vendas = append(vendas, venda)
	}
	return vendas, nil
}

func (r *VendaRepositoryImpl) GetVendasPorPeriodo(inicio, fim int64) ([]domain.Venda, error) {
	query := `SELECT id, cliente, data_venda, total FROM vendas WHERE data_venda BETWEEN ? AND ?`
	rows, err := r.db.Query(query, inicio, fim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		err := rows.Scan(&venda.ID, &venda.Cliente, &venda.DataVenda, &venda.Total)
		if err != nil {
			return nil, err
		}

		query = `SELECT produto_id, quantidade, preco_unitario FROM itens_venda WHERE venda_id = ?`
		itemRows, err := r.db.Query(query, venda.ID)
		if err != nil {
			return nil, err
		}

		for itemRows.Next() {
			var item domain.ItemVenda
			err := itemRows.Scan(&item.ProdutoID, &item.Quantidade, &item.PrecoUnitario)
			if err != nil {
				itemRows.Close()
				return nil, err
			}
			venda.Itens = append(venda.Itens, item)
		}
		itemRows.Close()

		vendas = append(vendas, venda)
	}
	return vendas, nil
}
