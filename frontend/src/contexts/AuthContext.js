'use client';

import { createContext, useContext, useState, useEffect } from 'react';
import authService from '@/services/authService';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';
import Loading from '@/components/Loading';

const AuthContext = createContext({});

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
        const token = Cookies.get('token');
        if (token) {
            authService.getUsuarioAtual()
                .then(response => {
                    setUser(response.data);
                })
                .catch(() => {
                    Cookies.remove('token');
                    setUser(null);
                })
                .finally(() => {
                    setLoading(false);
                });
        } else {
            setLoading(false);
        }
    }, []);

    const login = async (email, senha) => {
        try {
            const response = await authService.login(email, senha);
            const { token, usuario } = response.data;
            Cookies.set('token', token);
            setUser(usuario);
            router.push('/');
        } catch (error) {
            throw error;
        }
    };

    const logout = () => {
        authService.logout();
        Cookies.remove('token');
        setUser(null);
        router.push('/login');
    };

    const updateUser = (userData) => {
        setUser(userData);
    };

    if (loading) {
        return <Loading />;
    }

    return (
        <AuthContext.Provider value={{ user, login, logout, updateUser }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth deve ser usado dentro de um AuthProvider');
    }
    return context;
} 