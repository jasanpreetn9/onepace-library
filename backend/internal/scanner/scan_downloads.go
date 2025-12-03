package scanner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"onepace-library/internal/library"
	"onepace-library/internal/metadata"
	"onepace-library/internal/nfo"
)

func ScanDownloads(downloadRoot, libraryRoot string, lib *library.Library, meta *metadata.Client) error {

	log.Printf("Scanning downloads: %s\n", downloadRoot)

	entries, err := os.ReadDir(downloadRoot)
	if err != nil {
		return fmt.Errorf("read downloads: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if !(strings.HasSuffix(strings.ToLower(name), ".mkv") ||
			strings.HasSuffix(strings.ToLower(name), ".mp4")) {
			continue
		}

		parsed, err := ParseOnePaceFilename(name)
		if err != nil {
			log.Printf("Skipping non-OnePace file: %s", name)
			continue
		}

		// Metadata lookup
		epMeta, err := meta.GetEpisodeByCRC32(parsed.CRC32)
		if err != nil {
			log.Printf("No metadata for CRC %s (file: %s)", parsed.CRC32, name)
			continue
		}

		arcTitle := meta.GetArcTitle(epMeta.Arc)

		// Build arc folder in library
		arcFolder := filepath.Join(
			libraryRoot,
			fmt.Sprintf("%02d - %s", epMeta.Arc, sanitizeFilename(arcTitle)),
		)
		os.MkdirAll(arcFolder, 0755)

		destFilename := fmt.Sprintf(
			"S%02dE%02d - %s [%s].%s",
			epMeta.Arc,
			epMeta.Episode,
			sanitizeFilename(epMeta.Title),
			parsed.CRC32,
			parsed.Extension,
		)

		src := filepath.Join(downloadRoot, name)
		dst := filepath.Join(arcFolder, destFilename)

		if err := os.Rename(src, dst); err != nil {
			log.Printf("Failed to move file: %v", err)
			continue
		}

		// Merge into library
		entry := addOrUpdateEpisode(lib, dst, parsed, epMeta, arcTitle)

		// Write NFO
		nfoPath := nfo.NFOPathForVideo(dst)
		nfo.GenerateEpisodeNFO(entry, epMeta, arcTitle, nfoPath)

		log.Printf("Imported: %s → %s", name, dst)
	}

	log.Println("Downloads scan complete.")
	return nil
}
