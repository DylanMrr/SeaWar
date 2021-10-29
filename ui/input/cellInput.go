package input

import (
	"fmt"

	"github.com/DylanMrr/seawar/domain"
)

func InputCell() *domain.Cell {
	isValid := false
	for !isValid {
		fmt.Println("Введите ячейку для стрельбы в формате b4")
		var cell string
		fmt.Scan(&cell)
		fmt.Println()

		i, j, err := ValidateCell(&cell)
		if err != nil {
			fmt.Println(err)
			continue
		}
		firstCell := domain.Cell{I: i, J: j}

		return &firstCell
	}
	return nil
}
