// lib/types.ts
type UnifiedEpisode = {
    arc: number;
    episode: number;
    title: string;
    description: string;
    released: string;
    crc32: string;
    downloadStatus: "imported" | "missing" | "upgradable" | "queued" | "downloading" | "failed" | "none";
    monitored: boolean;
    version: "normal" | "extended";
    filePath?: string | null;
}

type UnifiedArc = {
    arc: number;
    title: string;
    audioLanguages: string;
    subtitleLanguages: string;
    resolution: string;
    status: string;
    episodeCount: number;
    episodesDownloaded: number;
    episodes: UnifiedEpisode[];
}
