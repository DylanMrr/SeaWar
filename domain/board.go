package domain

type Board struct {
	isPlayerBoard bool
	cells         [10][10]Cell
}

func (board *Board) CreateBoard() {
	
}
