package ui

import (
	"fmt"

	"github.com/DylanMrr/seawar/domain"
)

func PrintField(board *domain.Board) {
	fmt.Println("  |a b c d e f g h i j")
	fmt.Println("______________________")
	for i := 0; i < 10; i++ {
		fmt.Print(i + 1)
		if i+1 < 10 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < 10; j++ {
			if (*board).Cells[j][i].State == 0 {
				fmt.Print(". ")
			} else if (*board).Cells[j][i].State == 1 {
				fmt.Print("O ")
			} else if (*board).Cells[j][i].State == 2 {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
