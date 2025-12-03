package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
)

type UnifiedEpisode struct {
	Arc            int    `json:"arc"`
	Episode        int    `json:"episode"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Released       string `json:"released"`
	CRC32          string `json:"crc32"`
	DownloadStatus string `json:"downloadStatus"`
	Monitored      bool   `json:"monitored"`
	FilePath       string `json:"filePath"`
	Version        string `json:"version"`
}

type UnifiedArc struct {
	Arc               int              `json:"arc"`
	Title             string           `json:"title"`
	AudioLanguages    string           `json:"audio_languages"`
	SubtitleLanguages string           `json:"subtitle_languages"`
	Resolution        string           `json:"resolution"`
	Status            string           `json:"status"`
	Episodes          []UnifiedEpisode `json:"episodes"`
}

func HandleGetAllEpisodes(meta *metadata.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lib, err := library.Load("data/library.json")
		if err != nil {
			log.Printf("No library.json found, creating new empty library...")
			lib = library.New()
		}
		// Build grouped arcs
		arcs := map[int]*UnifiedArc{}

		// Loop over all episodes in metadata
		for crc, ep := range meta.Cache.EpisodesByCRC {

			// Ensure arc exists in output map
			if _, ok := arcs[ep.Arc]; !ok {
				arcTitle := meta.GetArcTitle(ep.Arc)
				arcMetadata, err := meta.GetArcByNumber(ep.Arc)
				if err != nil {
					log.Printf("Error fetching arc metadata for arc %d: %v", ep.Arc, err)
				}
				arcs[ep.Arc] = &UnifiedArc{
					Arc:               ep.Arc,
					Title:             arcTitle,
					AudioLanguages:    arcMetadata.AudioLanguages,
					SubtitleLanguages: arcMetadata.SubtitleLanguages,
					Resolution:        arcMetadata.Resolution,
					Status:            arcMetadata.Status,
					Episodes:          []UnifiedEpisode{},
				}
			}

			entry := UnifiedEpisode{
				Arc:            ep.Arc,
				Episode:        ep.Episode,
				Title:          ep.Title,
				Description:    ep.Description,
				Released:       ep.Released,
				Version:        ep.File.Version,
				CRC32:          crc,
				DownloadStatus: "missing",
				Monitored:      false,
			}

			// Merge library info if available
			if arc, ok := lib.Arcs[ep.Arc]; ok {
				key := fmt.Sprintf("%d", ep.Episode)
				if libEp, ok := arc.Episodes[key]; ok {
					entry.DownloadStatus = libEp.DownloadStatus
					entry.Monitored = libEp.Monitored
					entry.FilePath = libEp.FilePath
					// entry.Version = libEp.File.Version
				}
			}

			arcs[ep.Arc].Episodes = append(arcs[ep.Arc].Episodes, entry)
		}

		// Convert map → sorted slice
		result := []UnifiedArc{}
		for _, arc := range arcs {
			// Sort episodes by episode number
			sort.Slice(arc.Episodes, func(i, j int) bool {
				return arc.Episodes[i].Episode < arc.Episodes[j].Episode
			})
			result = append(result, *arc)
		}

		// Sort arcs by arc number
		sort.Slice(result, func(i, j int) bool {
			return result[i].Arc < result[j].Arc
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
