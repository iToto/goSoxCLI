package main

import (
	"github.com/iToto/goSoxCLI/soxCLI"
	"log"
)

func main() {
	path := "audio/test.wav"
	file, err := soxCLI.GetTrackInfo(path)

	if err != nil {
		log.Fatalf("Could not get File Info for %s", path)
	}

	log.Printf("File: %v", file)
}
