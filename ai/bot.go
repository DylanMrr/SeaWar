package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DylanMrr/seawar/domain"
)

type Bot struct {
	FirstHittedShipCell *domain.Cell

	Cells          []domain.Cell
	CellsAroundHit []domain.Cell

	CurrentDirection int //0 - up, 1 - right, 2 - down, 3 - left
	Direction        int //-1 - undef, 0 - up, 1 - right, 2 - down, 3 - left
}

func (bot *Bot) MakeMove(aiFigthBoard *domain.Board) (int, int) {

	//if (*bot).FirstHittedShipCell == nil {
	if len((*bot).Cells) == 0 { //первый выстрел
		rand.Seed(time.Now().UnixNano())
		for k := 0; k < 100; k++ {
			i := rand.Intn(10)
			j := rand.Intn(10)

			if (*aiFigthBoard).Cells[i][j].State == 0 {
				return i, j
			}
		}
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (*aiFigthBoard).Cells[i][j].State == 0 {
					return i, j
				}
			}
		}
		return -1, -1
	} else { //} if len((*bot).Cells) == 1 { //попал в 1 палубу
		//if (*bot).kx == 0 && (*bot).ky == 0 {
		for i := 0; i < 4; i++ {
			if (*bot).CellsAroundHit[i].I < 0 || (*bot).CellsAroundHit[i].I > 9 || (*bot).CellsAroundHit[i].J < 0 || (*bot).CellsAroundHit[i].J > 9 || (*bot).CellsAroundHit[i].State != 0 {
				continue
			}
			return (*bot).CellsAroundHit[i].I, (*bot).CellsAroundHit[i].J
		}
		return -1, -1
		//}
	}
}

func (bot *Bot) Miss() {
	//if (*bot).Cells[len((*bot).Cells)-1].State != 4 {
	(*bot).Cells = nil
	(*bot).CellsAroundHit = nil
	//}
}

func (bot *Bot) Shot(aiFightBoard *domain.Board, i int, j int) {
	(*bot).Cells = append((*bot).Cells, domain.Cell{I: i, J: j})
	(*bot).buildCellsAroundHit(aiFightBoard, len((*bot).Cells)-1)
	/*if len((*bot).Cells) == 2 {
		if math.Abs(float64((*bot).Cells[0].XIndex)-float64((*bot).Cells[1].XIndex)) == 1 {
			(*bot).kx = 1
		} else {
			(*bot).kx = 0
		}
		if math.Abs(float64((*bot).Cells[0].YIndex)-float64((*bot).Cells[1].YIndex)) == 1 {
			(*bot).ky = 1
		} else {
			(*bot).ky = 0
		}
	}*/
}

func (bot *Bot) ShipDestroyedCallback(aiFightBoard *domain.Board) {
	for _, cell := range (*bot).CellsAroundHit {
		(*aiFightBoard).Cells[cell.I][cell.J].State = 3
	}

	(*bot).Cells = nil
	(*bot).CellsAroundHit = nil
}

func (bot *Bot) buildCellsAroundHit(aiFightBoard *domain.Board, i int) {
	(*bot).CellsAroundHit = make([]domain.Cell, 0, 4)
	fmt.Println("Cells: ", (*bot).Cells)
	//(*bot).CellsAroundHit = nil
	if (*bot).Cells[i].J+1 <= 9 && (*aiFightBoard).Cells[(*bot).Cells[i].J+1][(*bot).Cells[i].I].State == 0 {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, domain.Cell{I: (*bot).Cells[i].I, J: (*bot).Cells[i].J + 1})
		fmt.Println((*bot).Cells[i].I, (*bot).Cells[i].J+1)
	}
	if (*bot).Cells[i].I+1 <= 9 && (*aiFightBoard).Cells[(*bot).Cells[i].I+1][(*bot).Cells[i].J].State == 0 {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, domain.Cell{I: (*bot).Cells[i].I + 1, J: (*bot).Cells[i].J})
		fmt.Println((*bot).Cells[i].I+1, (*bot).Cells[i].J)
	}
	if (*bot).Cells[i].J-1 >= 0 && (*aiFightBoard).Cells[(*bot).Cells[i].I][(*bot).Cells[i].J-1].State == 0 {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, domain.Cell{I: (*bot).Cells[i].I, J: (*bot).Cells[i].J - 1})
		fmt.Println((*bot).Cells[i].I, (*bot).Cells[i].J-1)
	}
	if (*bot).Cells[i].I-1 >= 0 && (*aiFightBoard).Cells[(*bot).Cells[i].I-1][(*bot).Cells[i].J].State == 0 {
		(*bot).CellsAroundHit = append((*bot).CellsAroundHit, domain.Cell{I: (*bot).Cells[i].I - 1, J: (*bot).Cells[i].J})
		fmt.Println((*bot).Cells[i].I-1, (*bot).Cells[i].J)
	}
}

func (bot *Bot) MarkCellChecked(aiFightBoard *domain.Board, i int, j int) {
	(*aiFightBoard).Cells[i][j].State = 3
	if len((*bot).Cells) > 0 {
		(*bot).findCellAroungHit(i, j).State = 3
	}
	//(*bot).Cells[len((*bot).Cells)-1].State = 3
}

func (bot *Bot) findCellAroungHit(i int, j int) *domain.Cell {
	for _, cell := range (*bot).CellsAroundHit {
		if cell.I == i && cell.J == j {
			return &cell
		}
	}
	return nil
}

func (bot *Bot) MarkCellHitted(aiFightBoard *domain.Board, i int, j int) {

	(*aiFightBoard).Cells[i][j].State = 4
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
		(*aiFightBoard).Cells[markI[k]][markJ[k]].State = 3
	}
}