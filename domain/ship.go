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
		iStart int
		iEnd   int
		jStart int
		jEnd   int
	)

	if (*ship).FirstCell.I == 0 {
		iStart = 0
	} else {
		iStart = (*ship).FirstCell.I - 1
	}

	if (*ship).FirstCell.J == 0 {
		jStart = 0
	} else {
		jStart = (*ship).FirstCell.J - 1
	}

	//h
	if (*ship).Orientation == core.Horizontal || (*ship).Orientation == core.Unit {
		if (*ship).FirstCell.J+(*ship).Length >= 9 {
			jEnd = 9
		} else {
			jEnd = (*ship).FirstCell.J + (*ship).Length
		}
		if (*ship).FirstCell.I+1 >= 9 {
			iEnd = 9
		} else {
			iEnd = (*ship).FirstCell.I + 1
		}
	} else { //v
		if (*ship).FirstCell.I+(*ship).Length >= 9 {
			iEnd = 9
		} else {
			iEnd = (*ship).FirstCell.I + (*ship).Length
		}
		if (*ship).FirstCell.J+1 >= 9 {
			jEnd = 9
		} else {
			jEnd = (*ship).FirstCell.J + 1
		}
	}
	return Rect{JStart: jStart, JEnd: jEnd, IStart: iStart, IEnd: iEnd}
}

func (ship *Ship) GetShip() Rect {
	var (
		iStart int
		iEnd   int
		jStart int
		jEnd   int
	)
	jStart = (*ship).FirstCell.J
	iStart = (*ship).FirstCell.I

	if (*ship).Orientation == core.Horizontal || (*ship).Orientation == core.Unit {
		jEnd = (*ship).FirstCell.J + (*ship).Length - 1
		iEnd = iStart
	} else {
		jEnd = jStart
		iEnd = (*ship).FirstCell.I + (*ship).Length - 1
	}

	return Rect{JStart: jStart, JEnd: jEnd, IStart: iStart, IEnd: iEnd}
}

func (ship *Ship) ValidateShip() bool {
	if (*ship).Orientation == core.Horizontal || (*ship).Orientation == core.Unit {
		if (*ship).FirstCell.J+(*ship).Length > 10 {
			return false
		}
	} else { //v
		if (*ship).FirstCell.I+(*ship).Length > 10 {
			return false
		}
	}
	return true
}
