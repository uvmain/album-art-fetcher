package getter

import (
	"album-art-fetcher/config"
	"album-art-fetcher/finder"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type AlbumInfoResponse struct {
	Album struct {
		Image []struct {
			Size string `json:"size"`
			URL  string `json:"#text"`
		} `json:"image"`
	} `json:"album"`
}

var imageHashRegex = regexp.MustCompile(`/([a-f0-9]{32})\.(jpg|png)$`)

func FetchAlbumArtURL(artist, album string) (string, error) {

	url := finder.BuildLastfmApiUrl(artist, album)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data AlbumInfoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	// Look for any valid image URL with a hash
	for _, img := range data.Album.Image {
		if img.URL == "" {
			continue
		}
		if match := imageHashRegex.FindStringSubmatch(img.URL); len(match) == 3 {
			hash := match[1]
			ext := match[2]
			customURL := fmt.Sprintf("https://lastfm.freetls.fastly.net/i/u/%sx%s/%s.%s", config.AlbumArtPixelsString, config.AlbumArtPixelsString, hash, ext)
			return customURL, nil
		}
	}

	return "", fmt.Errorf("no usable image URL found")
}
