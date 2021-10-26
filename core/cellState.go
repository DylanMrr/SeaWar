package core

type CellState int

const (
	Empty    CellState = iota //0
	Ship                      //1
	NearShip                  //2
	Checked                   //3
)

func (c CellState) Index() int {
	return int(c)
}
