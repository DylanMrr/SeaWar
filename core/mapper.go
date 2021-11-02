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
		"i": 8,
		"j": 9,
	}
	if val, ok := dict[symbol]; ok {
		return val, nil
	}
	return -1, errors.New(fmt.Sprintf("Cannot find symbol %s in board", symbol))
}

func MapIndexToChar(i, j int) string {
	var dict = map[int]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
		5: "f",
		6: "g",
		7: "h",
		8: "i",
		9: "j",
	}
	return fmt.Sprint(dict[j], i+1)
}
