import { useState, useCallback } from 'react';

export function useLoading() {
    const [loading, setLoading] = useState(false);

    const startLoading = useCallback(() => {
        setLoading(true);
    }, []);

    const stopLoading = useCallback(() => {
        setLoading(false);
    }, []);

    const withLoading = useCallback(
        async (fn) => {
            try {
                startLoading();
                await fn();
            } finally {
                stopLoading();
            }
        },
        [startLoading, stopLoading]
    );

    return {
        loading,
        startLoading,
        stopLoading,
        withLoading,
    };
} 