import axios from './axiosConfig';

const API_URL = process.env.NEXT_PUBLIC_API_URL;

const authService = {
    async login(email, senha) {
        const response = await axios.post('/v1/auth/login', { email, senha });
        return response;
    },

    async getUsuarioAtual() {
        const response = await axios.get('/v1/auth/me');
        return response;
    },

    async logout() {
        localStorage.removeItem('token');
    }
};

export default authService; 