import { Inter } from 'next/font/google';
import { AppRouterCacheProvider } from '@mui/material-nextjs/v13-appRouter';
import CssBaseline from '@mui/material/CssBaseline';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Providers from '@/components/Providers';

const inter = Inter({ subsets: ['latin'] });

export const metadata = {
    title: 'Sistema de Vendas',
    description: 'Sistema de gerenciamento de vendas',
};

export default function RootLayout({ children }) {
    return (
        <html lang="pt-BR">
            <body className={inter.className}>
                <AppRouterCacheProvider>
                    <CssBaseline />
                    <Providers>
                        {children}
                        <ToastContainer
                            position="top-right"
                            autoClose={5000}
                            hideProgressBar={false}
                            newestOnTop
                            closeOnClick
                            rtl={false}
                            pauseOnFocusLoss
                            draggable
                            pauseOnHover
                            theme="dark"
                        />
                    </Providers>
                </AppRouterCacheProvider>
            </body>
        </html>
    );
} 