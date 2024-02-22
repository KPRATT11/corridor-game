package gridElements

func MakeCellGrid(boardSize int) [][]*Cell {
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

func MakePointGrid(boardSize int, cellGrid [][]*Cell) [][]*Point {
	grid := make([][]*Point, boardSize)

	for row := range grid {
		newRow := make([]*Point, boardSize)
		for col := range newRow {
			var point = NewPoint([2]int{row, col}, cellGrid)
			newRow[col] = &point
		}
		grid[row] = newRow
	}
	return grid
}
