package ai

import (
	"math/rand"
	"time"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func InitField() (*domain.Board, bool) {
	var board domain.Board = domain.Board{IsPlayerBoard: false}

	for i := 0; i < 10; i++ {
		var ship *domain.Ship
		ok := false
		for j := 0; j < 50; j++ {
			shipTemp, isOk := canAddShipToBoard(&board, core.ShipsTypes[i])
			if isOk {
				ship = shipTemp
				ok = true
				break
			}
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
	x := rand.Intn(10)
	y := rand.Intn(10)
	orientation := rand.Intn(2)
	firstCell := domain.Cell{XIndex: x, YIndex: y, State: 1}
	ship := domain.Ship{Orientation: core.Orientation(orientation), Length: shipSize, FirstCell: firstCell}

	isOk := ship.ValidateShip()
	if !isOk {
		return nil, false
	}

	shipNearArea := ship.GetShip()

	for i := shipNearArea.XStart; i <= shipNearArea.XEnd; i++ {
		for j := shipNearArea.YStart; j <= shipNearArea.YEnd; j++ {
			if (*board).Cells[i][j].State != 0 {
				return nil, false
			}
		}
	}
	return &ship, true
}
