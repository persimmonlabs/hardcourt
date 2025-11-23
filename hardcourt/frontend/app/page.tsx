"use client";

import { useEffect, useState } from "react";
import { useLiveScores, Match } from "../hooks/useLiveScores";
import { MatchCard } from "../components/MatchCard";
import { Header } from "../components/Header";
import { Navigation } from "../components/Navigation";
import { Trophy } from "lucide-react";

export default function Home() {
    const { matches, isConnected, isLoading } = useLiveScores();
    const [favorites, setFavorites] = useState<string[]>([]);

    useEffect(() => {
        const saved = localStorage.getItem("hardcourt_favorites");
        if (saved) {
            setFavorites(JSON.parse(saved));
        }
    }, []);

    const toggleFavorite = (id: string) => {
        const newFavs = favorites.includes(id)
            ? favorites.filter((f) => f !== id)
            : [...favorites, id];
        setFavorites(newFavs);
        localStorage.setItem("hardcourt_favorites", JSON.stringify(newFavs));
    };

    const matchList = Object.values(matches);

    // Group by Tournament
    const groupedMatches = matchList.reduce((acc, match) => {
        const tName = "Live Tournaments"; // Simplified for now, could use match.tournament_id lookup
        if (!acc[tName]) acc[tName] = [];
        acc[tName].push(match);
        return acc;
    }, {} as Record<string, Match[]>);

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
            <Header isConnected={isConnected} />

            {/* Content */}
            <div className="space-y-6">
                {matchList.length === 0 && (
                    <div className="text-center py-20">
                        {isLoading ? (
                            <div className="space-y-3">
                                <div className="w-12 h-12 border-4 border-neon border-t-transparent rounded-full animate-spin mx-auto" />
                                <p className="text-zinc-500">Loading matches...</p>
                            </div>
                        ) : (
                            <div className="space-y-2">
                                <p className="text-xl font-bold text-zinc-400">No current live matches</p>
                                <p className="text-sm text-zinc-600">Check back during ATP/WTA tournament hours</p>
                            </div>
                        )}
                    </div>
                )}

                {Object.entries(groupedMatches).map(([tournament, tMatches]) => (
                    <div key={tournament} className="space-y-3">
                        <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2">
                            <Trophy size={14} />
                            {tournament}
                        </h2>
                        <div className="grid gap-3 md:grid-cols-2 lg:grid-cols-2">
                            {tMatches.map((match) => (
                                <MatchCard
                                    key={match.id}
                                    match={match}
                                    isFavorite={favorites.includes(match.id)}
                                    onToggleFavorite={toggleFavorite}
                                />
                            ))}
                        </div>
                    </div>
                ))}
            </div>

            <Navigation />
        </main>
    );
}
