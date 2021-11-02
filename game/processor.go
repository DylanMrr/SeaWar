package game

import (
	"fmt"
	"time"

	"github.com/DylanMrr/seawar/ai"
	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui/input"
	"github.com/DylanMrr/seawar/ui/output"
)

func StartGame() {

	var aiBoard *domain.Board

	channel := make(chan *domain.Board)
	go func(channel chan *domain.Board) {
		for true {
			boardTemp, ok := ai.InitField()
			if ok {
				channel <- boardTemp
				break
			}
		}
	}(channel)

	userBoard := InitField()

	userFightBoard := domain.New()

	aiFightBoard := domain.New()

	output.PrintBoards(userBoard, userFightBoard)

	userMove := true

	n := 1
	ai.BuildMoves()

	bot := ai.Bot{}

	aiBoard = <-channel
	userPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: userBoard, FightBoard: userFightBoard}
	aiPlayer := domain.Player{ShipCells: core.ShipsCellsCount, Board: aiBoard, FightBoard: aiFightBoard}

	for userPlayer.ShipCells > 0 && aiPlayer.ShipCells > 0 {
		output.PrintDelimiter(fmt.Sprint("Шаг ", n))
		if userMove {
			fmt.Println("Твой ход!")

			chosenCell := input.InputCell()
			for !validateCellState(userFightBoard, chosenCell) {
				chosenCell = input.InputCell()
			}

			if CheckHit(aiBoard, (*chosenCell).I, (*chosenCell).J) {
				aiPlayer.ShipCells--
				aiPlayer.Board.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Hitted

				userFightBoard.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Hitted

				userPlayer.ShootedCells = append(userPlayer.ShootedCells, &userFightBoard.Cells[(*chosenCell).I][(*chosenCell).J])
				output.PrintBoards(userBoard, userFightBoard)
				if IsShipDestroyed(userPlayer.ShootedCells, aiBoard) {
					fmt.Println("*** Убил!")
					userPlayer.ShootedCells = nil
				} else {
					fmt.Println("** Ранил!")
				}

			} else {
				userMove = false
				aiPlayer.Board.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Checked
				userFightBoard.Cells[(*chosenCell).I][(*chosenCell).J].State = core.Checked
				output.PrintBoards(userBoard, userFightBoard)
				fmt.Println("Мимо!")
				fmt.Println("Ход соперника!")
			}
		} else {
			i, j := bot.MakeMove(aiFightBoard)

			time.Sleep(2 * time.Second)

			fmt.Println("Соперник сходил - ", core.MapIndexToChar(i, j))

			if CheckHit(userBoard, i, j) {
				userPlayer.ShipCells--
				bot.MarkCellHitted(aiFightBoard, i, j)
				bot.Shot(aiFightBoard, i, j)
				userBoard.Cells[i][j].State = core.Hitted
				output.PrintBoards(userBoard, userFightBoard)
				if IsShipDestroyed(bot.Cells, userBoard) {
					fmt.Println("*** Твой корабль уничтожен")
					bot.ShipDestroyedCallback(aiFightBoard)
				} else {
					fmt.Println("** Соперник ранил!")
				}

			} else {
				userBoard.Cells[i][j].State = core.Checked
				bot.MarkCellChecked(aiFightBoard, i, j)
				bot.Miss()
				output.PrintBoards(userBoard, userFightBoard)
				fmt.Println("Соперник промахнулся!")
				userMove = true
			}
		}

		//output.PrintBoards(userBoard, &userFightBoard)
		fmt.Println()
		n++
	}
	output.PrintDelimiter("")
	output.PrintDelimiter("Игра завершена!")
	if userPlayer.ShipCells > 0 {
		fmt.Println("****** Ты победил! Поздравляю! ******")
	} else {
		fmt.Println("****** К сожалению, ты проиграл( Не отчаивайся и пробуй заново! ******")
		fmt.Println("Доска соперника")
		output.PrintBoard(aiBoard)
	}
}

func validateCellState(fightBoard *domain.Board, chosenCell *domain.Cell) bool {
	if (*fightBoard).Cells[(*chosenCell).I][(*chosenCell).J].State != core.Empty {
		fmt.Println("Ячейка уже была выбрана")
		return false
	}
	return true
}
