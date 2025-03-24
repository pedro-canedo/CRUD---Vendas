'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Paper,
    Typography,
    Grid,
    TextField,
    Button,
    Avatar,
    Divider,
    Alert,
} from '@mui/material';
import { usuarioService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';
import { toast } from 'react-toastify';
import { useAuth } from '@/contexts/AuthContext';

export default function Perfil() {
    const { user, updateUser } = useAuth();
    const [perfil, setPerfil] = useState({
        nome: '',
        email: '',
        senha: '',
        confirmarSenha: '',
        foto: '',
    });
    const { loading, withLoading } = useLoading();

    const loadPerfil = async () => {
        try {
            const response = await usuarioService.getPerfil();
            setPerfil({
                ...response.data,
                senha: '',
                confirmarSenha: '',
            });
        } catch (error) {
            toast.error('Erro ao carregar perfil');
        }
    };

    useEffect(() => {
        withLoading(loadPerfil);
    }, []);

    const handleChange = (event) => {
        const { name, value } = event.target;
        setPerfil((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        if (perfil.senha && perfil.senha !== perfil.confirmarSenha) {
            toast.error('As senhas não conferem');
            return;
        }

        try {
            const response = await usuarioService.updatePerfil(perfil);
            updateUser(response.data);
            toast.success('Perfil atualizado com sucesso');
            setPerfil((prev) => ({
                ...prev,
                senha: '',
                confirmarSenha: '',
            }));
        } catch (error) {
            toast.error('Erro ao atualizar perfil');
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
                    Meu Perfil
                </Typography>

                <Paper sx={{ p: 3 }}>
                    <form onSubmit={handleSubmit}>
                        <Grid container spacing={3}>
                            <Grid item xs={12} sx={{ textAlign: 'center' }}>
                                <Avatar
                                    src={perfil.foto}
                                    sx={{ width: 120, height: 120, mx: 'auto', mb: 2 }}
                                />
                                <Button
                                    variant="outlined"
                                    component="label"
                                    sx={{ mb: 3 }}
                                >
                                    Alterar Foto
                                    <input
                                        type="file"
                                        hidden
                                        accept="image/*"
                                        onChange={(e) => {
                                            const file = e.target.files[0];
                                            if (file) {
                                                const reader = new FileReader();
                                                reader.onloadend = () => {
                                                    setPerfil((prev) => ({
                                                        ...prev,
                                                        foto: reader.result,
                                                    }));
                                                };
                                                reader.readAsDataURL(file);
                                            }
                                        }}
                                    />
                                </Button>
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Nome"
                                    name="nome"
                                    value={perfil.nome}
                                    onChange={handleChange}
                                    required
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="E-mail"
                                    name="email"
                                    type="email"
                                    value={perfil.email}
                                    onChange={handleChange}
                                    required
                                    disabled
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <Divider sx={{ my: 2 }} />
                            </Grid>

                            <Grid item xs={12}>
                                <Typography variant="h6" gutterBottom>
                                    Alterar Senha
                                </Typography>
                                <Alert severity="info" sx={{ mb: 2 }}>
                                    Preencha apenas se desejar alterar sua senha
                                </Alert>
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Nova Senha"
                                    name="senha"
                                    type="password"
                                    value={perfil.senha}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12} md={6}>
                                <TextField
                                    fullWidth
                                    label="Confirmar Nova Senha"
                                    name="confirmarSenha"
                                    type="password"
                                    value={perfil.confirmarSenha}
                                    onChange={handleChange}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <Button
                                    type="submit"
                                    variant="contained"
                                    color="primary"
                                    size="large"
                                >
                                    Salvar Alterações
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Paper>
            </Box>
        </Box>
    );
} 