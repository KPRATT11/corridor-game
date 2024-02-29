package gridElements

import (
	"corridor-game/types"
	"errors"
)

type PointGrid [][]*Point

func MakePointGrid(boardSize int, cellGrid [][]*Cell) PointGrid {
	grid := make([][]*Point, boardSize+1)

	for row := range grid {
		newRow := make([]*Point, boardSize+1)
		for col := range newRow {
			var point = NewPoint([2]int{row, col}, cellGrid, boardSize)
			newRow[col] = &point
		}
		grid[row] = newRow
	}
	return grid
}

func (pointGrid *PointGrid) GetAt(coords [2]int) *Point {
	return (*pointGrid)[coords[0]][coords[1]]
}

//todo need to add some validation for the size of the walls
func (pointGrid *PointGrid) CreateWall(startPointCoords [2]int, endPointCoords [2]int) error {
	wallDirection := getDirection(startPointCoords, endPointCoords)

	//TODO need to make sure we wont have out of bound exceptions
	var middlePointCoords [2]int
	switch wallDirection {
	case types.North:
		middlePointCoords = [2]int{startPointCoords[0] - 1, startPointCoords[1]}
	case types.East:
		middlePointCoords = [2]int{startPointCoords[0], startPointCoords[1] + 1}
	case types.South:
		middlePointCoords = [2]int{startPointCoords[0] + 1, startPointCoords[1]}
	case types.West:
		middlePointCoords = [2]int{startPointCoords[0], startPointCoords[1] - 1}
	case types.None:
		return errors.New("unable to get a direction between 2 points")
	}

	//TODO need to check for intersecting walls
	var err error
	//Updates first cell
	err = pointGrid.updateWallForSingleCell(startPointCoords, middlePointCoords, false)
	if err != nil {
		return err
	}
	//Updates second cell
	err = pointGrid.updateWallForSingleCell(middlePointCoords, endPointCoords, true)
	if err != nil {
		return err
	}
	return nil
}

func (pointGrid *PointGrid) updateWallForSingleCell(startingCoords [2]int, endingCoords [2]int, lastWall bool) error {
	point1 := pointGrid.GetAt(startingCoords)
	point2 := pointGrid.GetAt(endingCoords)

	sharedCells, wallDirections, err := getSharedCell(point1, point2)
	if err != nil {
		return err
	}

	var endingDirection types.Direction //ensures that it always takes the middle point (out of the 3 points called in CreateWall) to use at the starting coord in getDirection
	if !lastWall {
		endingDirection = getDirection(endingCoords, startingCoords)
	} else {
		endingDirection = getDirection(startingCoords, endingCoords)
	}

	for i := range sharedCells {
		err := sharedCells[i].updateWall(wallDirections[i], endingDirection)
		if err!= nil {
            return err
        }
	}
	return nil
}

func getDirection(startingCoords [2]int, endingCoords [2]int) types.Direction {
	switch {
	case startingCoords[0] > endingCoords[0]:
		return types.North
	case startingCoords[1] < endingCoords[1]:
		return types.East
	case startingCoords[0] < endingCoords[0]:
		return types.South
	case startingCoords[1] > endingCoords[1]:
		return types.West
	default:
		return types.None
	}
}

// there is def a better way to do this and I might refactor it later (I won't)
func getSharedCell(point1 *Point, point2 *Point) ([]*Cell, []types.Direction, error) {
	var sharedCells []*Cell
	var sharedDirections []types.Direction
	for i := range point1.SurroundingCells {
		for y := range point2.SurroundingCells {
			if point1.SurroundingCells[i] == point2.SurroundingCells[y] {
				sharedDirections = append(sharedDirections, types.Direction(i))
				sharedCells = append(sharedCells, point1.SurroundingCells[i])
			}
		}
	}
	if len(sharedCells) == 0 {
		return nil, nil, errors.New("could not find shared cell")
	}
	return sharedCells, sharedDirections, nil
}
