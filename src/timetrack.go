package src

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("\nElapsed Time: %s\n", elapsed)
}
