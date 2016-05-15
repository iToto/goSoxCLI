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
