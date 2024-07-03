package src

func SearchRight(x, y, wordCount int, word, colums string, wordFound bool) bool {
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
		}
	}
	return wordFound
}
