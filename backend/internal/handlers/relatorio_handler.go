package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RelatorioHandler struct {
	db *sql.DB
}

func NewRelatorioHandler(db *sql.DB) *RelatorioHandler {
	return &RelatorioHandler{db: db}
}

// @Summary Obtém relatório geral
// @Description Retorna dados gerais do sistema como vendas do dia, total de clientes e produtos
// @Tags relatorios
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /relatorios [get]
func (h *RelatorioHandler) GetRelatorio(c *gin.Context) {
	// Vendas do dia
	var vendasDia float64
	err := h.db.QueryRow(`
		SELECT COALESCE(SUM(valor_total), 0)
		FROM vendas
		WHERE date(data_venda) = date('now')
	`).Scan(&vendasDia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter vendas do dia"})
		return
	}

	// Total de clientes
	var totalClientes int
	err = h.db.QueryRow(`
		SELECT COUNT(*)
		FROM usuarios
		WHERE role = 'cliente'
	`).Scan(&totalClientes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter total de clientes"})
		return
	}

	// Total de produtos
	var totalProdutos int
	err = h.db.QueryRow(`
		SELECT COUNT(*)
		FROM produtos
	`).Scan(&totalProdutos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter total de produtos"})
		return
	}

	// Vendas por mês (últimos 12 meses)
	rows, err := h.db.Query(`
		SELECT 
			strftime('%Y-%m', data_venda) as mes,
			COUNT(*) as quantidade,
			SUM(valor_total) as total
		FROM vendas
		WHERE data_venda >= datetime('now', '-12 months')
		GROUP BY strftime('%Y-%m', data_venda)
		ORDER BY mes DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter vendas por mês"})
		return
	}
	defer rows.Close()

	type VendaMes struct {
		Mes        string  `json:"mes"`
		Quantidade int     `json:"quantidade"`
		Total      float64 `json:"total"`
	}

	var vendasPorMes []VendaMes
	for rows.Next() {
		var v VendaMes
		err := rows.Scan(&v.Mes, &v.Quantidade, &v.Total)
		if err != nil {
			continue
		}
		vendasPorMes = append(vendasPorMes, v)
	}

	// Produtos mais vendidos (últimos 30 dias)
	rows, err = h.db.Query(`
		SELECT 
			p.id,
			p.nome,
			SUM(iv.quantidade) as quantidade,
			SUM(iv.quantidade * iv.preco_unitario) as total
		FROM itens_venda iv
		JOIN produtos p ON iv.produto_id = p.id
		JOIN vendas v ON iv.venda_id = v.id
		WHERE v.data_venda >= datetime('now', '-30 days')
		GROUP BY p.id, p.nome
		ORDER BY quantidade DESC
		LIMIT 5
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter produtos mais vendidos"})
		return
	}
	defer rows.Close()

	type ProdutoVenda struct {
		ID         string  `json:"id"`
		Nome       string  `json:"nome"`
		Quantidade int     `json:"quantidade"`
		Total      float64 `json:"total"`
	}

	var produtosMaisVendidos []ProdutoVenda
	for rows.Next() {
		var p ProdutoVenda
		err := rows.Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Total)
		if err != nil {
			continue
		}
		produtosMaisVendidos = append(produtosMaisVendidos, p)
	}

	// Vendas por vendedor (últimos 30 dias)
	rows, err = h.db.Query(`
		SELECT 
			u.id,
			u.nome,
			COUNT(*) as quantidade,
			SUM(v.valor_total) as total
		FROM vendas v
		JOIN usuarios u ON v.vendedor_id = u.id
		WHERE v.data_venda >= datetime('now', '-30 days')
		GROUP BY u.id, u.nome
		ORDER BY total DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter vendas por vendedor"})
		return
	}
	defer rows.Close()

	var vendasPorVendedor []ProdutoVenda
	for rows.Next() {
		var v ProdutoVenda
		err := rows.Scan(&v.ID, &v.Nome, &v.Quantidade, &v.Total)
		if err != nil {
			continue
		}
		vendasPorVendedor = append(vendasPorVendedor, v)
	}

	// Produtos com estoque baixo
	rows, err = h.db.Query(`
		SELECT 
			id,
			nome,
			quantidade,
			preco
		FROM produtos
		WHERE quantidade <= 10
		ORDER BY quantidade ASC
		LIMIT 5
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter produtos com estoque baixo"})
		return
	}
	defer rows.Close()

	type ProdutoEstoque struct {
		ID         string  `json:"id"`
		Nome       string  `json:"nome"`
		Quantidade int     `json:"quantidade"`
		Preco      float64 `json:"preco"`
	}

	var produtosEstoqueBaixo []ProdutoEstoque
	for rows.Next() {
		var p ProdutoEstoque
		err := rows.Scan(&p.ID, &p.Nome, &p.Quantidade, &p.Preco)
		if err != nil {
			continue
		}
		produtosEstoqueBaixo = append(produtosEstoqueBaixo, p)
	}

	c.JSON(http.StatusOK, gin.H{
		"vendas_dia":             vendasDia,
		"total_clientes":         totalClientes,
		"total_produtos":         totalProdutos,
		"vendas_por_mes":         vendasPorMes,
		"produtos_mais_vendidos": produtosMaisVendidos,
		"vendas_por_vendedor":    vendasPorVendedor,
		"produtos_estoque_baixo": produtosEstoqueBaixo,
	})
}
