"use client";

import { Header } from "../../components/Header";
import { Navigation } from "../../components/Navigation";
import { TrendingUp, Users, Globe, Award } from "lucide-react";

export default function StatsPage() {
    return (
        <main className="min-h-screen pb-24 px-4 pt-6 max-w-md mx-auto md:max-w-2xl lg:max-w-4xl">
            <Header />

            {/* Content */}
            <div className="space-y-6">
                <div className="text-center py-20">
                    <div className="bg-surface border border-white/10 rounded-2xl p-8 max-w-md mx-auto">
                        <div className="w-16 h-16 bg-neon/10 rounded-full flex items-center justify-center mx-auto mb-4">
                            <TrendingUp className="text-neon" size={32} />
                        </div>
                        <h2 className="text-2xl font-bold text-white mb-2">Player Stats Coming Soon</h2>
                        <p className="text-zinc-400 mb-6">
                            Advanced player statistics, head-to-head records, and tournament brackets
                            will be available here.
                        </p>

                        <div className="grid grid-cols-2 gap-3 text-sm">
                            <div className="bg-white/5 p-3 rounded-lg">
                                <Users className="text-neon mb-2" size={20} />
                                <div className="text-zinc-500">Head-to-Head</div>
                            </div>
                            <div className="bg-white/5 p-3 rounded-lg">
                                <Globe className="text-neon mb-2" size={20} />
                                <div className="text-zinc-500">Rankings</div>
                            </div>
                            <div className="bg-white/5 p-3 rounded-lg">
                                <Award className="text-neon mb-2" size={20} />
                                <div className="text-zinc-500">Tournament Draws</div>
                            </div>
                            <div className="bg-white/5 p-3 rounded-lg">
                                <TrendingUp className="text-neon mb-2" size={20} />
                                <div className="text-zinc-500">Performance</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <Navigation />
        </main>
    );
}
