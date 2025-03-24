'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Paper,
    Typography,
    Grid,
    FormControl,
    InputLabel,
    Select,
    MenuItem,
    TextField,
    Button,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { ptBR } from 'date-fns/locale';
import {
    BarChart,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer,
} from 'recharts';
import { relatorioService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';

export default function Relatorios() {
    const [relatorio, setRelatorio] = useState(null);
    const [filtro, setFiltro] = useState('diario');
    const [dataInicio, setDataInicio] = useState(null);
    const [dataFim, setDataFim] = useState(null);
    const { loading, withLoading } = useLoading();

    const loadRelatorio = async () => {
        try {
            const response = await relatorioService.getRelatorio(filtro, dataInicio, dataFim);
            setRelatorio(response.data);
        } catch (error) {
            toast.error('Erro ao carregar relatório');
        }
    };

    useEffect(() => {
        withLoading(loadRelatorio);
    }, []);

    const handleFiltroChange = (event) => {
        setFiltro(event.target.value);
    };

    const handleGerarRelatorio = () => {
        withLoading(loadRelatorio);
    };

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
                    Relatórios
                </Typography>

                <Paper sx={{ p: 3, mb: 3 }}>
                    <Grid container spacing={3}>
                        <Grid item xs={12} md={3}>
                            <FormControl fullWidth>
                                <InputLabel>Filtro</InputLabel>
                                <Select
                                    value={filtro}
                                    label="Filtro"
                                    onChange={handleFiltroChange}
                                >
                                    <MenuItem value="diario">Diário</MenuItem>
                                    <MenuItem value="semanal">Semanal</MenuItem>
                                    <MenuItem value="mensal">Mensal</MenuItem>
                                </Select>
                            </FormControl>
                        </Grid>
                        <Grid item xs={12} md={3}>
                            <LocalizationProvider dateAdapter={AdapterDateFns} adapterLocale={ptBR}>
                                <DatePicker
                                    label="Data Início"
                                    value={dataInicio}
                                    onChange={setDataInicio}
                                    renderInput={(params) => <TextField {...params} fullWidth />}
                                />
                            </LocalizationProvider>
                        </Grid>
                        <Grid item xs={12} md={3}>
                            <LocalizationProvider dateAdapter={AdapterDateFns} adapterLocale={ptBR}>
                                <DatePicker
                                    label="Data Fim"
                                    value={dataFim}
                                    onChange={setDataFim}
                                    renderInput={(params) => <TextField {...params} fullWidth />}
                                />
                            </LocalizationProvider>
                        </Grid>
                        <Grid item xs={12} md={3}>
                            <Button
                                variant="contained"
                                fullWidth
                                onClick={handleGerarRelatorio}
                                sx={{ height: '100%' }}
                            >
                                Gerar Relatório
                            </Button>
                        </Grid>
                    </Grid>
                </Paper>

                {relatorio && (
                    <Grid container spacing={3}>
                        <Grid item xs={12} md={6}>
                            <Paper sx={{ p: 3 }}>
                                <Typography variant="h6" gutterBottom>
                                    Vendas por Período
                                </Typography>
                                <Box sx={{ height: 300 }}>
                                    <ResponsiveContainer width="100%" height="100%">
                                        <BarChart data={relatorio.vendasPorPeriodo}>
                                            <CartesianGrid strokeDasharray="3 3" />
                                            <XAxis dataKey="periodo" />
                                            <YAxis />
                                            <Tooltip />
                                            <Legend />
                                            <Bar dataKey="total" fill="#8884d8" />
                                        </BarChart>
                                    </ResponsiveContainer>
                                </Box>
                            </Paper>
                        </Grid>

                        <Grid item xs={12} md={6}>
                            <Paper sx={{ p: 3 }}>
                                <Typography variant="h6" gutterBottom>
                                    Produtos Mais Vendidos
                                </Typography>
                                <Box sx={{ height: 300 }}>
                                    <ResponsiveContainer width="100%" height="100%">
                                        <BarChart data={relatorio.produtosMaisVendidos}>
                                            <CartesianGrid strokeDasharray="3 3" />
                                            <XAxis dataKey="nome" />
                                            <YAxis />
                                            <Tooltip />
                                            <Legend />
                                            <Bar dataKey="quantidade" fill="#82ca9d" />
                                        </BarChart>
                                    </ResponsiveContainer>
                                </Box>
                            </Paper>
                        </Grid>

                        <Grid item xs={12}>
                            <Paper sx={{ p: 3 }}>
                                <Typography variant="h6" gutterBottom>
                                    Resumo
                                </Typography>
                                <Grid container spacing={2}>
                                    <Grid item xs={12} md={3}>
                                        <Paper sx={{ p: 2, bgcolor: 'primary.main', color: 'white' }}>
                                            <Typography variant="h6">Total de Vendas</Typography>
                                            <Typography variant="h4">
                                                R$ {relatorio.totalVendas.toFixed(2)}
                                            </Typography>
                                        </Paper>
                                    </Grid>
                                    <Grid item xs={12} md={3}>
                                        <Paper sx={{ p: 2, bgcolor: 'success.main', color: 'white' }}>
                                            <Typography variant="h6">Ticket Médio</Typography>
                                            <Typography variant="h4">
                                                R$ {relatorio.ticketMedio.toFixed(2)}
                                            </Typography>
                                        </Paper>
                                    </Grid>
                                    <Grid item xs={12} md={3}>
                                        <Paper sx={{ p: 2, bgcolor: 'warning.main', color: 'white' }}>
                                            <Typography variant="h6">Total de Produtos</Typography>
                                            <Typography variant="h4">{relatorio.totalProdutos}</Typography>
                                        </Paper>
                                    </Grid>
                                    <Grid item xs={12} md={3}>
                                        <Paper sx={{ p: 2, bgcolor: 'info.main', color: 'white' }}>
                                            <Typography variant="h6">Produtos em Baixa</Typography>
                                            <Typography variant="h4">
                                                {relatorio.produtosBaixa}
                                            </Typography>
                                        </Paper>
                                    </Grid>
                                </Grid>
                            </Paper>
                        </Grid>
                    </Grid>
                )}
            </Box>
        </Box>
    );
} 