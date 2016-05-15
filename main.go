package main

import (
	"github.com/iToto/audicut/soxCLI"
	"log"
)

func main() {
	path := "/Users/iToto/Dev/golang/src/github.com/iToto/audicut/audio/test.wav"
	file, err := soxCLI.GetTrackInfo(path)

	if err != nil {
		log.Fatalf("Could not get File Info for %s", path)
	}

	log.Printf("File: %v", file)
}
