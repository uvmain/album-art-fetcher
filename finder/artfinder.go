package finder

import (
	"album-art-fetcher/config"
	"fmt"
	"net/url"
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
