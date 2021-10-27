package game

import (
	"fmt"

	"github.com/DylanMrr/seawar/ai"
	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui/input"
	"github.com/DylanMrr/seawar/ui/output"
)

func StartGame() {
	userBoard := InitField()
	//ui.PrintField(userBoard)

	var userFightBoard domain.Board
	//ui.PrintField(&userFightBoard)

	var aiBoard *domain.Board
	for true {
		boardTemp, ok := ai.InitField()
		if ok {
			aiBoard = boardTemp
			break
		}
	}
	//ui.PrintField(aiBoard)
	var aiFightBoard domain.Board
	//ui.PrintField(&aiFightBoard)

	//output.PrintBoard(userBoard)
	//output.PrintBoard(&userFightBoard)
	output.PrintBoards(userBoard, &userFightBoard)

	userMove := true

	userPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: userBoard, FightBoard: &userFightBoard}
	aiPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: aiBoard, FightBoard: &aiFightBoard}

	for userPlayer.ShipCells > 0 || aiPlayer.ShipCells > 0 {
		if userMove {
			fmt.Println("Ваш ход!")
			chosenCell := input.InputCell()
			for !validateCellState(&userFightBoard, chosenCell) {
				chosenCell = input.InputCell()
			}

			if aiPlayer.Board.Cells[(*chosenCell).YIndex][(*chosenCell).XIndex].State == 1 {
				aiPlayer.ShipCells--
				aiPlayer.Board.Cells[(*chosenCell).YIndex][(*chosenCell).XIndex].State = 4
			} else {
				fmt.Println("Ваш ход!")
				fmt.Println("Ход соперника!")
				userMove = false
				aiPlayer.Board.Cells[(*chosenCell).YIndex][(*chosenCell).XIndex].State = 3
			}
		} else {

		}
	}
}

func validateCellState(fightBoard *domain.Board, chosenCell *domain.Cell) bool {
	if (*fightBoard).Cells[(*chosenCell).YIndex][(*chosenCell).XIndex].State != 0 {
		fmt.Println("Ячейка уже была выбрана")
		return false
	}
	return true
}
