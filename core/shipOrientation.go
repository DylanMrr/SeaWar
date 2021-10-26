package core

type Orientation int

const (
	Horizontal Orientation = iota //0
	Vertical                      //1
	Unit                          //2
)

func (o Orientation) Index() int {
	return int(o)
}
