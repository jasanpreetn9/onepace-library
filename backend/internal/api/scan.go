package api

import (
	"encoding/json"
	"log"
	"net/http"
	"onepace-library/internal/config"
	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
	"onepace-library/internal/scanner"
)

func HandleScanLibrary(meta *metadata.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lib, err := library.Load("data/library.json")
		if err != nil {
			log.Printf("No library.json found, creating new empty library...")
			lib = library.New()
		}
		cfg, _ := config.Load("data/config.yml")

		if err := scanner.ScanLibrary(cfg.LibraryPath, lib, meta); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		library.Save(cfg.LibraryJSONPath, lib)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "ok",
			"message": "library scan complete",
		})
	}
}

func HandleScanDownloads(meta *metadata.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lib, err := library.Load("data/library.json")
		if err != nil {
			log.Printf("No library.json found, creating new empty library...")
			lib = library.New()
		}
		cfg, _ := config.Load("data/config.yml")

		if err := scanner.ScanDownloads(cfg.DownloadPath, cfg.LibraryPath, lib, meta); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		library.Save(cfg.LibraryJSONPath, lib)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "ok",
			"message": "download scan complete",
		})
	}
}
