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
					if x+1 < WordSearchGrid.Width && WordSearchGrid.Grid[y][x+1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y][x+1]
						count := 2
						for x+count < WordSearchGrid.Width && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y][x+count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y][x+count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y][x+i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y][x+i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search left.
					if x-1 >= 0 && WordSearchGrid.Grid[y][x-1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y][x-1]
						count := 2
						for x-count >= 0 && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y][x-count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y][x-count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y][x-i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y][x-i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search down.
					if y+1 < WordSearchGrid.Height && WordSearchGrid.Grid[y+1][x] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y+1][x]
						count := 2
						for y+count < WordSearchGrid.Height && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y+count][x] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y+count][x]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y+i][x] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y+i][x] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search up.
					if y-1 >= 0 && WordSearchGrid.Grid[y-1][x] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y-1][x]
						count := 2
						for y-count >= 0 && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y-count][x] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y-count][x]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y-i][x] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y-i][x] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search upper-right.
					if y-1 >= 0 && x+1 < WordSearchGrid.Width && WordSearchGrid.Grid[y-1][x+1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y-1][x+1]
						count := 2
						for y-count >= 0 && x+count < WordSearchGrid.Width && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y-count][x+count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y-count][x+count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y-i][x+i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y-i][x+i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search upper-left.
					if y-1 >= 0 && x-1 >= 0 && WordSearchGrid.Grid[y-1][x-1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y-1][x-1]
						count := 2
						for y-count >= 0 && x-count >= 0 && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y-count][x-count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y-count][x-count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y-i][x-i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y-i][x-i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search lower-right.
					if y+1 < WordSearchGrid.Height && x+1 < WordSearchGrid.Width && WordSearchGrid.Grid[y+1][x+1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y+1][x+1]
						count := 2
						for y+count < WordSearchGrid.Height && x+count < WordSearchGrid.Width && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y+count][x+count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y+count][x+count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y+i][x+i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y+i][x+i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
						word = colums
					}
					// Search lower-left.
					if y+1 < WordSearchGrid.Height && x-1 >= 0 && WordSearchGrid.Grid[y+1][x-1] == string(WordSearchGrid.ListWords()[wordCount][1]) {
						word += WordSearchGrid.Grid[y+1][x-1]
						count := 2
						for y+count < WordSearchGrid.Height && x-count >= 0 && count < len(WordSearchGrid.ListWords()[wordCount]) &&
							WordSearchGrid.Grid[y+count][x-count] == string(WordSearchGrid.ListWords()[wordCount][count]) {
							word += WordSearchGrid.Grid[y+count][x-count]
							count++
						}
						if word == WordSearchGrid.ListWords()[wordCount] {
							for i := 0; i < count; i++ {
								CopiedGrid.DispGrid[y+i][x-i] = "\033[" + Light + "" + Color + "m" + CopiedGrid.DispGrid[y+i][x-i] + "\033[0m"
							}
							CopiedGrid.DisplayGrid()
							Save = append(Save, word)
							wordFound = true
							break
						}
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
