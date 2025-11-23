"use client";

import { useTournaments } from "../../hooks/useTournaments";
import { Header } from "../../components/Header";
import { Navigation } from "../../components/Navigation";
import { Trophy, MapPin, Calendar, DollarSign, Circle, Award, ChevronDown } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import Link from "next/link";
import { clsx } from "clsx";
import { useState, useMemo } from "react";

const surfaceColors = {
    Hard: "bg-blue-500/20 text-blue-400 border-blue-500/30",
    Clay: "bg-orange-500/20 text-orange-400 border-orange-500/30",
    Grass: "bg-green-500/20 text-green-400 border-green-500/30",
};

const categoryColors = {
    "Grand Slam": "from-yellow-500/20 to-orange-500/20",
    "Masters 1000": "from-purple-500/20 to-pink-500/20",
    "ATP Finals": "from-blue-500/20 to-cyan-500/20",
    "ATP 500": "from-green-500/20 to-emerald-500/20",
    "ATP 250": "from-zinc-500/20 to-gray-500/20",
};

export default function TournamentsPage() {
    const { tournaments, isLoading } = useTournaments();
    const [selectedYear, setSelectedYear] = useState<number | "all">("all");
    const [selectedStatus, setSelectedStatus] = useState<string | "all">("all");
    const [isYearDropdownOpen, setIsYearDropdownOpen] = useState(false);

    // Extract unique years and sort descending
    const availableYears = useMemo(() => {
        const years = Array.from(new Set(tournaments.map(t => t.year).filter(Boolean))).sort((a, b) => (b as number) - (a as number));
        return years as number[];
    }, [tournaments]);

    // Filter tournaments by year and status
    const filteredTournaments = useMemo(() => {
        return tournaments.filter(t => {
            const yearMatch = selectedYear === "all" || t.year === selectedYear;
            const statusMatch = selectedStatus === "all" || t.status === selectedStatus;
            return yearMatch && statusMatch;
        });
    }, [tournaments, selectedYear, selectedStatus]);

    // Group by status
    const ongoing = filteredTournaments.filter((t) => t.status === "ongoing");
    const upcoming = filteredTournaments.filter((t) => t.status === "upcoming");
    const completed = filteredTournaments.filter((t) => t.status === "completed");

    // Group completed tournaments by year
    const completedByYear = useMemo(() => {
        const grouped: { [year: number]: typeof completed } = {};
        completed.forEach(t => {
            const year = t.year || new Date(t.start_date || "").getFullYear();
            if (!grouped[year]) grouped[year] = [];
            grouped[year].push(t);
        });
        return grouped;
    }, [completed]);

    const formatPrizeMoney = (amount?: number) => {
        if (!amount) return "N/A";
        return `$${(amount / 1000000).toFixed(1)}M`;
    };

    const TournamentCard = ({ tournament }: { tournament: any }) => {
        const isCompleted = tournament.status === "completed";
        const categoryGradient = categoryColors[tournament.category as keyof typeof categoryColors] || "from-zinc-500/20 to-gray-500/20";

        return (
            <Link href={`/tournaments/${tournament.id}`}>
                <motion.div
                    whileHover={{ y: -4, scale: 1.02 }}
                    className={clsx(
                        "relative bg-surface border border-white/10 rounded-xl p-5 overflow-hidden",
                        "hover:border-neon/50 transition-all cursor-pointer group"
                    )}
                >
                    {/* Background Gradient */}
                    <div className={clsx("absolute inset-0 bg-gradient-to-br opacity-30 group-hover:opacity-50 transition-opacity", categoryGradient)} />

                    {/* Content */}
                    <div className="relative z-10">
                        {/* Header */}
                        <div className="flex justify-between items-start mb-3">
                            <div className="flex-1 pr-2">
                                <div className="flex items-center gap-2 mb-1">
                                    <h3 className="text-lg font-bold text-white">{tournament.name}</h3>
                                    {tournament.year && (
                                        <span className="text-xs text-zinc-500 font-mono">{tournament.year}</span>
                                    )}
                                </div>
                                <div className="flex items-center gap-2 text-sm text-zinc-400">
                                    <MapPin size={14} />
                                    <span>{tournament.city}, {tournament.country}</span>
                                </div>
                            </div>
                            <div
                                className={clsx(
                                    "px-3 py-1 rounded-full text-xs font-bold border shrink-0",
                                    surfaceColors[tournament.surface as keyof typeof surfaceColors] ||
                                        "bg-zinc-500/20 text-zinc-400"
                                )}
                            >
                                {tournament.surface}
                            </div>
                        </div>

                        {/* Winner Info for Completed Tournaments */}
                        {isCompleted && tournament.winner_id && (
                            <div className="mb-3 p-3 bg-neon/10 border border-neon/30 rounded-lg">
                                <div className="flex items-center gap-2 text-neon text-sm font-bold">
                                    <Award size={16} className="text-yellow-400" />
                                    <span>Champion: {tournament.winner_id.replace(/-/g, " ").toUpperCase()}</span>
                                </div>
                                {tournament.runner_up_id && (
                                    <div className="mt-1 text-xs text-zinc-400 ml-6">
                                        Runner-up: {tournament.runner_up_id.replace(/-/g, " ").toUpperCase()}
                                    </div>
                                )}
                            </div>
                        )}

                        {/* Tournament Info */}
                        <div className="grid grid-cols-2 gap-3 text-sm">
                            <div className="flex items-center gap-2 text-zinc-400">
                                <Trophy size={14} className="text-neon" />
                                <span className="truncate">{tournament.category}</span>
                            </div>
                            <div className="flex items-center gap-2 text-zinc-400">
                                <DollarSign size={14} className="text-neon" />
                                <span>{formatPrizeMoney(tournament.prize_money)}</span>
                            </div>
                        </div>

                        {/* Dates */}
                        {tournament.start_date && (
                            <div className="mt-3 pt-3 border-t border-white/10 flex items-center gap-2 text-xs text-zinc-500">
                                <Calendar size={12} />
                                <span>
                                    {new Date(tournament.start_date).toLocaleDateString("en-US", {
                                        month: "short",
                                        day: "numeric",
                                        year: "numeric",
                                    })}
                                    {tournament.end_date &&
                                        ` - ${new Date(tournament.end_date).toLocaleDateString("en-US", {
                                            month: "short",
                                            day: "numeric",
                                        })}`}
                                </span>
                            </div>
                        )}

                        {/* Status Badge */}
                        {tournament.status === "ongoing" && (
                            <div className="absolute top-3 right-3 flex items-center gap-1 px-2 py-1 bg-green-500/20 border border-green-500/30 rounded-full">
                                <Circle className="fill-green-500 text-green-500 animate-pulse" size={8} />
                                <span className="text-xs text-green-400 font-bold">LIVE</span>
                            </div>
                        )}
                    </div>
                </motion.div>
            </Link>
        );
    };

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-6xl">
            <Header />

            {/* Filters */}
            <div className="mb-8 flex gap-3 flex-wrap">
                {/* Year Filter */}
                <div className="relative">
                    <button
                        onClick={() => setIsYearDropdownOpen(!isYearDropdownOpen)}
                        className="px-4 py-2 bg-surface border border-white/10 rounded-lg text-sm font-bold text-white hover:border-neon/50 transition-colors flex items-center gap-2"
                    >
                        <Calendar size={16} />
                        <span>{selectedYear === "all" ? "All Years" : selectedYear}</span>
                        <ChevronDown size={16} className={clsx("transition-transform", isYearDropdownOpen && "rotate-180")} />
                    </button>

                    <AnimatePresence>
                        {isYearDropdownOpen && (
                            <motion.div
                                initial={{ opacity: 0, y: -10 }}
                                animate={{ opacity: 1, y: 0 }}
                                exit={{ opacity: 0, y: -10 }}
                                className="absolute top-full mt-2 w-48 bg-surface border border-white/10 rounded-lg shadow-2xl overflow-hidden z-50"
                            >
                                <button
                                    onClick={() => { setSelectedYear("all"); setIsYearDropdownOpen(false); }}
                                    className={clsx(
                                        "w-full px-4 py-2 text-left text-sm hover:bg-neon/10 transition-colors",
                                        selectedYear === "all" ? "text-neon font-bold" : "text-white"
                                    )}
                                >
                                    All Years
                                </button>
                                {availableYears.map(year => (
                                    <button
                                        key={year}
                                        onClick={() => { setSelectedYear(year); setIsYearDropdownOpen(false); }}
                                        className={clsx(
                                            "w-full px-4 py-2 text-left text-sm hover:bg-neon/10 transition-colors",
                                            selectedYear === year ? "text-neon font-bold" : "text-white"
                                        )}
                                    >
                                        {year}
                                    </button>
                                ))}
                            </motion.div>
                        )}
                    </AnimatePresence>
                </div>

                {/* Status Filters */}
                <button
                    onClick={() => setSelectedStatus("all")}
                    className={clsx(
                        "px-4 py-2 rounded-lg text-sm font-bold transition-all",
                        selectedStatus === "all"
                            ? "bg-neon/20 text-neon border border-neon/50"
                            : "bg-surface border border-white/10 text-white hover:border-neon/30"
                    )}
                >
                    All ({filteredTournaments.length})
                </button>
                <button
                    onClick={() => setSelectedStatus("ongoing")}
                    className={clsx(
                        "px-4 py-2 rounded-lg text-sm font-bold transition-all flex items-center gap-2",
                        selectedStatus === "ongoing"
                            ? "bg-green-500/20 text-green-400 border border-green-500/50"
                            : "bg-surface border border-white/10 text-white hover:border-green-500/30"
                    )}
                >
                    <Circle className="fill-green-500 text-green-500" size={10} />
                    Live ({ongoing.length})
                </button>
                <button
                    onClick={() => setSelectedStatus("upcoming")}
                    className={clsx(
                        "px-4 py-2 rounded-lg text-sm font-bold transition-all flex items-center gap-2",
                        selectedStatus === "upcoming"
                            ? "bg-blue-500/20 text-blue-400 border border-blue-500/50"
                            : "bg-surface border border-white/10 text-white hover:border-blue-500/30"
                    )}
                >
                    <Circle className="fill-blue-500 text-blue-500" size={10} />
                    Upcoming ({upcoming.length})
                </button>
                <button
                    onClick={() => setSelectedStatus("completed")}
                    className={clsx(
                        "px-4 py-2 rounded-lg text-sm font-bold transition-all flex items-center gap-2",
                        selectedStatus === "completed"
                            ? "bg-zinc-500/20 text-zinc-300 border border-zinc-500/50"
                            : "bg-surface border border-white/10 text-white hover:border-zinc-500/30"
                    )}
                >
                    <Trophy size={14} />
                    Completed ({completed.length})
                </button>
            </div>

            {/* Content */}
            <div className="space-y-8">
                {isLoading ? (
                    <div className="text-center py-20">
                        <div className="w-12 h-12 border-4 border-neon border-t-transparent rounded-full animate-spin mx-auto mb-3" />
                        <p className="text-zinc-500">Loading tournaments...</p>
                    </div>
                ) : (
                    <>
                        {/* Live Tournaments */}
                        {ongoing.length > 0 && (selectedStatus === "all" || selectedStatus === "ongoing") && (
                            <section>
                                <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2 mb-4">
                                    <Circle className="fill-green-500 text-green-500 animate-pulse" size={12} />
                                    Live Now ({ongoing.length})
                                </h2>
                                <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                                    {ongoing.map((tournament) => (
                                        <TournamentCard key={tournament.id} tournament={tournament} />
                                    ))}
                                </div>
                            </section>
                        )}

                        {/* Upcoming Tournaments */}
                        {upcoming.length > 0 && (selectedStatus === "all" || selectedStatus === "upcoming") && (
                            <section>
                                <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2 mb-4">
                                    <Circle className="fill-blue-500 text-blue-500" size={12} />
                                    Upcoming ({upcoming.length})
                                </h2>
                                <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                                    {upcoming.map((tournament) => (
                                        <TournamentCard key={tournament.id} tournament={tournament} />
                                    ))}
                                </div>
                            </section>
                        )}

                        {/* Completed Tournaments - Grouped by Year */}
                        {completed.length > 0 && (selectedStatus === "all" || selectedStatus === "completed") && (
                            <section>
                                {Object.entries(completedByYear)
                                    .sort(([yearA], [yearB]) => Number(yearB) - Number(yearA))
                                    .map(([year, yearTournaments]) => (
                                        <div key={year} className="mb-10">
                                            <h2 className="text-2xl font-bold text-white mb-6 flex items-center gap-3">
                                                <div className="h-1 w-12 bg-gradient-to-r from-neon to-transparent rounded-full" />
                                                {year}
                                                <span className="text-sm text-zinc-500 font-normal">
                                                    ({yearTournaments.length} tournament{yearTournaments.length !== 1 ? 's' : ''})
                                                </span>
                                            </h2>
                                            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                                                {yearTournaments.map((tournament) => (
                                                    <TournamentCard key={tournament.id} tournament={tournament} />
                                                ))}
                                            </div>
                                        </div>
                                    ))}
                            </section>
                        )}

                        {filteredTournaments.length === 0 && (
                            <div className="text-center py-20">
                                <Trophy size={48} className="mx-auto mb-4 text-zinc-600" />
                                <p className="text-xl font-bold text-zinc-400 mb-2">No tournaments found</p>
                                <p className="text-sm text-zinc-500">
                                    Try selecting a different year or status filter
                                </p>
                            </div>
                        )}
                    </>
                )}
            </div>

            <Navigation />
        </main>
    );
}
