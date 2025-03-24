'use client';

import { useState, useEffect } from 'react';
import {
    Box,
    Button,
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    IconButton,
    Typography,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField,
} from '@mui/material';
import {
    Add as AddIcon,
    Edit as EditIcon,
    Delete as DeleteIcon,
} from '@mui/icons-material';
import { toast } from 'react-toastify';
import { produtoService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';

export default function Produtos() {
    const [produtos, setProdutos] = useState([]);
    const [openDialog, setOpenDialog] = useState(false);
    const [selectedProduto, setSelectedProduto] = useState(null);
    const [formData, setFormData] = useState({
        nome: '',
        descricao: '',
        preco: '',
        quantidade: '',
    });
    const { loading, withLoading } = useLoading();

    const loadProdutos = async () => {
        try {
            const response = await produtoService.getAll();
            setProdutos(response.data);
        } catch (error) {
            toast.error('Erro ao carregar produtos');
        }
    };

    useEffect(() => {
        withLoading(loadProdutos);
    }, []);

    const handleOpenDialog = (produto = null) => {
        if (produto) {
            setFormData({
                nome: produto.nome,
                descricao: produto.descricao,
                preco: produto.preco,
                quantidade: produto.quantidade,
            });
            setSelectedProduto(produto);
        } else {
            setFormData({
                nome: '',
                descricao: '',
                preco: '',
                quantidade: '',
            });
            setSelectedProduto(null);
        }
        setOpenDialog(true);
    };

    const handleCloseDialog = () => {
        setOpenDialog(false);
        setSelectedProduto(null);
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            if (selectedProduto) {
                await produtoService.update(selectedProduto.id, formData);
                toast.success('Produto atualizado com sucesso!');
            } else {
                await produtoService.create(formData);
                toast.success('Produto cadastrado com sucesso!');
            }
            handleCloseDialog();
            loadProdutos();
        } catch (error) {
            toast.error('Erro ao salvar produto');
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Tem certeza que deseja excluir este produto?')) {
            try {
                await produtoService.delete(id);
                toast.success('Produto excluído com sucesso!');
                loadProdutos();
            } catch (error) {
                toast.error('Erro ao excluir produto');
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
                <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
                    <Typography variant="h4">Produtos</Typography>
                    <Button
                        variant="contained"
                        startIcon={<AddIcon />}
                        onClick={() => handleOpenDialog()}
                    >
                        Novo Produto
                    </Button>
                </Box>

                <TableContainer component={Paper}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Nome</TableCell>
                                <TableCell>Descrição</TableCell>
                                <TableCell align="right">Preço</TableCell>
                                <TableCell align="right">Quantidade</TableCell>
                                <TableCell align="center">Ações</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {produtos.map((produto) => (
                                <TableRow key={produto.id}>
                                    <TableCell>{produto.nome}</TableCell>
                                    <TableCell>{produto.descricao}</TableCell>
                                    <TableCell align="right">
                                        R$ {produto.preco.toFixed(2)}
                                    </TableCell>
                                    <TableCell align="right">{produto.quantidade}</TableCell>
                                    <TableCell align="center">
                                        <IconButton
                                            color="primary"
                                            onClick={() => handleOpenDialog(produto)}
                                        >
                                            <EditIcon />
                                        </IconButton>
                                        <IconButton
                                            color="error"
                                            onClick={() => handleDelete(produto.id)}
                                        >
                                            <DeleteIcon />
                                        </IconButton>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>

                <Dialog open={openDialog} onClose={handleCloseDialog}>
                    <DialogTitle>
                        {selectedProduto ? 'Editar Produto' : 'Novo Produto'}
                    </DialogTitle>
                    <form onSubmit={handleSubmit}>
                        <DialogContent>
                            <TextField
                                autoFocus
                                margin="dense"
                                name="nome"
                                label="Nome"
                                type="text"
                                fullWidth
                                required
                                value={formData.nome}
                                onChange={handleChange}
                            />
                            <TextField
                                margin="dense"
                                name="descricao"
                                label="Descrição"
                                type="text"
                                fullWidth
                                multiline
                                rows={3}
                                value={formData.descricao}
                                onChange={handleChange}
                            />
                            <TextField
                                margin="dense"
                                name="preco"
                                label="Preço"
                                type="number"
                                fullWidth
                                required
                                value={formData.preco}
                                onChange={handleChange}
                            />
                            <TextField
                                margin="dense"
                                name="quantidade"
                                label="Quantidade"
                                type="number"
                                fullWidth
                                required
                                value={formData.quantidade}
                                onChange={handleChange}
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={handleCloseDialog}>Cancelar</Button>
                            <Button type="submit" variant="contained">
                                Salvar
                            </Button>
                        </DialogActions>
                    </form>
                </Dialog>
            </Box>
        </Box>
    );
} 