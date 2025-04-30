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

	countAlbums := len(albums)
	countFound := 0
	countDownloaded := 0
	countFailed := 0

	for _, album := range albums {
		albumName := ExtractAlbumName(album.AlbumDir)
		fmt.Printf("Directory: %s\n", album.AlbumDir)
		fmt.Printf("Artist: %s\n", album.Artist)
		fmt.Printf("Album: %s\n", albumName)

		if album.MissingArt {
			downloadUrl, err := FetchAlbumArtURL(album.Artist, albumName)
			if err != nil {
				fmt.Printf("Could not get %spx URL for %s - %v\n", config.AlbumArtPixelsString, album.AlbumDir, err)
				countFailed++
			} else {
				fmt.Printf("Artwork can be downloaded from: %s\n", downloadUrl)
				countDownloaded++
			}
		} else {
			fmt.Println("Artwork found")
			countFound++
		}

		fmt.Println()
	}

	fmt.Println("Summary:")
	fmt.Printf("Total albums: %d\n", countAlbums)
	fmt.Printf("Skipped: %d\n", countFound)
	fmt.Printf("Artworks downloaded: %d\n", countDownloaded)
	fmt.Printf("Failed downloads: %d\n", countFailed)
}
