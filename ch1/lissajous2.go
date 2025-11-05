// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
