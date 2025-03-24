'use client';

import { ThemeProvider } from '@mui/material/styles';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { AuthProvider } from '@/contexts/AuthContext';
import { theme } from '@/app/theme';

export default function Providers({ children }) {
    return (
        <ThemeProvider theme={theme}>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
                <AuthProvider>{children}</AuthProvider>
            </LocalizationProvider>
        </ThemeProvider>
    );
} 