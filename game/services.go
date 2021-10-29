package game

import (
	"fmt"

	"github.com/DylanMrr/seawar/domain"
)

func CheckHit(board *domain.Board, i int, j int) bool {
	fmt.Println((*board).Cells[i][j].State == 1)
	return (*board).Cells[i][j].State == 1
}

func IsShipDestroyed(shootedCells *[]domain.Cell, board *domain.Board) bool {
	for _, ship := range (*board).Ships {
		shipArea := ship.GetShipArea()
		if shipArea.Contains((*shootedCells)[0].I, (*shootedCells)[0].J) {
			return len(*shootedCells) == ship.Length
		}
	}
	return false
}
