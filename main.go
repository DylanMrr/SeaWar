package main

import (
	"github.com/DylanMrr/seawar/ai"
	"github.com/DylanMrr/seawar/domain"
	"github.com/DylanMrr/seawar/ui"
)

func main() {
	ui.PrintStartText()

	//board := game.InitField()
	//ui.PrintField(board)

	var aiBoard *domain.Board

	for true {
		boardTemp, ok := ai.InitField()
		if ok {
			aiBoard = boardTemp
			break
		}
	}
	ui.PrintField(aiBoard)
}
