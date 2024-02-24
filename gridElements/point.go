package gridElements

import (
	"corridor-game/types"
	"errors"
)

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

func updateWallForSingleCell(point1 *Point, point2 *Point) {
	sharedCell, sharedCellErr := getSharedCell(point1, point2)
	var direction types.Direction

	if sharedCellErr != nil {
		UpdateWall(sharedCell)
	}
}

func CreateWall(point1 *Point, point2 *Point, point1Coords [2]int, point2Coords [2]int) {

}

//Will need to pass both points coords (or just add them into the point objet) for this func and the one above
func getDirection(startingPoint Point, endingPoint Point, startingCoords [2]int, endingCoords [2]int) types.Direction {
	if (startingPoint)
}

// there is def a better way to do this and I might refactor it later (I won't)
func getSharedCell(point1 *Point, point2 *Point) (*Cell, error) {
	for i := range point1.SurroundingCells {
		for y := range point2.SurroundingCells {
			if point1.SurroundingCells[i] == point2.SurroundingCells[y] {
				return point1.SurroundingCells[i], nil
			}
		}
	}
	return nil, errors.New("cannot find and shared cell between the provided points")
}
