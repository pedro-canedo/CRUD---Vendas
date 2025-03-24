'use client';

import { Box, CircularProgress, Typography } from '@mui/material';

export default function Loading() {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                minHeight: '100vh',
            }}
        >
            <CircularProgress />
            <Typography variant="h6" sx={{ mt: 2 }}>
                Carregando...
            </Typography>
        </Box>
    );
} 