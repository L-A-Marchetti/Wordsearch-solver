package src

import (
	"reflect"
)

func found() bool {
	return reflect.DeepEqual(WordSearchGrid.ListWords(), Save)
}
