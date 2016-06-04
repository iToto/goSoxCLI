# goSoxCLI

A go wrapper for the [SoX](http://sox.sourceforge.net) cli audio util.

***NB:*** *Must have sox installed and have the bin in your $PATH*

# Usage
```go
import (
	"github.com/iToto/audicut/soxCLI"
	"log"
)

func main() {
	path := "PATH_TO_FILE.wav"
	file, err := soxCLI.GetTrackInfo(path)

	if err != nil {
		log.Fatalf("Could not get File Info for %s", path)
	}

	log.Printf("File: %v", file)
}

```

## TODO
- [x] Get file information
- [ ] Input desired interval in seconds: [0, 60]
- [ ] Trim the audio file to desired interval.
			eg: [0, 60] would return the first minute of audio file
- [ ] Input multiple intervals: [[0, 60], [70, 120]]
- [ ] Create new audio file by attaching intervals
		eg: [[0, 60], [70, 120]] would return the first minute, skip 10 seconds, and attach the next 50 seconds.
