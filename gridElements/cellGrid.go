package gridElements

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

func (cellGrid *CellGrid) GetAt(coords [2]int) *Cell {
	return (*cellGrid)[coords[0]][coords[1]]
}

