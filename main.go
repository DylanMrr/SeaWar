package main

import (
	"fmt"

	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui"
	"github.com/DylanMrr/seawar/ui/input"
)

func main() {
	ui.PrintStartText()

	var userBoard domain.Board = domain.Board{IsPlayerBoard: true}

	var ship *domain.Ship
	for ship, isOk := canAddShipToBoard(&userBoard); !isOk{
		
	}

}

func canAddShipToBoard(board *domain.Board) (*domain.Ship, bool) {

	var ship *domain.Ship
	var err error
	if ship, err = input.InputShip(1); err != nil {
		fmt.Println(err)
		return nil, false
	}

	shipNearArea := ship.GetShipArea()
	//shipArea := ship.GetShip()

	for i := shipNearArea.XStart; i <= shipNearArea.XEnd; i++ {
		for j := shipNearArea.YStart; i <= shipNearArea.YEnd; j++ {
			if (*board).Cells[j][i].State != 0 {
				fmt.Println("Занятые клетки")
				return nil, false
			}
		}
	}
	return &ship, true
}
