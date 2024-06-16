package playlist

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)

func GenerateM3U8(outputDir string, segmentDuration int) error {
	files, err := ioutil.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("failed to read output directory: %w", err)
	}

	playlistFile := filepath.Join(outputDir, "index.m3u8")
	f, err := os.Create(playlistFile)
	if err != nil {
		return fmt.Errorf("failed to create playlist file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString("#EXTM3U\n#EXT-X-VERSION:3\n")
	if err != nil {
		return fmt.Errorf("failed to write to playlist file: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ts" {
			_, err = f.WriteString(fmt.Sprintf("#EXTINF:%d,\n%s\n", segmentDuration, file.Name()))
			if err != nil {
				return fmt.Errorf("failed to write to playlist file: %w", err)
			}
		}
	}

	_, err = f.WriteString("#EXT-X-ENDLIST\n")
	if err != nil {
		return fmt.Errorf("failed to write to playlist file: %w", err)
	}

	return nil
}
