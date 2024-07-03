package src

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	CopiedGrid.DisplayGrid()
	fmt.Println("Solved:\n")
	for _, w := range Save {
		fmt.Println(w)
	}
	fmt.Printf("\nElapsed Time: %s\n", elapsed)
}
