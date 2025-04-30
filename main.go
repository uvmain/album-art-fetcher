package main

import (
	"album-art-fetcher/config"
	"fmt"
	"os"
)

func main() {
	config.LoadEnv()

	albums, err := scanMusicDir(config.MusicDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error scanning: %v\n", err)
		os.Exit(1)
	}

	for _, album := range albums {
		albumName := ExtractAlbumName(album.AlbumDir)
		fmt.Printf("Directory: %s\n", album.AlbumDir)
		fmt.Printf("Album: %s\n", albumName)

		if album.MissingArt {
			downloadUrl, err := FetchAlbumArtURL(album.Artist, albumName)
			if err != nil {
				fmt.Printf("Could not get 600px URL for %s - %v\n", album.AlbumDir, err)
			} else {
				fmt.Printf("Artwork can be downloaded from: %s\n", downloadUrl)
			}
		} else {
			fmt.Printf("Artwork found: %s\n", album.AlbumDir)
		}

		fmt.Println()
	}
}
