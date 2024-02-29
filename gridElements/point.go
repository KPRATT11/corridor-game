package gridElements

import "C"

type Point struct {
	SurroundingCells [4]*Cell
}

func NewPoint(coords [2]int, cellGrid [][]*Cell, boardSize int) Point {
	return Point{SurroundingCells: getCellsSurroundingPoint(coords, cellGrid, boardSize)}
}

func getCellsSurroundingPoint(coords [2]int, cellGrid [][]*Cell, boardSize int) [4]*Cell {
	var surrounds [4]*Cell
	//Top left
	if coords[0] != 0 && coords[1] != 0 {
		surrounds[0] = cellGrid[coords[0]-1][coords[1]-1]
	}
	//Bottom left
	if coords[1] != 0 && coords[0] < boardSize {
		surrounds[1] = cellGrid[coords[0]][coords[1]-1]
	}
	//Top right
	if coords[0] != 0 && coords[1] < boardSize && coords[0] < boardSize {
		surrounds[2] = cellGrid[coords[0]-1][coords[1]]
	}
	//Bottom right
	if coords[1] < boardSize && coords[0] < boardSize {
		surrounds[3] = cellGrid[coords[0]][coords[1]]
	}
	return surrounds
}

