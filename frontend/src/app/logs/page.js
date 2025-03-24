'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Paper,
    Typography,
    Grid,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    FormControl,
    InputLabel,
    Select,
    MenuItem,
    TextField,
    Button,
    Chip,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { ptBR } from 'date-fns/locale';
import { logService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';
import { toast } from 'react-toastify';

export default function Logs() {
    const [logs, setLogs] = useState([]);
    const [nivel, setNivel] = useState('todos');
    const [dataInicio, setDataInicio] = useState(null);
    const [dataFim, setDataFim] = useState(null);
    const [busca, setBusca] = useState('');
    const { loading, withLoading } = useLoading();

    const loadLogs = async () => {
        try {
            const response = await logService.getLogs({
                nivel,
                dataInicio,
                dataFim,
                busca,
            });
            setLogs(response.data);
        } catch (error) {
            toast.error('Erro ao carregar logs');
        }
    };

    useEffect(() => {
        withLoading(loadLogs);
    }, [nivel, dataInicio, dataFim, busca]);

    const handleNivelChange = (event) => {
        setNivel(event.target.value);
    };

    const handleBuscaChange = (event) => {
        setBusca(event.target.value);
    };

    const getNivelColor = (nivel) => {
        switch (nivel.toLowerCase()) {
            case 'info':
                return 'info';
            case 'warning':
                return 'warning';
            case 'error':
                return 'error';
            default:
                return 'default';
        }
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
                    Logs do Sistema
                </Typography>

                <Paper sx={{ p: 3, mb: 3 }}>
                    <Grid container spacing={3}>
                        <Grid item xs={12} md={3}>
                            <FormControl fullWidth>
                                <InputLabel>Nível</InputLabel>
                                <Select
                                    value={nivel}
                                    label="Nível"
                                    onChange={handleNivelChange}
                                >
                                    <MenuItem value="todos">Todos</MenuItem>
                                    <MenuItem value="info">Info</MenuItem>
                                    <MenuItem value="warning">Warning</MenuItem>
                                    <MenuItem value="error">Error</MenuItem>
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
                            <TextField
                                fullWidth
                                label="Buscar"
                                value={busca}
                                onChange={handleBuscaChange}
                            />
                        </Grid>
                    </Grid>
                </Paper>

                <Paper sx={{ p: 3 }}>
                    <TableContainer>
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Data/Hora</TableCell>
                                    <TableCell>Nível</TableCell>
                                    <TableCell>Usuário</TableCell>
                                    <TableCell>Ação</TableCell>
                                    <TableCell>Detalhes</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {logs.map((log) => (
                                    <TableRow key={log.id}>
                                        <TableCell>
                                            {new Date(log.data).toLocaleString('pt-BR')}
                                        </TableCell>
                                        <TableCell>
                                            <Chip
                                                label={log.nivel}
                                                color={getNivelColor(log.nivel)}
                                                size="small"
                                            />
                                        </TableCell>
                                        <TableCell>{log.usuario}</TableCell>
                                        <TableCell>{log.acao}</TableCell>
                                        <TableCell>{log.detalhes}</TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Paper>
            </Box>
        </Box>
    );
} 