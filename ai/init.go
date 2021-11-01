package ai

import (
	"math/rand"
	"time"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func InitField(k *int) (*domain.Board, bool) {
	var board domain.Board = domain.Board{}

	for i := 9; i >= 0; i-- {
		var ship *domain.Ship
		ok := false
		for j := 0; j < 100; j++ {
			shipTemp, isOk := canAddShipToBoard(&board, core.ShipsTypes[i])
			if isOk {
				ship = shipTemp
				ok = true
				break
			}
			(*k)++
		}
		if !ok {
			return nil, false
		}
		board.AddShipToBoard(ship)
	}

	return &board, true
}

func canAddShipToBoard(board *domain.Board, shipSize int) (*domain.Ship, bool) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	j := rand.Intn(10)
	orientation := rand.Intn(2)
	firstCell := domain.Cell{I: i, J: j, State: core.Ship}
	ship := domain.Ship{Orientation: core.Orientation(orientation), Length: shipSize, FirstCell: firstCell}

	isOk := ship.ValidateShip()
	if !isOk {
		return nil, false
	}

	shipNearArea := ship.GetShip()

	for i := shipNearArea.IStart; i <= shipNearArea.IEnd; i++ {
		for j := shipNearArea.JStart; j <= shipNearArea.JEnd; j++ {
			if (*board).Cells[i][j].State != core.Empty {
				return nil, false
			}
		}
	}
	return &ship, true
}
