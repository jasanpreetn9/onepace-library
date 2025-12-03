package library

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(path string, lib *Library) error {
	tmp := path + ".tmp"

	b, err := json.MarshalIndent(lib, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	// Write to temporary file first
	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return fmt.Errorf("write tmp: %w", err)
	}

	// Replace existing file
	if err := os.Rename(tmp, path); err != nil {
		return fmt.Errorf("rename tmp: %w", err)
	}

	return nil
}
