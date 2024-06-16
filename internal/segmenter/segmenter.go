package segmenter

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func SegmentVideo(inputFile, outputDir string, segmentDuration int) error {
	// Ensure output directory exists
	err := exec.Command("mkdir", "-p", outputDir).Run()
	if err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Use ffmpeg to segment the video
	segmentPattern := filepath.Join(outputDir, "segment_%03d.ts")
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c", "copy", "-map", "0", 
		"-f", "segment", "-segment_time", fmt.Sprintf("%d", segmentDuration), segmentPattern)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("ffmpeg command failed: %w", err)
	}

	return nil
}
