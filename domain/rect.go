package domain

type Rect struct {
	XStart int
	XEnd   int
	YStart int
	YEnd   int
}

func (rect *Rect) Contains(x int, y int) bool {
	return x >= (*rect).XStart && x <= (*rect).XEnd && y >= (*rect).YStart && y <= (*rect).YEnd
}
