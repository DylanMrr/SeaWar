package domain

import "github.com/DylanMrr/seawar/core"

type Board struct {
	Cells [10][10]Cell
	Ships []*Ship
}

func New() *Board {
	board := Board{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board.Cells[i][j] = Cell{I: i, J: j, State: core.Empty}
		}
	}
	return &board
}

func (board *Board) AddShipToBoard(ship *Ship) {
	shipNearArea := ship.GetShipArea()
	shipArea := ship.GetShip()

	for i := shipNearArea.IStart; i <= shipNearArea.IEnd; i++ {
		for j := shipNearArea.JStart; j <= shipNearArea.JEnd; j++ {
			if shipArea.Contains(i, j) {
				(*board).Cells[i][j].State = core.Ship
			} else {
				(*board).Cells[i][j].State = core.NearShip
			}
		}
	}
	(*board).Ships = append((*board).Ships, ship)
}
