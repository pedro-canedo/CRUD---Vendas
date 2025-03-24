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
	query := `INSERT INTO vendas (id, cliente_id, vendedor_id, data_venda, valor_total, data_criacao) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, venda.ID, venda.ClienteID, venda.VendedorID, venda.DataVenda, venda.ValorTotal, venda.DataCriacao)
	if err != nil {
		return err
	}

	// Insere os itens da venda
	for i := range venda.Items {
		// Gera UUID para o item
		venda.Items[i].ID = utils.GenerateUUID()

		// Insere o item
		query = `INSERT INTO itens_venda (id, venda_id, produto_id, quantidade, preco_unitario) 
			VALUES (?, ?, ?, ?, ?)`
		_, err = tx.Exec(query,
			venda.Items[i].ID,
			venda.ID,
			venda.Items[i].ProdutoID,
			venda.Items[i].Quantidade,
			venda.Items[i].PrecoUnitario)
		if err != nil {
			return err
		}

		// Atualiza o estoque do produto
		query = `UPDATE produtos 
			SET quantidade = quantidade - ? 
			WHERE id = ? AND quantidade >= ?`
		result, err := tx.Exec(query,
			venda.Items[i].Quantidade,
			venda.Items[i].ProdutoID,
			venda.Items[i].Quantidade)
		if err != nil {
			return err
		}

		// Verifica se o produto foi atualizado
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			return fmt.Errorf("estoque insuficiente para o produto %s", venda.Items[i].ProdutoID)
		}
	}

	return tx.Commit()
}

func (r *VendaRepositoryImpl) GetByID(id string) (*domain.Venda, error) {
	var venda domain.Venda
	var clienteID, vendedorID string
	var clienteNome, vendedorNome string

	// Busca os dados da venda
	err := r.db.QueryRow(`
		SELECT v.id, v.cliente_id, v.vendedor_id, v.data_venda, v.valor_total, v.data_criacao,
			   c.nome as cliente_nome, vd.nome as vendedor_nome
		FROM vendas v
		LEFT JOIN usuarios c ON v.cliente_id = c.id
		LEFT JOIN usuarios vd ON v.vendedor_id = vd.id
		WHERE v.id = ?
	`, id).Scan(&venda.ID, &clienteID, &vendedorID, &venda.DataVenda, &venda.ValorTotal, &venda.DataCriacao,
		&clienteNome, &vendedorNome)

	if err != nil {
		return nil, err
	}

	// Busca os itens da venda
	rows, err := r.db.Query(`
		SELECT iv.id, iv.produto_id, iv.quantidade, iv.preco_unitario,
			   COALESCE(p.id, '') as produto_id,
			   COALESCE(p.nome, 'Produto não encontrado') as produto_nome,
			   COALESCE(p.descricao, '') as produto_descricao,
			   COALESCE(p.preco, 0) as produto_preco,
			   COALESCE(p.quantidade, 0) as produto_quantidade,
			   COALESCE(p.imagem_url, '') as produto_imagem_url,
			   COALESCE(p.data_criacao, datetime('now')) as produto_data_criacao
		FROM itens_venda iv
		LEFT JOIN produtos p ON iv.produto_id = p.id
		WHERE iv.venda_id = ?
	`, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Adiciona os dados do cliente e vendedor
	venda.Cliente = &domain.Usuario{
		ID:   clienteID,
		Nome: clienteNome,
	}
	venda.Vendedor = &domain.Usuario{
		ID:   vendedorID,
		Nome: vendedorNome,
	}

	// Processa os itens
	for rows.Next() {
		var item domain.ItemVenda
		var produto domain.Produto
		err := rows.Scan(
			&item.ID,
			&item.ProdutoID,
			&item.Quantidade,
			&item.PrecoUnitario,
			&produto.ID,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
			&produto.ImagemURL,
			&produto.DataCriacao,
		)
		if err != nil {
			return nil, err
		}
		item.Produto = &produto
		venda.Items = append(venda.Items, item)
	}

	return &venda, nil
}

func (r *VendaRepositoryImpl) GetAll() ([]domain.Venda, error) {
	query := `
		SELECT 
			v.id, 
			v.cliente_id, 
			v.vendedor_id, 
			v.data_venda, 
			v.valor_total, 
			v.data_criacao,
			c.nome as cliente_nome,
			u.nome as vendedor_nome
		FROM vendas v
		JOIN usuarios c ON c.id = v.cliente_id
		JOIN usuarios u ON u.id = v.vendedor_id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		var clienteNome, vendedorNome string
		err := rows.Scan(
			&venda.ID,
			&venda.ClienteID,
			&venda.VendedorID,
			&venda.DataVenda,
			&venda.ValorTotal,
			&venda.DataCriacao,
			&clienteNome,
			&vendedorNome,
		)
		if err != nil {
			return nil, err
		}

		// Adiciona os dados do cliente e vendedor
		venda.Cliente = &domain.Usuario{
			ID:   venda.ClienteID,
			Nome: clienteNome,
		}
		venda.Vendedor = &domain.Usuario{
			ID:   venda.VendedorID,
			Nome: vendedorNome,
		}

		query = `SELECT id, produto_id, quantidade, preco_unitario FROM itens_venda WHERE venda_id = ?`
		itemRows, err := r.db.Query(query, venda.ID)
		if err != nil {
			return nil, err
		}

		for itemRows.Next() {
			var item domain.ItemVenda
			err := itemRows.Scan(&item.ID, &item.ProdutoID, &item.Quantidade, &item.PrecoUnitario)
			if err != nil {
				itemRows.Close()
				return nil, err
			}
			venda.Items = append(venda.Items, item)
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
	query = `UPDATE vendas SET cliente_id = ?, vendedor_id = ?, data_venda = ?, valor_total = ? WHERE id = ?`
	_, err = tx.Exec(query, venda.ClienteID, venda.VendedorID, venda.DataVenda, venda.ValorTotal, venda.ID)
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
	for _, item := range venda.Items {
		query = `INSERT INTO itens_venda (venda_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?)`
		_, err = tx.Exec(query, venda.ID, item.ProdutoID, item.Quantidade, item.PrecoUnitario)
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
	query := `SELECT id, cliente_id, vendedor_id, data_venda, valor_total, data_criacao FROM vendas WHERE cliente_id = ?`
	rows, err := r.db.Query(query, cliente)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		err := rows.Scan(&venda.ID, &venda.ClienteID, &venda.VendedorID, &venda.DataVenda, &venda.ValorTotal, &venda.DataCriacao)
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
			venda.Items = append(venda.Items, item)
		}
		itemRows.Close()

		vendas = append(vendas, venda)
	}
	return vendas, nil
}

func (r *VendaRepositoryImpl) GetVendasPorPeriodo(inicio, fim int64) ([]domain.Venda, error) {
	query := `SELECT id, cliente_id, vendedor_id, data_venda, valor_total, data_criacao FROM vendas WHERE data_venda BETWEEN ? AND ?`
	rows, err := r.db.Query(query, inicio, fim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendas []domain.Venda
	for rows.Next() {
		var venda domain.Venda
		err := rows.Scan(&venda.ID, &venda.ClienteID, &venda.VendedorID, &venda.DataVenda, &venda.ValorTotal, &venda.DataCriacao)
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
			venda.Items = append(venda.Items, item)
		}
		itemRows.Close()

		vendas = append(vendas, venda)
	}
	return vendas, nil
}
