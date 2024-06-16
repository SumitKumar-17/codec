package main

import (
	"fmt"
	"log"
	"video-compression/internal/segmenter"
	"video-compression/internal/playlist"
)

func main() {
	// ask for the input file
	inputFile := "./input/Insight-GenX-AI.mp4"

	outputDir := "./output"
	segmentDuration := 10 // in seconds

	err := segmenter.SegmentVideo(inputFile, outputDir, segmentDuration)
	if err != nil {
		log.Fatalf("Error segmenting video: %v", err)
	}

	err = playlist.GenerateM3U8(outputDir, segmentDuration)
	if err != nil {
		log.Fatalf("Error generating playlist: %v", err)
	}

	fmt.Println("Video segmented and playlist generated successfully.")
}
