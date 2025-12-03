package api

import (
	"encoding/json"
	"net/http"

	"onepace-library/internal/config"
	"onepace-library/internal/metadata"
	"onepace-library/internal/qbittorrent"
)

type AddDownloadRequest struct {
	CRC32 string `json:"crc32"`
}

func HandleAddToQbit(meta *metadata.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req AddDownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Lookup metadata
		ep, err := meta.GetEpisodeByCRC32(req.CRC32)
		if err != nil {
			http.Error(w, "episode not found", http.StatusNotFound)
			return
		}

		// Episodes.json contains only a single "file" entry
		if ep.File.URL == "" {
			http.Error(w, "no download URL available", http.StatusInternalServerError)
			return
		}

		downloadURL := ep.File.URL

		// Config
		cfg, _ := config.Load("data/config.yml")
		qb := qbittorrent.NewClient(
			cfg.QBittorrent.Host,
			cfg.QBittorrent.Username,
			cfg.QBittorrent.Password,
		)

		// Send to qBit
		if err := qb.AddTorrent(downloadURL); err != nil {
			http.Error(w, "qBit add failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "ok",
			"message": "download added to qBittorrent",
			"url":     downloadURL,
		})
	}
}
