package downloader

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadAndSaveAsJPG(imageURL, outputPath string) error {
	resp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status downloading image: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	var img image.Image

	switch contentType {
	case "image/jpeg", "image/jpg":
		img, err = jpeg.Decode(resp.Body)
	case "image/png":
		img, err = png.Decode(resp.Body)
	default:
		img, _, err = image.Decode(resp.Body)
	}

	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	if filepath.Ext(outputPath) != ".jpg" {
		outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + ".jpg"
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	opts := jpeg.Options{Quality: 90}
	if err := jpeg.Encode(outFile, img, &opts); err != nil {
		return fmt.Errorf("failed to encode image to jpg: %w", err)
	}

	return nil
}
