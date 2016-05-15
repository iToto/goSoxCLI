package soxCLI

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoxCLI_testGetLengthOfTrack(t *testing.T) {
	testFile := "/Users/iToto/Dev/GitHub/audicut/audio/test.wav"
	file, err := GetTrackInfo(testFile)

	assert.Nil(t, err, "Caught error: %s", err)
	assert.ObjectsAreEqual(FileInfo{}, file)
	assert.IsType(t, FileInfo{}, file, "FileInfo not returned: %v", file)
}
