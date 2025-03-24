'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Paper,
    Typography,
    Grid,
    TextField,
    Button,
    Switch,
    FormControlLabel,
    Divider,
    Alert,
} from '@mui/material';
import { configuracaoService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';
import { toast } from 'react-toastify';

export default function Configuracoes() {
    const [configuracoes, setConfiguracoes] = useState({
        nomeEmpresa: '',
        cnpj: '',
        endereco: '',
        telefone: '',
        email: '',
        logo: '',
        tema: 'dark',
        notificacoes: true,
        backupAutomatico: true,
        intervaloBackup: 24,
    });
    const { loading, withLoading } = useLoading();

    const loadConfiguracoes = async () => {
        try {
            const response = await configuracaoService.getConfiguracoes();
            setConfiguracoes(response.data);
        } catch (error) {
            toast.error('Erro ao carregar configurações');
        }
    };

    useEffect(() => {
        withLoading(loadConfiguracoes);
    }, []);

    const handleChange = (event) => {
        const { name, value, checked } = event.target;
        setConfiguracoes((prev) => ({
            ...prev,
            [name]: event.target.type === 'checkbox' ? checked : value,
        }));
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await configuracaoService.updateConfiguracoes(configuracoes);
            toast.success('Configurações atualizadas com sucesso');
        } catch (error) {
            toast.error('Erro ao atualizar configurações');
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
                    Configurações do Sistema
                </Typography>

                <Paper sx={{ p: 3 }}>
                    <form onSubmit={handleSubmit}>
                        <Grid container spacing={3}>
                            <Grid item xs={12}>
                                <Typography variant="h6" gutterBottom>
                                    Informações da Empresa
                                </Typography>
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Nome da Empresa"
                                    name="nomeEmpresa"
                                    value={configuracoes.nomeEmpresa}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="CNPJ"
                                    name="cnpj"
                                    value={configuracoes.cnpj}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Endereço"
                                    name="endereco"
                                    value={configuracoes.endereco}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Telefone"
                                    name="telefone"
                                    value={configuracoes.telefone}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="E-mail"
                                    name="email"
                                    type="email"
                                    value={configuracoes.email}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="URL do Logo"
                                    name="logo"
                                    value={configuracoes.logo}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <Divider sx={{ my: 2 }} />
                            </Grid>

                            <Grid item xs={12}>
                                <Typography variant="h6" gutterBottom>
                                    Preferências do Sistema
                                </Typography>
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <FormControlLabel
                                    control={
                                        <Switch
                                            checked={configuracoes.notificacoes}
                                            onChange={handleChange}
                                            name="notificacoes"
                                        />
                                    }
                                    label="Ativar Notificações"
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <FormControlLabel
                                    control={
                                        <Switch
                                            checked={configuracoes.backupAutomatico}
                                            onChange={handleChange}
                                            name="backupAutomatico"
                                        />
                                    }
                                    label="Backup Automático"
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Intervalo de Backup (horas)"
                                    name="intervaloBackup"
                                    type="number"
                                    value={configuracoes.intervaloBackup}
                                    onChange={handleChange}
                                    disabled={!configuracoes.backupAutomatico}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <Alert severity="info" sx={{ mb: 2 }}>
                                    As alterações serão aplicadas após reiniciar o sistema
                                </Alert>

                                <Button
                                    type="submit"
                                    variant="contained"
                                    color="primary"
                                    size="large"
                                >
                                    Salvar Configurações
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Paper>
            </Box>
        </Box>
    );
} 