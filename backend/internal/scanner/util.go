package scanner

import "strings"

// normalizePath makes paths consistent for JSON output.
func normalizePath(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}
func sanitizeFilename(s string) string {
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, c := range invalid {
		s = strings.ReplaceAll(s, c, "_")
	}
	return s
}
