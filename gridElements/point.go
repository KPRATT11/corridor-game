package gridElements

type Point struct {
	SurroundingCells [4]*Cell
}

func NewPoint(coords [2]int, cellGrid [][]*Cell) Point {
	return Point{SurroundingCells: getCellsSurroundingPoint(coords, cellGrid)}
}

func getCellsSurroundingPoint(coords [2]int, cellGrid [][]*Cell) [4]*Cell {
	var surrounds [4]*Cell
	if coords[0] != 0 && coords[1] != 0 {
		surrounds[0] = cellGrid[coords[0]-1][coords[1]-1]
	}
	if coords[1] != 0 {
		surrounds[1] = cellGrid[coords[0]][coords[1]-1]
	}
	if coords[0] != 0 {
		surrounds[2] = cellGrid[coords[0]-1][coords[1]]
	}
	surrounds[3] = cellGrid[coords[0]][coords[1]]
	return surrounds
}
