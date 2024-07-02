package main

import (
	"src"
	"time"
)

func main() {
	src.WordSearchGrid.GetWordSearch()
	src.WordSearchGrid.DisplayGrid()
	defer src.TimeTrack(time.Now())
	src.Solve(0)
}
