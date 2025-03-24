'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Grid,
    Paper,
    Typography,
    Card,
    CardContent,
    List,
    ListItem,
    ListItemText,
    Divider,
} from '@mui/material';
import {
    TrendingUp as TrendingUpIcon,
    ShoppingCart as ShoppingCartIcon,
    Inventory as InventoryIcon,
    Warning as WarningIcon,
} from '@mui/icons-material';
import { vendaService, produtoService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';
import { toast } from 'react-toastify';

export default function Dashboard() {
    const [dados, setDados] = useState({
        totalVendas: 0,
        totalProdutos: 0,
        produtosBaixa: 0,
        vendasHoje: 0,
        ultimasVendas: [],
        produtosBaixaEstoque: [],
    });
    const { loading, withLoading } = useLoading();

    const loadDados = async () => {
        try {
            const [vendasResponse, produtosResponse] = await Promise.all([
                vendaService.getAll(),
                produtoService.getAll(),
            ]);

            const vendas = vendasResponse.data;
            const produtos = produtosResponse.data;

            const hoje = new Date();
            hoje.setHours(0, 0, 0, 0);

            const vendasHoje = vendas.filter(
                (venda) => new Date(venda.dataVenda) >= hoje
            );

            const produtosBaixa = produtos.filter(
                (produto) => produto.quantidade <= 5
            );

            setDados({
                totalVendas: vendas.reduce(
                    (total, venda) =>
                        total +
                        venda.itens.reduce(
                            (subtotal, item) =>
                                subtotal + item.produto.preco * item.quantidade,
                            0
                        ),
                    0
                ),
                totalProdutos: produtos.length,
                produtosBaixa: produtosBaixa.length,
                vendasHoje: vendasHoje.length,
                ultimasVendas: vendas.slice(0, 5),
                produtosBaixaEstoque: produtosBaixa.slice(0, 5),
            });
        } catch (error) {
            toast.error('Erro ao carregar dados do dashboard');
        }
    };

    useEffect(() => {
        withLoading(loadDados);
    }, []);

    if (loading) {
        return <Loading />;
    }

    return (
        <Box>
            <Navbar />
            <Box
                component="main"
                sx={{
                    flexGrow: 1,
                    p: 3,
                    width: { sm: `calc(100% - 240px)` },
                    ml: { sm: '240px' },
                }}
            >
                <Typography variant="h4" gutterBottom>
                    Dashboard
                </Typography>

                <Grid container spacing={3}>
                    <Grid item xs={12} sm={6} md={3}>
                        <Card>
                            <CardContent>
                                <Box
                                    sx={{
                                        display: 'flex',
                                        alignItems: 'center',
                                        mb: 2,
                                    }}
                                >
                                    <TrendingUpIcon
                                        sx={{ fontSize: 40, color: 'primary.main', mr: 1 }}
                                    />
                                    <Typography variant="h6">Total de Vendas</Typography>
                                </Box>
                                <Typography variant="h4">
                                    R$ {dados.totalVendas.toFixed(2)}
                                </Typography>
                            </CardContent>
                        </Card>
                    </Grid>

                    <Grid item xs={12} sm={6} md={3}>
                        <Card>
                            <CardContent>
                                <Box
                                    sx={{
                                        display: 'flex',
                                        alignItems: 'center',
                                        mb: 2,
                                    }}
                                >
                                    <ShoppingCartIcon
                                        sx={{ fontSize: 40, color: 'success.main', mr: 1 }}
                                    />
                                    <Typography variant="h6">Vendas Hoje</Typography>
                                </Box>
                                <Typography variant="h4">{dados.vendasHoje}</Typography>
                            </CardContent>
                        </Card>
                    </Grid>

                    <Grid item xs={12} sm={6} md={3}>
                        <Card>
                            <CardContent>
                                <Box
                                    sx={{
                                        display: 'flex',
                                        alignItems: 'center',
                                        mb: 2,
                                    }}
                                >
                                    <InventoryIcon
                                        sx={{ fontSize: 40, color: 'info.main', mr: 1 }}
                                    />
                                    <Typography variant="h6">Total de Produtos</Typography>
                                </Box>
                                <Typography variant="h4">{dados.totalProdutos}</Typography>
                            </CardContent>
                        </Card>
                    </Grid>

                    <Grid item xs={12} sm={6} md={3}>
                        <Card>
                            <CardContent>
                                <Box
                                    sx={{
                                        display: 'flex',
                                        alignItems: 'center',
                                        mb: 2,
                                    }}
                                >
                                    <WarningIcon
                                        sx={{ fontSize: 40, color: 'warning.main', mr: 1 }}
                                    />
                                    <Typography variant="h6">Produtos em Baixa</Typography>
                                </Box>
                                <Typography variant="h4">{dados.produtosBaixa}</Typography>
                            </CardContent>
                        </Card>
                    </Grid>

                    <Grid item xs={12} md={6}>
                        <Paper sx={{ p: 3 }}>
                            <Typography variant="h6" gutterBottom>
                                Últimas Vendas
                            </Typography>
                            <List>
                                {dados.ultimasVendas.map((venda, index) => (
                                    <div key={venda.id}>
                                        <ListItem>
                                            <ListItemText
                                                primary={`Venda #${venda.id}`}
                                                secondary={`${new Date(
                                                    venda.dataVenda
                                                ).toLocaleDateString()} - R$ ${venda.itens
                                                    .reduce(
                                                        (total, item) =>
                                                            total +
                                                            item.produto.preco *
                                                            item.quantidade,
                                                        0
                                                    )
                                                    .toFixed(2)}`}
                                            />
                                        </ListItem>
                                        {index < dados.ultimasVendas.length - 1 && (
                                            <Divider />
                                        )}
                                    </div>
                                ))}
                            </List>
                        </Paper>
                    </Grid>

                    <Grid item xs={12} md={6}>
                        <Paper sx={{ p: 3 }}>
                            <Typography variant="h6" gutterBottom>
                                Produtos em Baixa Estoque
                            </Typography>
                            <List>
                                {dados.produtosBaixaEstoque.map((produto, index) => (
                                    <div key={produto.id}>
                                        <ListItem>
                                            <ListItemText
                                                primary={produto.nome}
                                                secondary={`Quantidade: ${produto.quantidade}`}
                                            />
                                        </ListItem>
                                        {index < dados.produtosBaixaEstoque.length - 1 && (
                                            <Divider />
                                        )}
                                    </div>
                                ))}
                            </List>
                        </Paper>
                    </Grid>
                </Grid>
            </Box>
        </Box>
    );
} 