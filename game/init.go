package game

import (
	"fmt"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui/input"
	"github.com/DylanMrr/seawar/ui/output"
)

func InitField() *domain.Board {
	var userBoard domain.Board = domain.Board{IsPlayerBoard: true}

	for i := 0; i < 10; i++ {
		var ship *domain.Ship
		for true {
			shipTemp, isOk := canAddShipToBoard(&userBoard, core.ShipsTypes[i])
			if isOk {
				ship = shipTemp
				break
			}
			fmt.Println("Попробуйте еще раз")
		}
		userBoard.AddShipToBoard(ship)
		output.PrintBoard(&userBoard)
	}
	return &userBoard
}

func canAddShipToBoard(board *domain.Board, shipSize int) (*domain.Ship, bool) {
	var err error
	var ship *domain.Ship
	if ship, err = input.InputShip(shipSize); err != nil {
		fmt.Println(err)
		return nil, false
	}

	shipNearArea := ship.GetShip()

	for i := shipNearArea.XStart; i <= shipNearArea.XEnd; i++ {
		for j := shipNearArea.YStart; j <= shipNearArea.YEnd; j++ {
			if (*board).Cells[i][j].State != 0 {
				fmt.Println("Занятые клетки")
				return nil, false
			}
		}
	}
	return ship, true
}
