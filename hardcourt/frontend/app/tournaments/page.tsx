"use client";

import { useTournaments } from "../../hooks/useTournaments";
import { Header } from "../../components/Header";
import { Navigation } from "../../components/Navigation";
import { Trophy, MapPin, Calendar, DollarSign, Circle } from "lucide-react";
import { motion } from "framer-motion";
import Link from "next/link";
import { clsx } from "clsx";

const surfaceColors = {
    Hard: "bg-blue-500/20 text-blue-400 border-blue-500/30",
    Clay: "bg-orange-500/20 text-orange-400 border-orange-500/30",
    Grass: "bg-green-500/20 text-green-400 border-green-500/30",
};

export default function TournamentsPage() {
    const { tournaments, isLoading } = useTournaments();

    // Group by status
    const ongoing = tournaments.filter((t) => t.status === "ongoing");
    const upcoming = tournaments.filter((t) => t.status === "upcoming");
    const completed = tournaments.filter((t) => t.status === "completed");

    const formatPrizeMoney = (amount?: number) => {
        if (!amount) return "N/A";
        return `$${(amount / 1000000).toFixed(1)}M`;
    };

    const TournamentCard = ({ tournament }: { tournament: any }) => (
        <Link href={`/tournaments/${tournament.id}`}>
            <motion.div
                whileHover={{ y: -4 }}
                className="bg-surface border border-white/10 rounded-xl p-5 hover:border-neon/50 transition-colors cursor-pointer"
            >
                <div className="flex justify-between items-start mb-3">
                    <div className="flex-1">
                        <h3 className="text-lg font-bold text-white mb-1">{tournament.name}</h3>
                        <div className="flex items-center gap-2 text-sm text-zinc-500">
                            <MapPin size={14} />
                            {tournament.city}, {tournament.country}
                        </div>
                    </div>
                    <div
                        className={clsx(
                            "px-3 py-1 rounded-full text-xs font-bold border",
                            surfaceColors[tournament.surface as keyof typeof surfaceColors] ||
                                "bg-zinc-500/20 text-zinc-400"
                        )}
                    >
                        {tournament.surface}
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-3 text-sm">
                    <div className="flex items-center gap-2 text-zinc-400">
                        <Trophy size={14} className="text-neon" />
                        <span>{tournament.category}</span>
                    </div>
                    <div className="flex items-center gap-2 text-zinc-400">
                        <DollarSign size={14} className="text-neon" />
                        <span>{formatPrizeMoney(tournament.prize_money)}</span>
                    </div>
                </div>

                {tournament.start_date && (
                    <div className="mt-3 pt-3 border-t border-white/10 flex items-center gap-2 text-xs text-zinc-500">
                        <Calendar size={12} />
                        <span>
                            {new Date(tournament.start_date).toLocaleDateString("en-US", {
                                month: "short",
                                day: "numeric",
                            })}
                            {tournament.end_date &&
                                ` - ${new Date(tournament.end_date).toLocaleDateString("en-US", {
                                    month: "short",
                                    day: "numeric",
                                })}`}
                        </span>
                    </div>
                )}
            </motion.div>
        </Link>
    );

    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
            <Header />

            {/* Content */}
            <div className="space-y-8">
                {isLoading ? (
                    <div className="text-center py-20">
                        <div className="w-12 h-12 border-4 border-neon border-t-transparent rounded-full animate-spin mx-auto mb-3" />
                        <p className="text-zinc-500">Loading tournaments...</p>
                    </div>
                ) : (
                    <>
                        {ongoing.length > 0 && (
                            <section>
                                <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2 mb-4">
                                    <Circle className="fill-green-500 text-green-500" size={12} />
                                    Live Now ({ongoing.length})
                                </h2>
                                <div className="grid gap-4 md:grid-cols-2">
                                    {ongoing.map((tournament) => (
                                        <TournamentCard key={tournament.id} tournament={tournament} />
                                    ))}
                                </div>
                            </section>
                        )}

                        {upcoming.length > 0 && (
                            <section>
                                <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2 mb-4">
                                    <Circle className="fill-blue-500 text-blue-500" size={12} />
                                    Upcoming ({upcoming.length})
                                </h2>
                                <div className="grid gap-4 md:grid-cols-2">
                                    {upcoming.map((tournament) => (
                                        <TournamentCard key={tournament.id} tournament={tournament} />
                                    ))}
                                </div>
                            </section>
                        )}

                        {completed.length > 0 && (
                            <section>
                                <h2 className="text-sm font-bold text-zinc-400 uppercase tracking-widest flex items-center gap-2 mb-4">
                                    <Circle className="fill-zinc-600 text-zinc-600" size={12} />
                                    Completed ({completed.length})
                                </h2>
                                <div className="grid gap-4 md:grid-cols-2">
                                    {completed.map((tournament) => (
                                        <TournamentCard key={tournament.id} tournament={tournament} />
                                    ))}
                                </div>
                            </section>
                        )}

                        {tournaments.length === 0 && (
                            <div className="text-center py-20">
                                <p className="text-xl font-bold text-zinc-400">No tournaments available</p>
                            </div>
                        )}
                    </>
                )}
            </div>

            <Navigation />
        </main>
    );
}
