package domain

import "github.com/DylanMrr/seawar/core"

type Ship struct {
	Length      int
	FirstCell   Cell
	Orientation core.Orientation
}
