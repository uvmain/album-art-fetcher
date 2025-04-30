package scanner

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Album struct {
	Artist     string
	AlbumDir   string // Full path to the album folder
	MissingArt bool
}

func normalizePath(path string) string {
	// Replace problematic characters (e.g., '+') with underscores
	return strings.Map(func(r rune) rune {
		if r == '+' || !unicode.IsPrint(r) {
			return '_'
		}
		return r
	}, path)
}

func ScanMusicDir(root string) ([]Album, error) {
	var albums []Album

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Normalize the path to handle problematic characters
		path = normalizePath(path)

		// Check if this is a folder and directly under root/Artist/Album
		if d.IsDir() {
			rel, _ := filepath.Rel(root, path)
			parts := strings.Split(rel, string(os.PathSeparator))

			if len(parts) == 2 {
				// This is an album folder
				folderJpg := filepath.Join(path, "folder.jpg")
				_, err := os.Stat(folderJpg)
				missing := os.IsNotExist(err)

				albums = append(albums, Album{
					Artist:     parts[0],
					AlbumDir:   path,
					MissingArt: missing,
				})
			}
		}

		return nil
	})

	return albums, err
}
