package input

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func InputShip(size int) (*domain.Ship, error) {
	if size == 1 {
		fmt.Printf("Введите корабль из %d ячейки в формате b4..\n", size)

		var ship string
		fmt.Scan(&ship)

		x, y, err := validateCell(&ship)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		firstCell := domain.Cell{XIndex: x, YIndex: y, State: 1}
		var orientation core.Orientation = core.Unit
		shipObject := domain.Ship{Length: 1, FirstCell: firstCell, Orientation: orientation}

		isOk := validateShip(&shipObject)
		if !isOk {
			return nil, errors.New("Корабль не помещается")
		}
		return &shipObject, nil
	} else {
		fmt.Printf("Введите верхнюю левую ячейку для корабля из %d ячейкеек в формате b4\n", size)
		var ship string
		fmt.Scan(&ship)

		x, y, err := validateCell(&ship)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		firstCell := domain.Cell{XIndex: x, YIndex: y, State: 1}

		fmt.Print("Введите ориентацию корабля. 0 - горизонтально, 1 - вертикально \n", size)
		var orientation core.Orientation
		fmt.Scan(&orientation)

		shipObject := domain.Ship{Length: size, FirstCell: firstCell, Orientation: orientation}

		isOk := validateShip(&shipObject)
		if !isOk {
			return nil, errors.New("Корабль не помещается")
		}
		return &shipObject, nil
	}
}

func validateCell(ship *string) (int, int, error) {
	if len(*ship) != 2 && len(*ship) != 3 {
		return -1, -1, errors.New("Некорректный формат ввода корабля")
	}

	index, err := strconv.Atoi(string((*ship)[1:]))
	if err != nil {
		return -1, -1, errors.New("Некорректный числовой индекс")
	}

	if strings.Contains(Symbols, string((*ship)[0])) && index >= 1 && index <= 10 {
		xIndex, err := core.MapCharToIndex(string((*ship)[0]))
		if err != nil {
			return -1, -1, err
		}
		return xIndex, index - 1, nil
	}
	return -1, -1, errors.New("Некорректный корабль")
}

func validateShip(ship *domain.Ship) bool {
	if (*ship).Orientation == 0 || (*ship).Orientation == 2 {
		if (*ship).FirstCell.XIndex+(*ship).Length > 10 {
			return false
		}
	} else { //v
		if (*ship).FirstCell.YIndex+(*ship).Length > 10 {
			return false
		}
	}
	return true
}
