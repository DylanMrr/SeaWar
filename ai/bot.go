package ai

import (
	"github.com/DylanMrr/seawar/core"
	"github.com/DylanMrr/seawar/domain"
)

type Bot struct {
	Cells          []*domain.Cell
	CellsAroundHit []*domain.Cell
}

func (bot *Bot) MakeMove(aiFigthBoard *domain.Board) (int, int) {
	if len((*bot).Cells) == 0 {
		return GetMove()
	} else {
		for i := len((*bot).CellsAroundHit) - 1; i >= 0; i-- {
			if (*bot).CellsAroundHit[i].I < 0 || (*bot).CellsAroundHit[i].I > 9 || (*bot).CellsAroundHit[i].J < 0 || (*bot).CellsAroundHit[i].J > 9 || (*bot).CellsAroundHit[i].State != core.Empty {
				continue
			}
			return (*bot).CellsAroundHit[i].I, (*bot).CellsAroundHit[i].J
		}
		return -1, -1
	}
}

func (bot *Bot) Miss() {
	if len((*bot).Cells) > 0 {
		(*bot).Cells = (*bot).Cells[0:1]
	}
}

func (bot *Bot) Shot(aiFightBoard *domain.Board, i int, j int) {
	(*bot).Cells = append((*bot).Cells, &(*aiFightBoard).Cells[i][j])
	(*bot).buildCellsAroundHit(aiFightBoard, len((*bot).Cells)-1)
}

func (bot *Bot) ShipDestroyedCallback(aiFightBoard *domain.Board) {
	for _, cell := range (*bot).CellsAroundHit {
		cell.State = core.Checked
	}

	(*bot).Cells = nil
	(*bot).CellsAroundHit = nil
}

func (bot *Bot) buildCellsAroundHit(aiFightBoard *domain.Board, i int) {
	if (*bot).Cells[i].J+1 <= 9 && (*aiFightBoard).Cells[(*bot).Cells[i].I][(*bot).Cells[i].J+1].State == core.Empty {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, &(*aiFightBoard).Cells[(*bot).Cells[i].I][(*bot).Cells[i].J+1])
	}
	if (*bot).Cells[i].I+1 <= 9 && (*aiFightBoard).Cells[(*bot).Cells[i].I+1][(*bot).Cells[i].J].State == core.Empty {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, &(*aiFightBoard).Cells[(*bot).Cells[i].I+1][(*bot).Cells[i].J])
	}
	if (*bot).Cells[i].J-1 >= 0 && (*aiFightBoard).Cells[(*bot).Cells[i].I][(*bot).Cells[i].J-1].State == core.Empty {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, &(*aiFightBoard).Cells[(*bot).Cells[i].I][(*bot).Cells[i].J-1])
	}
	if (*bot).Cells[i].I-1 >= 0 && (*aiFightBoard).Cells[(*bot).Cells[i].I-1][(*bot).Cells[i].J].State == core.Empty {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, &(*aiFightBoard).Cells[(*bot).Cells[i].I-1][(*bot).Cells[i].J])
	}
}

func (bot *Bot) MarkCellChecked(aiFightBoard *domain.Board, i int, j int) {
	(*aiFightBoard).Cells[i][j].State = core.Checked
}

func (bot *Bot) MarkCellHitted(aiFightBoard *domain.Board, i int, j int) {

	(*aiFightBoard).Cells[i][j].State = core.Hitted
	var markI []int
	var markJ []int
	if j == 0 {
		if i == 0 {
			markI = append(markI, i+1)
			markJ = append(markJ, j+1)
		} else if i == 9 {
			markI = append(markI, i-1)
			markJ = append(markJ, j+1)
		} else {
			markJ = append(markJ, j+1)
			markI = append(markI, i-1)

			markJ = append(markJ, j+1)
			markI = append(markI, i+1)
		}
	} else if j == 9 {
		if i == 0 {
			markJ = append(markJ, j-1)
			markI = append(markI, i+1)
		} else if i == 9 {
			markJ = append(markJ, j-1)
			markI = append(markI, i-1)
		} else {
			markJ = append(markJ, j-1)
			markI = append(markI, i-1)

			markJ = append(markJ, j-1)
			markI = append(markI, i+1)
		}
	} else if i == 0 {
		markJ = append(markJ, j-1)
		markI = append(markI, i+1)

		markJ = append(markJ, j+1)
		markI = append(markI, i+1)
	} else if i == 9 {
		markJ = append(markJ, j-1)
		markI = append(markI, i-1)

		markJ = append(markJ, j+1)
		markI = append(markI, i-1)
	} else {
		markJ = append(markJ, j-1)
		markI = append(markI, i-1)

		markJ = append(markJ, j+1)
		markI = append(markI, i-1)

		markJ = append(markJ, j-1)
		markI = append(markI, i+1)

		markJ = append(markJ, j+1)
		markI = append(markI, i+1)
	}

	for k := 0; k < len(markI); k++ {
		(*aiFightBoard).Cells[markI[k]][markJ[k]].State = core.Checked
	}
}
