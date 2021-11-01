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
	//var aiFightBoard domain.Board
	aiFightBoard := domain.New()
	//ui.PrintField(&aiFightBoard)

	//output.PrintBoard(userBoard)
	//output.PrintBoard(&userFightBoard)
	output.PrintBoards(userBoard, &userFightBoard)

	userMove := false

	userPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: userBoard, FightBoard: &userFightBoard}
	aiPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: aiBoard, FightBoard: aiFightBoard}
	fmt.Println(userBoard.Cells)
	n := 1
	ai.BuildMoves()

	bot := ai.Bot{}

	for userPlayer.ShipCells > 0 && aiPlayer.ShipCells > 0 {
		//todo уничтожение корабля соперника
		if userMove {
			fmt.Println("Ваш ход!")
			chosenCell := input.InputCell()
			for !validateCellState(&userFightBoard, chosenCell) {
				chosenCell = input.InputCell()
			}

			if CheckHit(aiBoard, (*chosenCell).I, (*chosenCell).J) {
				aiPlayer.ShipCells--
				aiPlayer.Board.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Hitted

				userFightBoard.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Hitted

			} else {
				fmt.Println("Ход соперника!")
				userMove = false
				aiPlayer.Board.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Checked
				userFightBoard.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Checked
			}
		} else {
			i, j := bot.MakeMove(aiFightBoard)
			fmt.Println("i ", i, "j", j)
			if CheckHit(userBoard, i, j) {
				userPlayer.ShipCells--
				bot.MarkCellHitted(aiFightBoard, i, j)
				bot.Shot(aiFightBoard, i, j)
				userBoard.Cells[i][j].State = core.Hitted

				if IsShipDestroyed(bot.Cells, userBoard) {
					fmt.Println("Ваш корабль уничтожен")
					bot.ShipDestroyedCallback(aiFightBoard)
				}

			} else {
				userBoard.Cells[i][j].State = core.Checked
				bot.MarkCellChecked(aiFightBoard, i, j)
				bot.Miss()
				userMove = true
			}
		}

		fmt.Println("Шаг ", n)
		output.PrintBoards(userBoard, &userFightBoard)
		//output.PrintBoards(userBoard, aiFightBoard)
		fmt.Println()
		n++
	}
}

func validateCellState(fightBoard *domain.Board, chosenCell *domain.Cell) bool {
	if (*fightBoard).Cells[(*chosenCell).I][(*chosenCell).J].State != core.Empty {
		fmt.Println("Ячейка уже была выбрана")
		return false
	}
	return true
}
