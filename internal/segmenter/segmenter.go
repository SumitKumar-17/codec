package segmenter

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func SegmentVideo(inputFile, outputDir string, segmentDuration int, quality int) error {
	// Ensure output directory exists
	err := exec.Command("mkdir", "-p", outputDir).Run()
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	scaleFilter := fmt.Sprintf("scale=1280:%d", quality)
	fmt.Printf("scaleFilter: %s\n", scaleFilter)
	// Use ffmpeg to segment the video
	segmentPattern := filepath.Join(outputDir, "segment_%03d.ts")

	cmd := exec.Command("ffmpeg", "-i", inputFile , "-c", "copy", "-map", "0",
		"-f", "segment", "-segment_time", fmt.Sprintf("%d", segmentDuration), segmentPattern)

	// cmd := exec.Command("ffmpeg", "-i", inputFile, "-vf", fmt.Sprintf(`"%s"`, scaleFilter), "-c", "copy", "-map", "0",
	// 	"-f", "segment", "-segment_time", fmt.Sprintf("%d", segmentDuration), segmentPattern)
	fmt.Println(cmd.String())
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("ffmpeg command failed: %w", err)
	}

	return nil
}
