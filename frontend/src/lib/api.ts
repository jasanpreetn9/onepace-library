// lib/api.ts
export const api = {
    async getAllEpisodes(): Promise<UnifiedArc[]> {
        const res = await fetch("http://localhost:8989/api/episodes/all");
        const resp = await res.json();

        // Convert backend response → UnifiedArc[]
        const unifiedArcs: UnifiedArc[] = resp.map((arc: any) => ({
            arc: arc.arc,
            title: arc.title,
            audio_languages: arc.audio_languages,
            subtitle_languages: arc.subtitle_languages,
            resolution: arc.resolution,
            status: arc.status,
            episodeCount: 12,
            episodesDownloaded: 5,
            episodes: arc.episodes.map((ep: any) => ({
                arc: ep.arc,
                episode: ep.episode,
                title: ep.title,
                description: ep.description,
                released: ep.released,
                crc32: ep.crc32,
                version: ep.version,
                downloadStatus: ep.downloadStatus,
                monitored: ep.monitored,
                filePath: ep.filePath
            })) as UnifiedEpisode[],
        }));

        return unifiedArcs;
    },

    async scanLibrary() {
        return fetch("http://localhost:8989/api/scan/library", { method: "POST" }).then(r => r.json());
    },

    async scanDownloads() {
        return fetch("http://localhost:8989/api/scan/downloads", { method: "POST" }).then(r => r.json());
    },

    async toggleMonitor(arc: number, episode: number, monitored: boolean) {
        return fetch("http://localhost:8989/api/episodes/monitor", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ arc, episode, monitored }),
        }).then(r => r.json());
    },

    async downloadEpisode(crc32: string) {
        return fetch("http://localhost:8989/api/download/add", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ crc32 }),
        }).then(r => r.json());
    }
};
