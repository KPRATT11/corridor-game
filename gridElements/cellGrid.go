package gridElements

import "errors"

type CellGrid [][]*Cell

func MakeCellGrid(boardSize int) CellGrid {
	grid := make([][]*Cell, boardSize)

	for col := range grid {
		row := make([]*Cell, boardSize)
		for spot := range row {
			var cell = NewCell()
			row[spot] = &cell
		}
		grid[col] = row
	}
	return grid
}

func (cellGrid *CellGrid) GetAt(coords [2]int) (*Cell, error) {
	if !cellGrid.isInBounds(coords){
        return nil, errors.New("out of bounds")
    } else {
		return (*cellGrid)[coords[0]][coords[1]], nil
	}
}

func (cellGrid *CellGrid) isInBounds(coords [2]int) bool {
	if coords[0] < 0 || coords[0] >= len(*cellGrid) || coords[1] < 0 || coords[1] >= len((*cellGrid)[coords[0]]) {
        return false
    } else {
        return true
    }
}