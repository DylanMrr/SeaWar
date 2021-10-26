package domain

import "github.com/DylanMrr/seawar/core"

type Cell struct {
	XIndex int
	YIndex int
	State  core.CellState
}
