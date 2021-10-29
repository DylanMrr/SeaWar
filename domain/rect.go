package domain

type Rect struct {
	IStart int
	IEnd   int
	JStart int
	JEnd   int
}

func (rect *Rect) Contains(i int, j int) bool {
	return i >= (*rect).IStart && i <= (*rect).IEnd && j >= (*rect).JStart && j <= (*rect).JEnd
}
