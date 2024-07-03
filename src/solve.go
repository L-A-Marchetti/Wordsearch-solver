package src

import (
	"fmt"
)

func Solve(wordCount int) {
	//time.Sleep(100 * time.Millisecond)
	Clear()
	if found() {
		CopiedGrid.DisplayGrid()
		fmt.Println("Solved:\n")
		for _, w := range Save {
			fmt.Println(w)
		}
		return
	} else {
		wordFound := false
		for y, line := range WordSearchGrid.Grid {
			for x, colums := range line {
				if colums == string(WordSearchGrid.ListWords()[wordCount][0]) {
					word := colums
					// Search right.
					wordFound = SearchRight(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search left.
					wordFound = SearchLeft(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search down.
					wordFound = SearchDown(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search up.
					wordFound = SearchUp(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search upper-right.
					wordFound = SearchUpperRight(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search upper-left.
					wordFound = SearchUpperLeft(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search lower-right.
					wordFound = SearchLowerRight(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
					// Search lower-left.
					wordFound = SearchLowerLeft(x, y, wordCount, word, colums, wordFound)
					if wordFound {
						break
					}
				}
			}
			if wordFound {
				if Color == "7" {
					Color = "1"
					Light = "3"
				} else {
					Color = string(Color[0] + 1)
				}
				break
			}
		}
		Solve(wordCount + 1)
	}
}
