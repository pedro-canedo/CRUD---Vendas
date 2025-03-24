'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Paper,
    Typography,
    Grid,
    Button,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    IconButton,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField,
    Alert,
} from '@mui/material';
import {
    Download as DownloadIcon,
    Delete as DeleteIcon,
    Add as AddIcon,
} from '@mui/icons-material';
import { backupService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';
import { toast } from 'react-toastify';

export default function Backup() {
    const [backups, setBackups] = useState([]);
    const [openDialog, setOpenDialog] = useState(false);
    const [descricao, setDescricao] = useState('');
    const { loading, withLoading } = useLoading();

    const loadBackups = async () => {
        try {
            const response = await backupService.getBackups();
            setBackups(response.data);
        } catch (error) {
            toast.error('Erro ao carregar backups');
        }
    };

    useEffect(() => {
        withLoading(loadBackups);
    }, []);

    const handleOpenDialog = () => {
        setOpenDialog(true);
    };

    const handleCloseDialog = () => {
        setOpenDialog(false);
        setDescricao('');
    };

    const handleCreateBackup = async () => {
        try {
            await backupService.createBackup(descricao);
            toast.success('Backup criado com sucesso');
            handleCloseDialog();
            withLoading(loadBackups);
        } catch (error) {
            toast.error('Erro ao criar backup');
        }
    };

    const handleDownloadBackup = async (id) => {
        try {
            const response = await backupService.downloadBackup(id);
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', `backup-${id}.sql`);
            document.body.appendChild(link);
            link.click();
            link.remove();
        } catch (error) {
            toast.error('Erro ao baixar backup');
        }
    };

    const handleDeleteBackup = async (id) => {
        if (window.confirm('Tem certeza que deseja excluir este backup?')) {
            try {
                await backupService.deleteBackup(id);
                toast.success('Backup excluído com sucesso');
                withLoading(loadBackups);
            } catch (error) {
                toast.error('Erro ao excluir backup');
            }
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
                    Backup do Sistema
                </Typography>

                <Paper sx={{ p: 3, mb: 3 }}>
                    <Grid container spacing={2} alignItems="center">
                        <Grid item xs={12} md={6}>
                            <Alert severity="info">
                                Os backups são realizados automaticamente conforme configurado nas
                                configurações do sistema
                            </Alert>
                        </Grid>
                        <Grid item xs={12} md={6} sx={{ textAlign: 'right' }}>
                            <Button
                                variant="contained"
                                startIcon={<AddIcon />}
                                onClick={handleOpenDialog}
                            >
                                Criar Backup Manual
                            </Button>
                        </Grid>
                    </Grid>
                </Paper>

                <Paper sx={{ p: 3 }}>
                    <TableContainer>
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Data</TableCell>
                                    <TableCell>Descrição</TableCell>
                                    <TableCell>Tamanho</TableCell>
                                    <TableCell align="right">Ações</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {backups.map((backup) => (
                                    <TableRow key={backup.id}>
                                        <TableCell>
                                            {new Date(backup.data).toLocaleString('pt-BR')}
                                        </TableCell>
                                        <TableCell>{backup.descricao}</TableCell>
                                        <TableCell>{backup.tamanho}</TableCell>
                                        <TableCell align="right">
                                            <IconButton
                                                color="primary"
                                                onClick={() => handleDownloadBackup(backup.id)}
                                            >
                                                <DownloadIcon />
                                            </IconButton>
                                            <IconButton
                                                color="error"
                                                onClick={() => handleDeleteBackup(backup.id)}
                                            >
                                                <DeleteIcon />
                                            </IconButton>
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Paper>

                <Dialog open={openDialog} onClose={handleCloseDialog}>
                    <DialogTitle>Criar Backup Manual</DialogTitle>
                    <DialogContent>
                        <TextField
                            autoFocus
                            margin="dense"
                            label="Descrição"
                            fullWidth
                            value={descricao}
                            onChange={(e) => setDescricao(e.target.value)}
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleCloseDialog}>Cancelar</Button>
                        <Button onClick={handleCreateBackup} variant="contained">
                            Criar
                        </Button>
                    </DialogActions>
                </Dialog>
            </Box>
        </Box>
    );
} 