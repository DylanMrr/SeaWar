package domain

type Board struct {
	IsPlayerBoard bool
	Cells         [10][10]Cell
}

func (board *Board) AddShipToBoard(ship *Ship) {
	shipNearArea := ship.GetShipArea()
	shipArea := ship.GetShip()

	for i := shipNearArea.XStart; i <= shipNearArea.XEnd; i++ {
		for j := shipNearArea.YStart; j <= shipNearArea.YEnd; j++ {
			if shipArea.Contains(i, j) {
				(*board).Cells[i][j].State = 1
			} else {
				(*board).Cells[i][j].State = 2
			}
		}
	}
}
