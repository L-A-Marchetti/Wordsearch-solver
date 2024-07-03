package src

func SearchLowerLeft(x, y, wordCount int, word, colums string, wordFound bool) bool {
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
			//CopiedGrid.DisplayGrid()
			Save = append(Save, word)
			wordFound = true
		}
	}
	return wordFound
}
