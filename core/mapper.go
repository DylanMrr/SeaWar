package core

import (
	"errors"
	"fmt"
)

func MapCharToIndex(symbol string) (int, error) {
	var dict = map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
	}
	if val, ok := dict[symbol]; ok {
		return val, nil
	}
	return -1, errors.New(fmt.Sprintf("Cannot find symbol %s in board", symbol))
}
