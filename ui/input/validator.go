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

	i, err := strconv.Atoi(string((*cell)[1:]))
	if err != nil {
		return -1, -1, errors.New("Некорректный числовой индекс")
	}

	if strings.Contains(Symbols, string((*cell)[0])) && i >= 1 && i <= 10 {
		j, err := core.MapCharToIndex(string((*cell)[0]))
		if err != nil {
			return -1, -1, err
		}
		return i - 1, j, nil
	}
	return -1, -1, errors.New("Некорректная ячейка")
}
