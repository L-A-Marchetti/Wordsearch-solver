package src

func SearchDown(x, y, wordCount int, word, colums string, wordFound bool) bool {
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
			//CopiedGrid.DisplayGrid()
			Save = append(Save, word)
			wordFound = true
		}
	}
	return wordFound
}
