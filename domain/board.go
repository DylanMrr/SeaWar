package domain

type Board struct {
	IsPlayerBoard bool
	Cells         [10][10]Cell
	Ships         []*Ship
}

func (board *Board) AddShipToBoard(ship *Ship) {
	shipNearArea := ship.GetShipArea()
	shipArea := ship.GetShip()

	for i := shipNearArea.IStart; i <= shipNearArea.IEnd; i++ {
		for j := shipNearArea.JStart; j <= shipNearArea.JEnd; j++ {
			if shipArea.Contains(i, j) {
				(*board).Cells[i][j].State = 1
			} else {
				(*board).Cells[i][j].State = 2
			}
		}
	}
	(*board).Ships = append((*board).Ships, ship)
}
