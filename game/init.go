package game

import (
	"fmt"

	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui"
	"github.com/DylanMrr/seawar/ui/input"
)

func InitField() *domain.Board {
	var userBoard domain.Board = domain.Board{IsPlayerBoard: true}

	for i := 0; i < 4; i++ {
		var ship *domain.Ship
		for true {
			shipTemp, isOk := canAddShipToBoard(&userBoard, 1)
			if isOk {
				ship = shipTemp
				break
			}
			fmt.Println("Попробуйте еще раз")
		}
		addShipToBoard(ship, &userBoard)
		ui.PrintField(&userBoard)
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

	shipNearArea := ship.GetShipArea()

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

func addShipToBoard(ship *domain.Ship, board *domain.Board) {
	shipNearArea := ship.GetShipArea()
	shipArea := ship.GetShip()

	for i := shipNearArea.XStart; i <= shipNearArea.XEnd; i++ {
		for j := shipNearArea.YStart; j <= shipNearArea.YEnd; j++ {
			if shipArea.Contains(i, j) {
				(*board).Cells[i][j].State = 1
			} else {
				(*board).Cells[i][j].State = 2
			}
		}
	}
}
