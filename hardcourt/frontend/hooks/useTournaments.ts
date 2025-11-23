import { useState, useEffect } from 'react';

export type Tournament = {
    id: string;
    name: string;
    surface: string;
    city: string;
    country: string;
    category: string;
    status: string;
    start_date?: string;
    end_date?: string;
    prize_money?: number;
    year?: number;
    winner_id?: string;
    runner_up_id?: string;
    logo_url?: string;
};

export const useTournaments = (status?: string) => {
    const [tournaments, setTournaments] = useState<Tournament[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchTournaments = async () => {
            try {
                const backendUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
                const url = status
                    ? `${backendUrl}/api/tournaments?status=${status}`
                    : `${backendUrl}/api/tournaments`;

                const response = await fetch(url);
                if (!response.ok) {
                    throw new Error('Failed to fetch tournaments');
                }

                const data = await response.json();
                setTournaments(data);
            } catch (err) {
                setError(err instanceof Error ? err.message : 'Unknown error');
            } finally {
                setIsLoading(false);
            }
        };

        fetchTournaments();
    }, [status]);

    return { tournaments, isLoading, error };
};
