package gridElements

import "C"

type Point struct {
	SurroundingCells [4]*Cell
}

func NewPoint(coords [2]int, cellGrid CellGrid, boardSize int) Point {
	return Point{SurroundingCells: getCellsSurroundingPoint(coords, cellGrid, boardSize)}
}

func getCellsSurroundingPoint(coords [2]int, cellGrid CellGrid, boardSize int) [4]*Cell {
	var surrounds [4]*Cell
	//Top left
	if cellGrid.isInBounds([2]int{coords[0] - 1, coords[1] - 1}) {
		surrounds[0] = cellGrid[coords[0]-1][coords[1]-1]
	}
	//Bottom left
	if cellGrid.isInBounds([2]int{coords[0], coords[1]-1}) {
		surrounds[1] = cellGrid[coords[0]][coords[1]-1]
	}
	//Top right
	if cellGrid.isInBounds([2]int{coords[0] - 1, coords[1]}) {
		surrounds[2] = cellGrid[coords[0]-1][coords[1]]
	}
	//Bottom right
	if cellGrid.isInBounds([2]int{coords[0], coords[1]}) {
		surrounds[3] = cellGrid[coords[0]][coords[1]]
	}
	return surrounds
}

