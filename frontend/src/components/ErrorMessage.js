import { Alert, AlertTitle, Box } from '@mui/material';

export default function ErrorMessage({ error, title = 'Erro' }) {
    if (!error) return null;

    return (
        <Box sx={{ mb: 2 }}>
            <Alert severity="error">
                <AlertTitle>{title}</AlertTitle>
                {error.message || 'Ocorreu um erro inesperado. Tente novamente.'}
            </Alert>
        </Box>
    );
} 