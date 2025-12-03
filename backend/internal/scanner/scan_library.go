package scanner

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
	"onepace-library/internal/nfo"
)

func ScanLibrary(root string, lib *library.Library, meta *metadata.Client) error {

	log.Printf("Starting library scan: %s", root)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		filename := info.Name()

		if !(strings.HasSuffix(strings.ToLower(filename), ".mkv") ||
			strings.HasSuffix(strings.ToLower(filename), ".mp4")) {
			return nil
		}

		parsed, err := ParseOnePaceFilename(filename)
		if err != nil {
			return nil
		}

		epMeta, err := meta.GetEpisodeByCRC32(parsed.CRC32)
		if err != nil {
			log.Printf("Metadata missing for CRC %s (%s)", parsed.CRC32, filename)
			return nil
		}

		arcTitle := meta.GetArcTitle(epMeta.Arc)

		entry := addOrUpdateEpisode(lib, path, parsed, epMeta, arcTitle)

		nfoPath := nfo.NFOPathForVideo(path)
		nfo.GenerateEpisodeNFO(entry, epMeta, arcTitle, nfoPath)

		return nil
	})

	log.Println("Library scan complete.")
	return err
}
