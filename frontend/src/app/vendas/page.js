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
    Grid,
    MenuItem,
    FormControl,
    InputLabel,
    Select,
} from '@mui/material';
import {
    Add as AddIcon,
    Edit as EditIcon,
    Delete as DeleteIcon,
} from '@mui/icons-material';
import { toast } from 'react-toastify';
import { vendaService, produtoService } from '@/services/api';
import Navbar from '@/components/Navbar';
import Loading from '@/components/Loading';
import ErrorMessage from '@/components/ErrorMessage';
import { useLoading } from '@/hooks/useLoading';

export default function Vendas() {
    const [vendas, setVendas] = useState([]);
    const [produtos, setProdutos] = useState([]);
    const [openDialog, setOpenDialog] = useState(false);
    const [selectedVenda, setSelectedVenda] = useState(null);
    const [formData, setFormData] = useState({
        cliente: '',
        itens: [{ produtoId: '', quantidade: 1 }],
    });
    const { loading, withLoading } = useLoading();

    const loadVendas = async () => {
        try {
            const response = await vendaService.getAll();
            setVendas(response.data);
        } catch (error) {
            toast.error('Erro ao carregar vendas');
        }
    };

    const loadProdutos = async () => {
        try {
            const response = await produtoService.getAll();
            setProdutos(response.data);
        } catch (error) {
            toast.error('Erro ao carregar produtos');
        }
    };

    useEffect(() => {
        withLoading(async () => {
            await Promise.all([loadVendas(), loadProdutos()]);
        });
    }, []);

    const handleOpenDialog = (venda = null) => {
        if (venda) {
            setFormData({
                cliente: venda.cliente,
                itens: venda.itens.map((item) => ({
                    produtoId: item.produto.id,
                    quantidade: item.quantidade,
                })),
            });
            setSelectedVenda(venda);
        } else {
            setFormData({
                cliente: '',
                itens: [{ produtoId: '', quantidade: 1 }],
            });
            setSelectedVenda(null);
        }
        setOpenDialog(true);
    };

    const handleCloseDialog = () => {
        setOpenDialog(false);
        setSelectedVenda(null);
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleItemChange = (index, field, value) => {
        setFormData((prev) => ({
            ...prev,
            itens: prev.itens.map((item, i) =>
                i === index ? { ...item, [field]: value } : item
            ),
        }));
    };

    const addItem = () => {
        setFormData((prev) => ({
            ...prev,
            itens: [...prev.itens, { produtoId: '', quantidade: 1 }],
        }));
    };

    const removeItem = (index) => {
        setFormData((prev) => ({
            ...prev,
            itens: prev.itens.filter((_, i) => i !== index),
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            if (selectedVenda) {
                await vendaService.update(selectedVenda.id, formData);
                toast.success('Venda atualizada com sucesso!');
            } else {
                await vendaService.create(formData);
                toast.success('Venda cadastrada com sucesso!');
            }
            handleCloseDialog();
            loadVendas();
        } catch (error) {
            toast.error('Erro ao salvar venda');
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Tem certeza que deseja excluir esta venda?')) {
            try {
                await vendaService.delete(id);
                toast.success('Venda excluída com sucesso!');
                loadVendas();
            } catch (error) {
                toast.error('Erro ao excluir venda');
            }
        }
    };

    const calcularTotal = (venda) => {
        return venda.itens.reduce(
            (total, item) => total + item.produto.preco * item.quantidade,
            0
        );
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
                    <Typography variant="h4">Vendas</Typography>
                    <Button
                        variant="contained"
                        startIcon={<AddIcon />}
                        onClick={() => handleOpenDialog()}
                    >
                        Nova Venda
                    </Button>
                </Box>

                <TableContainer component={Paper}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Cliente</TableCell>
                                <TableCell>Data</TableCell>
                                <TableCell>Itens</TableCell>
                                <TableCell align="right">Total</TableCell>
                                <TableCell align="center">Ações</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {vendas.map((venda) => (
                                <TableRow key={venda.id}>
                                    <TableCell>{venda.cliente}</TableCell>
                                    <TableCell>
                                        {new Date(venda.dataVenda).toLocaleDateString()}
                                    </TableCell>
                                    <TableCell>
                                        {venda.itens
                                            .map(
                                                (item) =>
                                                    `${item.produto.nome} (${item.quantidade})`
                                            )
                                            .join(', ')}
                                    </TableCell>
                                    <TableCell align="right">
                                        R$ {calcularTotal(venda).toFixed(2)}
                                    </TableCell>
                                    <TableCell align="center">
                                        <IconButton
                                            color="primary"
                                            onClick={() => handleOpenDialog(venda)}
                                        >
                                            <EditIcon />
                                        </IconButton>
                                        <IconButton
                                            color="error"
                                            onClick={() => handleDelete(venda.id)}
                                        >
                                            <DeleteIcon />
                                        </IconButton>
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>

                <Dialog open={openDialog} onClose={handleCloseDialog} maxWidth="md" fullWidth>
                    <DialogTitle>
                        {selectedVenda ? 'Editar Venda' : 'Nova Venda'}
                    </DialogTitle>
                    <form onSubmit={handleSubmit}>
                        <DialogContent>
                            <TextField
                                autoFocus
                                margin="dense"
                                name="cliente"
                                label="Cliente"
                                type="text"
                                fullWidth
                                required
                                value={formData.cliente}
                                onChange={handleChange}
                            />

                            <Typography variant="h6" sx={{ mt: 2, mb: 1 }}>
                                Itens
                            </Typography>
                            {formData.itens.map((item, index) => (
                                <Grid container spacing={2} key={index} sx={{ mb: 2 }}>
                                    <Grid item xs={5}>
                                        <FormControl fullWidth>
                                            <InputLabel>Produto</InputLabel>
                                            <Select
                                                value={item.produtoId}
                                                label="Produto"
                                                onChange={(e) =>
                                                    handleItemChange(index, 'produtoId', e.target.value)
                                                }
                                                required
                                            >
                                                {produtos.map((produto) => (
                                                    <MenuItem key={produto.id} value={produto.id}>
                                                        {produto.nome}
                                                    </MenuItem>
                                                ))}
                                            </Select>
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={5}>
                                        <TextField
                                            fullWidth
                                            type="number"
                                            label="Quantidade"
                                            value={item.quantidade}
                                            onChange={(e) =>
                                                handleItemChange(index, 'quantidade', e.target.value)
                                            }
                                            required
                                        />
                                    </Grid>
                                    <Grid item xs={2}>
                                        <Button
                                            color="error"
                                            onClick={() => removeItem(index)}
                                            disabled={formData.itens.length === 1}
                                        >
                                            Remover
                                        </Button>
                                    </Grid>
                                </Grid>
                            ))}
                            <Button
                                onClick={addItem}
                                startIcon={<AddIcon />}
                                sx={{ mt: 1 }}
                            >
                                Adicionar Item
                            </Button>
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