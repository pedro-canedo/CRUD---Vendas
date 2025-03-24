import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080/api/v1',
});

// Interceptor para adicionar o token em todas as requisições
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Interceptor para tratar erros de autenticação
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export const authService = {
    login: (email, senha) => api.post('/auth/login', { email, senha }),
    logout: () => api.post('/auth/logout'),
    getUser: () => api.get('/auth/user'),
};

export const produtoService = {
    getAll: () => api.get('/produtos'),
    getById: (id) => api.get(`/produtos/${id}`),
    create: (produto) => api.post('/produtos', produto),
    update: (id, produto) => api.put(`/produtos/${id}`, produto),
    delete: (id) => api.delete(`/produtos/${id}`),
};

export const vendaService = {
    getAll: () => api.get('/vendas'),
    getById: (id) => api.get(`/vendas/${id}`),
    create: (venda) => api.post('/vendas', venda),
    update: (id, venda) => api.put(`/vendas/${id}`, venda),
    delete: (id) => api.delete(`/vendas/${id}`),
    getByCliente: (clienteId) => api.get(`/vendas/cliente/${clienteId}`),
    getByPeriodo: (inicio, fim) =>
        api.get(`/vendas/periodo/${inicio}/${fim}`),
};

export const relatorioService = {
    getRelatorio: async (filtro = 'diario', dataInicio = null, dataFim = null) => {
        const params = new URLSearchParams();
        if (filtro) params.append('filtro', filtro);
        if (dataInicio) params.append('dataInicio', dataInicio.toISOString());
        if (dataFim) params.append('dataFim', dataFim.toISOString());

        const response = await api.get(`/relatorios?${params.toString()}`);
        return response;
    },
};

export const configuracaoService = {
    getConfiguracoes: () => api.get('/configuracoes'),
    updateConfiguracoes: (configuracoes) => api.put('/configuracoes', configuracoes),
};

export const usuarioService = {
    getPerfil: () => api.get('/usuarios/perfil'),
    updatePerfil: (perfil) => api.put('/usuarios/perfil', perfil),
};

export const backupService = {
    getBackups: () => api.get('/backups'),
    createBackup: (descricao) => api.post('/backups', { descricao }),
    downloadBackup: (id) => api.get(`/backups/${id}/download`, { responseType: 'blob' }),
    deleteBackup: (id) => api.delete(`/backups/${id}`),
};

export const logService = {
    getLogs: (params) => api.get('/logs', { params }),
};

export default api; 