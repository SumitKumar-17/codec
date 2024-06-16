package main

import (
	"fmt"
	"log"

	"github.com/SumitKumar-17/codec/internal/playlist"
	"github.com/SumitKumar-17/codec/internal/segmenter"
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
