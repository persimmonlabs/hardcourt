"use client";

export const Header = ({ isConnected }: { isConnected?: boolean }) => {
    return (
        <header className="mb-8 flex justify-between items-center">
            <div>
                <h1 className="text-3xl font-black tracking-tighter italic text-white">
                    HARD<span className="text-neon">COURT</span>
                </h1>
                {isConnected !== undefined && (
                    <div className="flex items-center gap-2 text-xs text-zinc-500 mt-1">
                        <span
                            className={`w-2 h-2 rounded-full ${
                                isConnected ? "bg-green-500" : "bg-red-500"
                            }`}
                        />
                        {isConnected ? "LIVE FEED ACTIVE" : "CONNECTING..."}
                    </div>
                )}
            </div>
        </header>
    );
};
