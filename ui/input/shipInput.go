package input

import (
	"errors"
	"fmt"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func InputShip(size int) (*domain.Ship, error) {
	if size == 1 {
		fmt.Printf("Введите корабль из %d ячейки в формате b4..\n", size)

		var ship string
		fmt.Scan(&ship)

		i, j, err := ValidateCell(&ship)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		firstCell := domain.Cell{I: i, J: j, State: 1}
		var orientation core.Orientation = core.Unit
		shipObject := domain.Ship{Length: 1, FirstCell: firstCell, Orientation: orientation}

		isOk := shipObject.ValidateShip()
		if !isOk {
			return nil, errors.New("Корабль не помещается")
		}
		return &shipObject, nil
	} else {
		fmt.Printf("Введите верхнюю левую ячейку для корабля из %d ячейкеек в формате b4\n", size)
		var ship string
		fmt.Scan(&ship)
		fmt.Println()

		i, j, err := ValidateCell(&ship)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		firstCell := domain.Cell{I: i, J: j, State: 1}

		fmt.Print("Введите ориентацию корабля. 0 - горизонтально, 1 - вертикально \n")
		var orientation core.Orientation
		fmt.Scan(&orientation)

		shipObject := domain.Ship{Length: size, FirstCell: firstCell, Orientation: orientation}

		isOk := shipObject.ValidateShip()
		if !isOk {
			return nil, errors.New("Корабль не помещается")
		}
		return &shipObject, nil
	}
}
