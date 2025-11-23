"use client";

import { motion, AnimatePresence } from "framer-motion";
import { useState, useEffect } from "react";
import { ChevronDown, Star, Activity, Zap } from "lucide-react";
import { clsx } from "clsx";
import { Match } from "../hooks/useLiveScores";

interface MatchCardProps {
    match: Match;
    isFavorite: boolean;
    onToggleFavorite: (id: string) => void;
}

export const MatchCard = ({ match, isFavorite, onToggleFavorite }: MatchCardProps) => {
    const [expanded, setExpanded] = useState(false);

    // High leverage pulse effect
    const isHighLeverage = match.leverage_index > 0.3;

    return (
        <motion.div
            layout
            className={clsx(
                "bg-surface rounded-xl overflow-hidden border transition-colors duration-300",
                isHighLeverage ? "border-red-500/50 shadow-[0_0_15px_rgba(239,68,68,0.2)]" : "border-white/5"
            )}
        >
            {/* Header / Main Row */}
            <div
                className="p-4 cursor-pointer hover:bg-surface-hover transition-colors"
                onClick={() => setExpanded(!expanded)}
            >
                <div className="flex justify-between items-center mb-2">
                    <span className="text-xs font-bold text-zinc-500 uppercase tracking-wider flex items-center gap-2">
                        {match.status === "Live" && (
                            <span className="flex h-2 w-2 relative">
                                <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-neon opacity-75"></span>
                                <span className="relative inline-flex rounded-full h-2 w-2 bg-neon"></span>
                            </span>
                        )}
                        {match.status}
                    </span>
                    <button
                        onClick={(e) => {
                            e.stopPropagation();
                            onToggleFavorite(match.id);
                        }}
                        className={clsx("transition-colors", isFavorite ? "text-neon" : "text-zinc-600 hover:text-zinc-400")}
                    >
                        <Star size={16} fill={isFavorite ? "currentColor" : "none"} />
                    </button>
                </div>

                {/* Players & Scores */}
                <div className="grid grid-cols-[1fr_auto] gap-4">
                    <div className="space-y-2">
                        <PlayerRow
                            name={match.player1.name}
                            isServing={match.score.serving === 1}
                            sets={match.score.sets_p1}
                            games={match.score.games_p1}
                            points={match.score.points_p1}
                            winner={match.win_prob_p1 > 0.5}
                        />
                        <PlayerRow
                            name={match.player2.name}
                            isServing={match.score.serving === 2}
                            sets={match.score.sets_p2}
                            games={match.score.games_p2}
                            points={match.score.points_p2}
                            winner={match.win_prob_p1 <= 0.5}
                        />
                    </div>
                </div>
            </div>

            {/* Expanded Stats (Accordion) */}
            <AnimatePresence>
                {expanded && (
                    <motion.div
                        initial={{ height: 0, opacity: 0 }}
                        animate={{ height: "auto", opacity: 1 }}
                        exit={{ height: 0, opacity: 0 }}
                        className="bg-black/20 border-t border-white/5"
                    >
                        <div className="p-4 space-y-4">
                            {/* Win Probability Bar */}
                            <div>
                                <div className="flex justify-between text-xs text-zinc-400 mb-1">
                                    <span>Win Probability</span>
                                    <span>{(match.win_prob_p1 * 100).toFixed(1)}%</span>
                                </div>
                                <div className="h-2 bg-zinc-800 rounded-full overflow-hidden flex">
                                    <motion.div
                                        className="h-full bg-neon"
                                        animate={{ width: `${match.win_prob_p1 * 100}%` }}
                                        transition={{ type: "spring", stiffness: 50 }}
                                    />
                                </div>
                            </div>

                            {/* Leverage & Fatigue */}
                            <div className="grid grid-cols-2 gap-4">
                                <StatBox
                                    label="Leverage Idx"
                                    value={match.leverage_index.toFixed(2)}
                                    icon={<Zap size={14} className={isHighLeverage ? "text-red-500" : "text-zinc-500"} />}
                                    highlight={isHighLeverage}
                                />
                                <StatBox
                                    label="Rally Count"
                                    value={match.stats.rally_count.toString()}
                                    icon={<Activity size={14} className="text-zinc-500" />}
                                />
                            </div>

                            {/* Deep Stats Grid */}
                            <div className="grid grid-cols-3 text-center text-xs text-zinc-400 pt-2 border-t border-white/5">
                                <div>
                                    <div className="text-white font-mono">{match.stats.aces_p1}</div>
                                    <div>Aces</div>
                                    <div className="text-white font-mono">{match.stats.aces_p2}</div>
                                </div>
                                <div>
                                    <div className="text-white font-mono">{match.stats.df_p1}</div>
                                    <div>DFs</div>
                                    <div className="text-white font-mono">{match.stats.df_p2}</div>
                                </div>
                                <div>
                                    <div className="text-white font-mono">{Math.round(match.fatigue_p1)}%</div>
                                    <div>Fatigue</div>
                                    <div className="text-white font-mono">{Math.round(match.fatigue_p2)}%</div>
                                </div>
                            </div>
                        </div>
                    </motion.div>
                )}
            </AnimatePresence>
        </motion.div>
    );
};

const PlayerRow = ({ name, isServing, sets, games, points, winner }: any) => (
    <div className="flex justify-between items-center">
        <div className="flex items-center gap-2">
            <span className={clsx("w-2 h-2 rounded-full", isServing ? "bg-neon" : "bg-transparent")} />
            <span className={clsx("font-medium", winner ? "text-white" : "text-zinc-400")}>{name}</span>
        </div>
        <div className="flex gap-4 font-mono">
            <span className="text-zinc-500 w-4 text-center">{sets}</span>
            <span className="text-white w-4 text-center">{games}</span>
            <span className="text-neon w-6 text-right font-bold">{points}</span>
        </div>
    </div>
);

const StatBox = ({ label, value, icon, highlight }: any) => (
    <div className={clsx("bg-white/5 p-2 rounded-lg flex flex-col items-center justify-center", highlight && "bg-red-500/10 border border-red-500/20")}>
        <div className="flex items-center gap-1 text-xs text-zinc-400 mb-1">
            {icon}
            {label}
        </div>
        <div className={clsx("text-lg font-mono font-bold", highlight ? "text-red-400" : "text-white")}>{value}</div>
    </div>
);
