package domain

type Board struct {
	IsPlayerBoard bool
	Cells         [10][10]Cell
}
