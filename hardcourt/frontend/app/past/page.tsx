"use client";

import { useState, useEffect } from "react";
import { usePastMatches } from "../../hooks/usePastMatches";
import { MatchCard } from "../../components/MatchCard";
import { Header } from "../../components/Header";
import { Navigation } from "../../components/Navigation";
import { Search, Filter, Calendar } from "lucide-react";

export default function PastMatchesPage() {
    const { matches, isLoading } = usePastMatches({ limit: 50 });
    const [favorites, setFavorites] = useState<string[]>([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [selectedSurface, setSelectedSurface] = useState<string>("all");

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

    const filteredMatches = matches.filter((match) => {
        const matchesSearch =
            searchTerm === "" ||
            match.player1.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
            match.player2.name.toLowerCase().includes(searchTerm.toLowerCase());

        return matchesSearch;
    });

    // Group by date
    const groupedByDate = filteredMatches.reduce((acc, match) => {
        const date = match.start_time
            ? new Date(match.start_time).toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                  year: "numeric",
              })
            : "Recent";
        if (!acc[date]) acc[date] = [];
        acc[date].push(match);
        return acc;
    }, {} as Record<string, typeof matches>);

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
            <Header />

            {/* Search and Filters */}
            <div className="mb-6 space-y-3">
                <div className="relative">
                    <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-zinc-500" size={18} />
                    <input
                        type="text"
                        placeholder="Search players..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="w-full bg-surface border border-white/10 rounded-xl pl-10 pr-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-neon/50 transition-colors"
                    />
                </div>

                <div className="flex gap-2 overflow-x-auto pb-2">
                    {["all", "Hard", "Clay", "Grass"].map((surface) => (
                        <button
                            key={surface}
                            onClick={() => setSelectedSurface(surface)}
                            className={`px-4 py-2 rounded-full text-sm font-bold whitespace-nowrap transition-colors ${
                                selectedSurface === surface
                                    ? "bg-neon text-black"
                                    : "bg-surface text-zinc-400 hover:text-white"
                            }`}
                        >
                            {surface}
                        </button>
                    ))}
                </div>
            </div>

            {/* Content */}
            <div className="space-y-6">
                {isLoading ? (
                    <div className="text-center py-20">
                        <div className="w-12 h-12 border-4 border-neon border-t-transparent rounded-full animate-spin mx-auto mb-3" />
                        <p className="text-zinc-500">Loading past matches...</p>
                    </div>
                ) : filteredMatches.length === 0 ? (
                    <div className="text-center py-20">
                        <p className="text-xl font-bold text-zinc-400">No matches found</p>
                        <p className="text-sm text-zinc-600 mt-2">
                            Try adjusting your search or filters
                        </p>
                    </div>
                ) : (
                    Object.entries(groupedByDate).map(([date, dateMatches]) => (
                        <div key={date} className="space-y-3">
                            <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2">
                                <Calendar size={14} />
                                {date}
                            </h2>
                            <div className="grid gap-3 md:grid-cols-2 lg:grid-cols-2">
                                {dateMatches.map((match) => (
                                    <MatchCard
                                        key={match.id}
                                        match={match}
                                        isFavorite={favorites.includes(match.id)}
                                        onToggleFavorite={toggleFavorite}
                                    />
                                ))}
                            </div>
                        </div>
                    ))
                )}
            </div>

            <Navigation />
        </main>
    );
}
