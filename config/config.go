package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var MusicDirectory string
var AlbumArtPixels uint
var AlbumArtPixelsString string
var LastfmAPIKey string

func LoadEnv() {

	godotenv.Load(".env")

	musicDirectory := os.Getenv("MUSIC_DIRECTORY")
	if musicDirectory == "" {
		musicDirectory = "./"
	}

	MusicDirectory, _ = filepath.Abs(musicDirectory)

	lastfmAPIKey := os.Getenv("LASTFM_API_KEY")
	if lastfmAPIKey == "" {
		fmt.Println("LASTFM_API_KEY not set in .env file")
		os.Exit(1)
	}
	LastfmAPIKey = lastfmAPIKey

	u, _ := strconv.ParseUint(os.Getenv("ALBUM_ART_PIXELS"), 10, 64)
	if u > 0 {
		AlbumArtPixels = uint(u)
	} else {
		AlbumArtPixels = 600
	}

	AlbumArtPixelsString = strconv.Itoa(int(AlbumArtPixels))

}
