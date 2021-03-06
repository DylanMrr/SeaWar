package output

import (
	"fmt"

	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

func PrintBoard(board *domain.Board) {
	fmt.Println("  |a b c d e f g h i j")
	fmt.Println("______________________")
	for i := 0; i < 10; i++ {
		fmt.Print(i + 1)
		if i+1 < 10 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < 10; j++ {
			if (*board).Cells[i][j].State == core.Empty {
				fmt.Print(". ")
			} else if (*board).Cells[i][j].State == core.Ship {
				fmt.Print("O ")
			} else if (*board).Cells[i][j].State == core.NearShip {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintBoards(userBoard *domain.Board, userFightBoard *domain.Board) {
	fmt.Println()
	fmt.Println("Твоя доска                    Доска соперника")
	fmt.Println("  |a b c d e f g h i j          |a b c d e f g h i j")
	fmt.Println("______________________         ______________________")
	for i := 0; i < 10; i++ {
		fmt.Print(i + 1)
		if i+1 < 10 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < 10; j++ {
			if (*userBoard).Cells[i][j].State == core.Ship {
				fmt.Print("O ")
			} else if (*userBoard).Cells[i][j].State == core.Checked {
				fmt.Print("* ")
			} else if (*userBoard).Cells[i][j].State == core.Hitted {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}

		fmt.Print("       ")
		fmt.Print(i + 1)
		if i+1 < 10 {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for j := 0; j < 10; j++ {
			if (*userFightBoard).Cells[i][j].State == core.Ship {
				fmt.Print("O ")
			} else if (*userFightBoard).Cells[i][j].State == core.Checked {
				fmt.Print("* ")
			} else if (*userFightBoard).Cells[i][j].State == core.Hitted {
				fmt.Print("x ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
