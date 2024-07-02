package src

import "fmt"

func Solve(word int) {
	if found() {
		return
	} else {
		for y, line := range WordSearchGrid.Grid {
			for x, colums := range line {
				if colums == string(WordSearchGrid.ListWords()[0][0]) {
					word := colums
					fmt.Println(word)
					// Search right.
					if x+1 < WordSearchGrid.Width && WordSearchGrid.Grid[y][x+1] == string(WordSearchGrid.ListWords()[0][1]) {
						word += WordSearchGrid.Grid[y][x+1]
						count := 2
						fmt.Println(word + " right")
						for x+count < WordSearchGrid.Width && count < len(WordSearchGrid.ListWords()[0]) &&
							WordSearchGrid.Grid[y][x+count] == string(WordSearchGrid.ListWords()[0][count]) {
							word += WordSearchGrid.Grid[y][x+count]
							count++
						}
						if word == WordSearchGrid.ListWords()[0] {
							fmt.Println("found " + word)
							WordSearchGrid.Grid[y][x] = "\033[37m" + WordSearchGrid.Grid[y][x]
							WordSearchGrid.Grid[y][x+count-1] += "\033[0m"
							WordSearchGrid.DisplayAll()
							return
						}
						word = colums
						// Search left.
					}
					if x-1 >= 0 && WordSearchGrid.Grid[y][x-1] == string(WordSearchGrid.ListWords()[0][1]) {
						word += WordSearchGrid.Grid[y][x-1]
						count := 2
						fmt.Println(word + " left")
						for x-count >= 0 && count < len(WordSearchGrid.ListWords()[0]) &&
							WordSearchGrid.Grid[y][x-count] == string(WordSearchGrid.ListWords()[0][count]) {
							word += WordSearchGrid.Grid[y][x-count]
							count++
						}
						if word == WordSearchGrid.ListWords()[0] {
							fmt.Println("found " + word)
							WordSearchGrid.Grid[y][x-count+1] = "\033[37m" + WordSearchGrid.Grid[y][x-count+1]
							WordSearchGrid.Grid[y][x] += "\033[0m"
							WordSearchGrid.DisplayAll()
							return
						}
						word = colums
						// Search down.
					}
					if y+1 < WordSearchGrid.Height && WordSearchGrid.Grid[y+1][x] == string(WordSearchGrid.ListWords()[0][1]) {
						word += WordSearchGrid.Grid[y+1][x]
						count := 2
						fmt.Println(word + " down")
						for y+count < WordSearchGrid.Height && count < len(WordSearchGrid.ListWords()[0]) &&
							WordSearchGrid.Grid[y+count][x] == string(WordSearchGrid.ListWords()[0][count]) {
							word += WordSearchGrid.Grid[y+count][x]
							count++
						}
						if word == WordSearchGrid.ListWords()[0] {
							fmt.Println("found " + word)
							for i := 0; i < count; i++ {
								WordSearchGrid.Grid[y+i][x] = "\033[37m" + WordSearchGrid.Grid[y+i][x] + "\033[0m"
							}
							WordSearchGrid.DisplayAll()
							return
						}
						word = colums
						// Search up.
					}
					if y-1 >= 0 && WordSearchGrid.Grid[y-1][x] == string(WordSearchGrid.ListWords()[0][1]) {
						word += WordSearchGrid.Grid[y-1][x]
						count := 2
						fmt.Println(word + " up")
						for y-count >= 0 && count < len(WordSearchGrid.ListWords()[0]) &&
							WordSearchGrid.Grid[y-count][x] == string(WordSearchGrid.ListWords()[0][count]) {
							word += WordSearchGrid.Grid[y-count][x]
							count++
						}
						if word == WordSearchGrid.ListWords()[0] {
							fmt.Println("found " + word)
							for i := 0; i < count; i++ {
								WordSearchGrid.Grid[y-i][x] = "\033[37m" + WordSearchGrid.Grid[y-i][x] + "\033[0m"
							}
							WordSearchGrid.DisplayAll()
							return
						}
					}
				}
			}
		}
		return
	}
}
