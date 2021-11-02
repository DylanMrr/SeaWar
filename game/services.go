package game

import (
	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func CheckHit(board *domain.Board, i int, j int) bool {
	return (*board).Cells[i][j].State == core.Ship
}

func IsShipDestroyed(shootedCells []*domain.Cell, board *domain.Board) bool {
	for _, ship := range (*board).Ships {
		shipArea := ship.GetShip()
		if shipArea.Contains(shootedCells[0].I, shootedCells[0].J) {
			return len(shootedCells) == ship.Length
		}
	}
	return false
}
