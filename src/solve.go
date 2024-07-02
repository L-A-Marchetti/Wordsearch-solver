package src

import "fmt"

func Solve(word int) {
	if found() {
		return
	} else {
		for y, line := range WordSearchGrid.Grid {
			for x, colums := range line {
				if colums == string(WordSearchGrid.ListWords()[0][0]) {
					fmt.Println(colums)
					if x+1 < WordSearchGrid.Width && WordSearchGrid.Grid[y][x+1] == string(WordSearchGrid.ListWords()[0][1]) {
						fmt.Println(WordSearchGrid.Grid[y][x+1] + " right")
					} else if x-1 >= 0 && WordSearchGrid.Grid[y][x-1] == string(WordSearchGrid.ListWords()[0][1]) {
						fmt.Println(WordSearchGrid.Grid[y][x-1] + " left")
					} else if y+1 < WordSearchGrid.Height && WordSearchGrid.Grid[y+1][x] == string(WordSearchGrid.ListWords()[0][1]) {
						fmt.Println(WordSearchGrid.Grid[y+1][x] + " down")
					} else if y-1 >= 0 && WordSearchGrid.Grid[y-1][x] == string(WordSearchGrid.ListWords()[0][1]) {
						fmt.Println(WordSearchGrid.Grid[y-1][x] + " up")
					}
				}
			}
		}
		return
	}
}
