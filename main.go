package main

import (
	"src"
	"time"
)

func main() {
	src.WordSearchGrid.GetWordSearch()
	src.CopiedGrid.CopyGrid()
	defer src.TimeTrack(time.Now())
	src.Solve(0)
}
