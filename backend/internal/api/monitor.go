package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"onepace-library/internal/library"
)

type MonitorRequest struct {
	Arc       int  `json:"arc"`
	Episode   int  `json:"episode"`
	Monitored bool `json:"monitored"`
}

func HandleMonitorEpisode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lib, err := library.Load("data/library.json")
		if err != nil {
			log.Printf("No library.json found, creating new empty library...")
			lib = library.New()
		}
		var req MonitorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		arc := lib.GetOrCreateArc(req.Arc, "")
		key := fmt.Sprintf("%d", req.Episode)

		// --- If unmonitoring, delete the entry instead of creating empty metadata ---
		if !req.Monitored {
			delete(arc.Episodes, key)
			library.Save("data/library.json", lib)

			json.NewEncoder(w).Encode(map[string]string{
				"status": "unmonitored",
			})
			return
		}

		// --- Otherwise: ensure the episode exists ---
		ep, ok := arc.Episodes[key]
		if !ok {
			ep = library.Episode{
				EpisodeNumber: req.Episode,
			}
		}

		ep.Monitored = true
		arc.Episodes[key] = ep

		library.Save("data/library.json", lib)

		json.NewEncoder(w).Encode(map[string]string{
			"status": "monitored",
		})
	}
}
