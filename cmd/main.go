package main

import (
	"fmt"
	"log"

	"github.com/SumitKumar-17/codec/internal/playlist"
	"github.com/SumitKumar-17/codec/internal/segmenter"
)

func main() {
	// ask for the input file
	var inputFile string
	
	fmt.Print("Enter the input file path: ")
	_, err := fmt.Scanln(&inputFile)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
    
	var quality int
	// ask for the quality of the video
	fmt.Print("Enter the quality of the video: ")
	_,err=fmt.Scanln(&quality)
	if err != nil {
		log.Fatalf("Error reading quality: %v", err)
	}

	outputDir := "./output"
	segmentDuration := 10 // in seconds

	err = segmenter.SegmentVideo(inputFile, outputDir, segmentDuration,quality)
	if err != nil {
		log.Fatalf("Error segmenting video: %v", err)
	}

	err = playlist.GenerateM3U8(outputDir, segmentDuration)
	if err != nil {
		log.Fatalf("Error generating playlist: %v", err)
	}

	fmt.Println("Video segmented and playlist generated successfully.")
}
