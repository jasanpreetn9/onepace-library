package api

import (
	"github.com/go-chi/chi/v5"

	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
)

func RegisterRoutes(r chi.Router, lib *library.Library, meta *metadata.Client) {
	r.Route("/api", func(api chi.Router) {

		api.Get("/library", HandleGetLibrary())

		api.Post("/scan/library", HandleScanLibrary(meta))
		api.Post("/scan/downloads", HandleScanDownloads(meta))

		api.Get("/episodes/all", HandleGetAllEpisodes(meta))
		api.Get("/episodes/{crc}", HandleGetEpisode(meta))
		api.Post("/episodes/monitor", HandleMonitorEpisode())

		api.Post("/download/add", HandleAddToQbit(meta))

	})
}
