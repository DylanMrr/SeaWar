package main

import (
	"github.com/DylanMrr/seawar/game"
	"github.com/DylanMrr/seawar/ui"
)

func main() {
	ui.PrintStartText()

	board := game.InitField()
	ui.PrintField(board)
}
