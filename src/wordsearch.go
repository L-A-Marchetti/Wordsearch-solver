package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WordSearch struct {
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	WordsCount int        `json:"wordsCount"`
	Grid       [][]string `json:"grid"`
	Words      []Word     `json:"words"`
}

type Word struct {
	Word     string   `json:"word"`
	Position Position `json:"position"`
}

type Position struct {
	Start []int `json:"start"`
	End   []int `json:"end"`
}

func (w *WordSearch) GetWordSearch() {
	api, _ := http.Get("https://shadify.dev/api/wordsearch/generator")
	defer api.Body.Close()
	data, _ := ioutil.ReadAll(api.Body)
	json.Unmarshal(data, &w)
}

func (w *WordSearch) StartX(i int) int {
	return w.Words[i].Position.Start[0]
}

func (w *WordSearch) StartY(i int) int {
	return w.Words[i].Position.Start[1]
}

func (w *WordSearch) EndX(i int) int {
	return w.Words[i].Position.End[0]
}

func (w *WordSearch) EndY(i int) int {
	return w.Words[i].Position.End[1]
}

func (w *WordSearch) Display() {
	fmt.Printf("width: %d\nheight: %d\nwords count: %d\n", w.Width, w.Height, w.WordsCount)
	for _, line := range w.Grid {
		fmt.Println(line)
	}
	for i, word := range w.Words {
		fmt.Printf("word : %s, position -> start: [%d][%d], end: [%d][%d]\n", word.Word, w.StartX(i), w.StartY(i), w.EndX(i), w.EndY(i))
	}
}
