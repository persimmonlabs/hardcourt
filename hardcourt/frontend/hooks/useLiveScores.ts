import { useEffect, useState, useRef } from 'react';

export type Match = {
    id: string;
    tournament_id: string;
    player1: { name: string; country_code: string; rank: number };
    player2: { name: string; country_code: string; rank: number };
    status: string;
    start_time?: string;
    score: {
        sets_p1: number;
        sets_p2: number;
        games_p1: number;
        games_p2: number;
        points_p1: string;
        points_p2: string;
        serving: number;
    };
    stats: {
        aces_p1: number;
        aces_p2: number;
        df_p1: number;
        df_p2: number;
        rally_count: number;
    };
    win_prob_p1: number;
    leverage_index: number;
    fatigue_p1: number;
    fatigue_p2: number;
};

export const useLiveScores = () => {
    const [matches, setMatches] = useState<Record<string, Match>>({});
    const [isConnected, setIsConnected] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const wsRef = useRef<WebSocket | null>(null);

    useEffect(() => {
        // Determine backend URL
        const backendUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
        const wsUrl = process.env.NEXT_PUBLIC_WS_URL || backendUrl.replace('http', 'ws') + '/ws';

        // Fetch initial match data via REST API
        const fetchInitialData = async () => {
            try {
                const response = await fetch(`${backendUrl}/api/matches?status=Live`);
                if (response.ok) {
                    const data: Match[] = await response.json();
                    const matchesMap = data.reduce((acc, match) => {
                        acc[match.id] = match;
                        return acc;
                    }, {} as Record<string, Match>);
                    setMatches(matchesMap);
                    console.log(`Loaded ${data.length} initial matches`);
                }
            } catch (error) {
                console.error('Failed to fetch initial matches:', error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchInitialData();

        // Connect to WebSocket for live updates
        const ws = new WebSocket(wsUrl);
        wsRef.current = ws;

        ws.onopen = () => {
            console.log('Connected to Live Scores WebSocket');
            setIsConnected(true);
        };

        ws.onmessage = (event: MessageEvent) => {
            try {
                const matchData: Match = JSON.parse(event.data);
                setMatches((prev) => ({
                    ...prev,
                    [matchData.id]: matchData,
                }));
            } catch (e) {
                console.error('Failed to parse match data', e);
            }
        };

        ws.onclose = () => {
            console.log('WebSocket disconnected');
            setIsConnected(false);
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };

        return () => {
            ws.close();
        };
    }, []);

    return { matches, isConnected, isLoading };
};
