package domain

import (
	"github.com/DylanMrr/seawar/core"
)

type Ship struct {
	Length      int
	FirstCell   Cell
	Orientation core.Orientation
}

func (ship *Ship) GetShipArea() Rect {
	var (
		xStart int
		xEnd   int
		yStart int
		yEnd   int
	)

	if (*ship).FirstCell.XIndex == 0 {
		xStart = 0
	} else {
		xStart = (*ship).FirstCell.XIndex - 1
	}

	if (*ship).FirstCell.YIndex == 0 {
		yStart = 0
	} else {
		yStart = (*ship).FirstCell.YIndex - 1
	}

	//h
	if (*ship).Orientation == 0 || (*ship).Orientation == 2 {
		if (*ship).FirstCell.XIndex+(*ship).Length >= 9 {
			xEnd = 9
		} else {
			xEnd = (*ship).FirstCell.XIndex + (*ship).Length
		}
		if (*ship).FirstCell.YIndex+(*ship).Length >= 9 {
			yEnd = 9
		} else {
			yEnd = (*ship).FirstCell.YIndex + (*ship).Length
		}
	} else { //v
		if (*ship).FirstCell.YIndex+(*ship).Length >= 9 {
			yEnd = 9
		} else {
			yEnd = (*ship).FirstCell.YIndex + (*ship).Length
		}
		if (*ship).FirstCell.XIndex+(*ship).Length >= 9 {
			xEnd = 9
		} else {
			xEnd = (*ship).FirstCell.XIndex + (*ship).Length
		}
	}
	return Rect{XStart: xStart, XEnd: xEnd, YStart: yStart, YEnd: yEnd}
}

func (ship *Ship) GetShip() Rect {
	var (
		xStart int
		xEnd   int
		yStart int
		yEnd   int
	)
	xStart = (*ship).FirstCell.XIndex
	yStart = (*ship).FirstCell.YIndex

	if (*ship).Orientation == 0 || (*ship).Orientation == 2 {
		xEnd = (*ship).FirstCell.XIndex + (*ship).Length - 1
		yEnd = yStart
	} else {
		xEnd = xStart
		yEnd = (*ship).FirstCell.YIndex + (*ship).Length - 1
	}

	return Rect{XStart: xStart, XEnd: xEnd, YStart: yStart, YEnd: yEnd}
}
