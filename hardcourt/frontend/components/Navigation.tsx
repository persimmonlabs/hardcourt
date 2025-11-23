"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Trophy, Clock, TrendingUp, Activity } from "lucide-react";
import { clsx } from "clsx";

const navItems = [
    { label: "Live", href: "/", icon: Activity },
    { label: "Past", href: "/past", icon: Clock },
    { label: "Tournaments", href: "/tournaments", icon: Trophy },
    { label: "Stats", href: "/stats", icon: TrendingUp },
];

export const Navigation = () => {
    const pathname = usePathname();

    return (
        <nav className="fixed bottom-6 left-1/2 -translate-x-1/2 w-[90%] max-w-md bg-white/10 backdrop-blur-md border border-white/10 rounded-full p-1 flex shadow-2xl z-50">
            {navItems.map((item) => {
                const Icon = item.icon;
                const isActive = pathname === item.href;

                return (
                    <Link
                        key={item.href}
                        href={item.href}
                        className={clsx(
                            "flex-1 py-3 rounded-full text-sm font-bold transition-all flex items-center justify-center gap-2",
                            isActive
                                ? "bg-neon text-black shadow-lg"
                                : "text-zinc-400 hover:text-white"
                        )}
                    >
                        <Icon size={16} />
                        <span className="hidden sm:inline">{item.label}</span>
                    </Link>
                );
            })}
        </nav>
    );
};
