// ffmpegCLI will interface with the sox cli comamnd
// NB: sox bin must be installed and in your $PATH
package soxCLI

import (
	"log"
	"os/exec"
	"strings"
)

type FileInfo struct {
	InputFile      string // song.wav
	Channels       string // 2
	SampleRate     string // 44100
	Precision      string // 16-bit
	Duration       string // 00:03:14.07
	Samples        string // 8558371
	Sectors        string // 14555.1 CDDA sectors
	FileSize       string // 35.0M
	BitRate        string // 1.44M
	SampleEncoding string // 16-bit Signed Integer PCM
}

func GetTrackInfo(path string) (FileInfo, error) {
	command := "soxi"
	out, err := exec.Command(command, path).Output()

	if err != nil {
		log.Fatalf("Error when running soxi: %v", err)
		return FileInfo{}, err
	}

	outString := string(out)
	outArray := strings.Split(outString, "\n")
	outSlice := outArray[1 : len(outArray)-2]
	insertIndex := 0
	var fileValues [10]string

	for _, element := range outSlice {
		array := strings.Split(element, ": ")

		if strings.TrimSpace(array[0]) == "Duration" {
			splitDuration := strings.Split(array[1], "=")

			for index, element := range splitDuration {
				array := strings.Split(element, " ")
				if index == 0 {
					fileValues[insertIndex] = array[0]
				} else {
					fileValues[insertIndex] = array[1]
				}

				insertIndex++
			}
		} else {
			fileValues[insertIndex] = array[1]
			insertIndex++
		}
	}

	file := FileInfo{
		InputFile:      fileValues[0],
		Channels:       fileValues[1],
		SampleRate:     fileValues[2],
		Precision:      fileValues[3],
		Duration:       fileValues[4],
		Samples:        fileValues[5],
		Sectors:        fileValues[6],
		FileSize:       fileValues[7],
		BitRate:        fileValues[8],
		SampleEncoding: fileValues[9],
	}

	return file, nil
}
