package main

import (
	"album-art-fetcher/config"
	"album-art-fetcher/downloader"
	"album-art-fetcher/getter"
	"album-art-fetcher/logic"
	"album-art-fetcher/scanner"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	config.LoadEnv()

	albums, err := scanner.ScanMusicDir(config.MusicDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error scanning: %v\n", err)
		os.Exit(1)
	}

	countAlbums := len(albums)
	countFound := 0
	countDownloaded := 0
	countFailed := 0

	failedDownloads := make(map[string]error)

	for _, album := range albums {
		albumName := logic.ExtractAlbumName(album.AlbumDir)
		fmt.Printf("Directory: %s\n", album.AlbumDir)
		fmt.Printf("Artist: %s\n", album.Artist)
		fmt.Printf("Album: %s\n", albumName)
		fileTarget := filepath.Join(album.AlbumDir, "folder.jpg")

		if album.MissingArt {
			downloadUrl, err := getter.FetchAlbumArtURL(album.Artist, albumName)
			if err != nil {
				fmt.Printf("Could not get %spx URL for %s - %v\n", config.AlbumArtPixelsString, album.AlbumDir, err)
				countFailed++
				failedDownloads[album.AlbumDir] = err
			} else {
				fmt.Printf("Downloading artwork from: %s\n", downloadUrl)
				err = downloader.DownloadAndSaveAsJPG(downloadUrl, fileTarget)
				if err != nil {
					fmt.Printf("Failed to download artwork: %s\n", err)
					countFailed++
					failedDownloads[album.AlbumDir] = err
				}
				countDownloaded++
			}
		} else {
			fmt.Println("Artwork found")
			countFound++
		}

		fmt.Println()
	}

	if countFailed > 0 {
		fmt.Println("\nFailed downloads details:")
		for directory, err := range failedDownloads {
			fmt.Printf(" - Directory: %s\n   Error: %v\n", directory, err)
		}
	}
	fmt.Println()

	fmt.Println("\n==================== Summary ====================")
	fmt.Printf("Total albums:          %d\n", countAlbums)
	fmt.Printf("Albums with artwork:   %d\n", countFound)
	fmt.Printf("Artworks downloaded:   %d\n", countDownloaded)
	fmt.Printf("Failed downloads:      %d\n", countFailed)
	fmt.Println("================================================")
}
