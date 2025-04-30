package logic

import (
	"path/filepath"
	"regexp"
)

// Parse album name from the folder name (e.g., "Pump (1989)" -> "Pump")
var albumNameRegex = regexp.MustCompile(`^(.*) \(\d{4}\)$`)

func ExtractAlbumName(folder string) string {
	base := filepath.Base(folder)
	match := albumNameRegex.FindStringSubmatch(base)
	if len(match) == 2 {
		return match[1]
	}
	return base
}
