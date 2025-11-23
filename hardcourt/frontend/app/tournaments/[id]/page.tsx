"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";
import { Header } from "../../../components/Header";
import { Navigation } from "../../../components/Navigation";
import { ArrowLeft, Trophy, MapPin, Calendar, DollarSign, Users } from "lucide-react";
import { clsx } from "clsx";

const surfaceColors = {
    Hard: "bg-blue-500/20 text-blue-400 border-blue-500/30",
    Clay: "bg-orange-500/20 text-orange-400 border-orange-500/30",
    Grass: "bg-green-500/20 text-green-400 border-green-500/30",
};

export default function TournamentDetailPage() {
    const params = useParams();
    const tournamentId = params.id;

    const [tournament, setTournament] = useState<any>(null);
    const [matches, setMatches] = useState<any[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [activeTab, setActiveTab] = useState<"matches" | "draw">("matches");

    useEffect(() => {
        const fetchTournamentData = async () => {
            try {
                const backendUrl = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

                // Fetch tournament details
                const tournamentRes = await fetch(`${backendUrl}/api/tournaments/${tournamentId}`);
                if (tournamentRes.ok) {
                    const tournamentData = await tournamentRes.json();
                    setTournament(tournamentData);
                }

                // Fetch tournament matches
                const matchesRes = await fetch(
                    `${backendUrl}/api/tournaments/${tournamentId}/matches`
                );
                if (matchesRes.ok) {
                    const matchesData = await matchesRes.json();
                    setMatches(matchesData.matches || []);
                }
            } catch (error) {
                console.error("Failed to fetch tournament data:", error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchTournamentData();
    }, [tournamentId]);

    const formatPrizeMoney = (amount?: number) => {
        if (!amount) return "N/A";
        return `$${(amount / 1000000).toFixed(1)}M`;
    };

    if (isLoading) {
        return (
            <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
                <Header />
                <div className="text-center py-20">
                    <div className="w-12 h-12 border-4 border-neon border-t-transparent rounded-full animate-spin mx-auto mb-3" />
                    <p className="text-zinc-500">Loading tournament...</p>
                </div>
                <Navigation />
            </main>
        );
    }

    if (!tournament) {
        return (
            <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
                <Header />
                <div className="text-center py-20">
                    <p className="text-xl font-bold text-zinc-400">Tournament not found</p>
                </div>
                <Navigation />
            </main>
        );
    }

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
            {/* Back Button */}
            <Link
                href="/tournaments"
                className="inline-flex items-center gap-2 text-zinc-400 hover:text-white transition-colors mb-4"
            >
                <ArrowLeft size={16} />
                <span>Back to Tournaments</span>
            </Link>

            {/* Tournament Header */}
            <div className="bg-surface border border-white/10 rounded-2xl p-6 mb-6">
                <div className="flex justify-between items-start mb-4">
                    <div className="flex-1">
                        <h1 className="text-2xl font-black text-white mb-2">{tournament.name}</h1>
                        <div className="flex items-center gap-2 text-zinc-400">
                            <MapPin size={16} />
                            <span>
                                {tournament.city}, {tournament.country}
                            </span>
                        </div>
                    </div>
                    <div
                        className={clsx(
                            "px-4 py-2 rounded-full text-sm font-bold border",
                            surfaceColors[tournament.surface as keyof typeof surfaceColors] ||
                                "bg-zinc-500/20 text-zinc-400"
                        )}
                    >
                        {tournament.surface}
                    </div>
                </div>

                <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mt-4">
                    <div className="bg-white/5 p-3 rounded-lg">
                        <div className="flex items-center gap-2 text-neon text-sm mb-1">
                            <Trophy size={14} />
                            Category
                        </div>
                        <div className="text-white font-bold">{tournament.category}</div>
                    </div>
                    <div className="bg-white/5 p-3 rounded-lg">
                        <div className="flex items-center gap-2 text-neon text-sm mb-1">
                            <DollarSign size={14} />
                            Prize Money
                        </div>
                        <div className="text-white font-bold">
                            {formatPrizeMoney(tournament.prize_money)}
                        </div>
                    </div>
                    <div className="bg-white/5 p-3 rounded-lg">
                        <div className="flex items-center gap-2 text-neon text-sm mb-1">
                            <Calendar size={14} />
                            Dates
                        </div>
                        <div className="text-white font-bold text-xs">
                            {tournament.start_date &&
                                new Date(tournament.start_date).toLocaleDateString("en-US", {
                                    month: "short",
                                    day: "numeric",
                                })}
                            {tournament.end_date && " - "}
                            {tournament.end_date &&
                                new Date(tournament.end_date).toLocaleDateString("en-US", {
                                    month: "short",
                                    day: "numeric",
                                })}
                        </div>
                    </div>
                    <div className="bg-white/5 p-3 rounded-lg">
                        <div className="flex items-center gap-2 text-neon text-sm mb-1">
                            <Users size={14} />
                            Status
                        </div>
                        <div className="text-white font-bold capitalize">{tournament.status}</div>
                    </div>
                </div>
            </div>

            {/* Tabs */}
            <div className="flex gap-2 mb-6">
                <button
                    onClick={() => setActiveTab("matches")}
                    className={clsx(
                        "flex-1 py-3 rounded-xl text-sm font-bold transition-colors",
                        activeTab === "matches"
                            ? "bg-neon text-black"
                            : "bg-surface text-zinc-400 hover:text-white border border-white/10"
                    )}
                >
                    Matches
                </button>
                <button
                    onClick={() => setActiveTab("draw")}
                    className={clsx(
                        "flex-1 py-3 rounded-xl text-sm font-bold transition-colors",
                        activeTab === "draw"
                            ? "bg-neon text-black"
                            : "bg-surface text-zinc-400 hover:text-white border border-white/10"
                    )}
                >
                    Draw
                </button>
            </div>

            {/* Content */}
            <div>
                {activeTab === "matches" && (
                    <div className="space-y-3">
                        {matches.length === 0 ? (
                            <div className="text-center py-10 bg-surface border border-white/10 rounded-xl">
                                <p className="text-zinc-400">No matches available</p>
                            </div>
                        ) : (
                            matches.map((match) => (
                                <div
                                    key={match.id}
                                    className="bg-surface border border-white/10 rounded-xl p-4"
                                >
                                    <p className="text-zinc-400 text-sm">{match.player1_name}</p>
                                    <p className="text-zinc-400 text-sm">vs</p>
                                    <p className="text-zinc-400 text-sm">{match.player2_name}</p>
                                </div>
                            ))
                        )}
                    </div>
                )}

                {activeTab === "draw" && (
                    <div className="text-center py-10 bg-surface border border-white/10 rounded-xl">
                        <p className="text-zinc-400">Tournament draw coming soon</p>
                    </div>
                )}
            </div>

            <Navigation />
        </main>
    );
}
