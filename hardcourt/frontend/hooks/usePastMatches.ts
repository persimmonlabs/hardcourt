import { useState, useEffect } from 'react';
import { Match } from './useLiveScores';

export const usePastMatches = (filters?: {
    player?: string;
    tournament?: string;
    limit?: number;
}) => {
    const [matches, setMatches] = useState<Match[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchPastMatches = async () => {
            try {
                const backendUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
                const params = new URLSearchParams();

                if (filters?.player) params.append('player', filters.player);
                if (filters?.tournament) params.append('tournament', filters.tournament);
                if (filters?.limit) params.append('limit', filters.limit.toString());

                const response = await fetch(
                    `${backendUrl}/api/matches/past?${params.toString()}`
                );

                if (!response.ok) {
                    throw new Error('Failed to fetch past matches');
                }

                const data = await response.json();
                setMatches(data.matches || []);
            } catch (err) {
                setError(err instanceof Error ? err.message : 'Unknown error');
            } finally {
                setIsLoading(false);
            }
        };

        fetchPastMatches();
    }, [filters?.player, filters?.tournament, filters?.limit]);

    return { matches, isLoading, error };
};
