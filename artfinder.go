package main

import (
	"album-art-fetcher/config"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
)

func BuildLastfmApiUrl(artist, album string) string {
	escapedArtist := url.QueryEscape(artist)
	escapedAlbum := url.QueryEscape(album)

	return fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=%s&artist=%s&album=%s&format=json",
		config.LastfmAPIKey,
		escapedArtist,
		escapedAlbum,
	)
}

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
