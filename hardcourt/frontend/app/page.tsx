"use client";

import { useEffect, useState } from "react";
import { useLiveScores, Match } from "../hooks/useLiveScores";
import { MatchCard } from "../components/MatchCard";
import { Trophy, Star } from "lucide-react";

export default function Home() {
    const { matches, isConnected, isLoading } = useLiveScores();
    const [favorites, setFavorites] = useState<string[]>([]);
    const [activeTab, setActiveTab] = useState<'all' | 'favorites'>('all');

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

    const displayedMatches = activeTab === 'favorites'
        ? matchList.filter(m => favorites.includes(m.id))
        : matchList;

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl">
            {/* Header */}
            <header className="mb-8 flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-black tracking-tighter italic text-white">
                        HARD<span className="text-neon">COURT</span>
                    </h1>
                    <div className="flex items-center gap-2 text-xs text-zinc-500 mt-1">
                        <span className={`w-2 h-2 rounded-full ${isConnected ? 'bg-green-500' : 'bg-red-500'}`} />
                        {isConnected ? 'LIVE FEED ACTIVE' : 'CONNECTING...'}
                    </div>
                </div>
            </header>

            {/* Content */}
            <div className="space-y-6">
                {displayedMatches.length === 0 && (
                    <div className="text-center py-20">
                        {isLoading ? (
                            <p className="text-zinc-500">Loading matches...</p>
                        ) : activeTab === 'favorites' ? (
                            <p className="text-zinc-600">No favorites pinned.</p>
                        ) : (
                            <div className="space-y-2">
                                <p className="text-xl font-bold text-zinc-400">No current live matches</p>
                                <p className="text-sm text-zinc-600">Check back during ATP/WTA tournament hours</p>
                            </div>
                        )}
                    </div>
                )}

                {Object.entries(groupedMatches).map(([tournament, tMatches]) => {
                    const filteredTMatches = activeTab === 'favorites'
                        ? tMatches.filter(m => favorites.includes(m.id))
                        : tMatches;

                    if (filteredTMatches.length === 0) return null;

                    return (
                        <div key={tournament} className="space-y-3">
                            <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2">
                                <Trophy size={14} />
                                {tournament}
                            </h2>
                            {filteredTMatches.map((match) => (
                                <MatchCard
                                    key={match.id}
                                    match={match}
                                    isFavorite={favorites.includes(match.id)}
                                    onToggleFavorite={toggleFavorite}
                                />
                            ))}
                        </div>
                    );
                })}
            </div>

            {/* Bottom Nav */}
            <div className="fixed bottom-6 left-1/2 -translate-x-1/2 w-[90%] max-w-md bg-white/10 backdrop-blur-md border border-white/10 rounded-full p-1 flex shadow-2xl z-50">
                <button
                    onClick={() => setActiveTab('all')}
                    className={`flex-1 py-3 rounded-full text-sm font-bold transition-all ${activeTab === 'all' ? 'bg-neon text-black shadow-lg' : 'text-zinc-400 hover:text-white'
                        }`}
                >
                    Scores
                </button>
                <button
                    onClick={() => setActiveTab('favorites')}
                    className={`flex-1 py-3 rounded-full text-sm font-bold transition-all flex items-center justify-center gap-2 ${activeTab === 'favorites' ? 'bg-neon text-black shadow-lg' : 'text-zinc-400 hover:text-white'
                        }`}
                >
                    <Star size={14} fill={activeTab === 'favorites' ? "currentColor" : "none"} />
                    Favorites
                </button>
            </div>
        </main>
    );
}
