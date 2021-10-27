package input

import (
	"errors"
	"strconv"
	"strings"

	"github.com/DylanMrr/seawar/core"
)

func ValidateCell(cell *string) (int, int, error) {
	if len(*cell) != 2 && len(*cell) != 3 {
		return -1, -1, errors.New("Некорректный формат ввода ячейки")
	}

	index, err := strconv.Atoi(string((*cell)[1:]))
	if err != nil {
		return -1, -1, errors.New("Некорректный числовой индекс")
	}

	if strings.Contains(Symbols, string((*cell)[0])) && index >= 1 && index <= 10 {
		xIndex, err := core.MapCharToIndex(string((*cell)[0]))
		if err != nil {
			return -1, -1, err
		}
		return xIndex, index - 1, nil
	}
	return -1, -1, errors.New("Некорректная ячейка")
}
