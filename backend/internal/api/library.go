package api

import (
	"encoding/json"
	"log"
	"net/http"
	"onepace-library/internal/library"
)

func HandleGetLibrary() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lib, err := library.Load("data/library.json")
		if err != nil {
			log.Printf("No library.json found, creating new empty library...")
			lib = library.New()
		}
		w.Header().Set("Content-Type", "application/json")

		b, err := json.MarshalIndent(lib, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(b)
	}
}
