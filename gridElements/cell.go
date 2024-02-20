package gridElements

import "corridor-game/types"

func NewEmptyCell() Cell {
	var returnWalls [4]SurroundingWall
	for index, value := range returnWalls {
		value.Direction = types.Direction(index)
		value.Ending = types.None
	}
	return Cell{SurroundingWalls: returnWalls}
}

type Cell struct {
	SurroundingWalls [4]SurroundingWall
}

type SurroundingWall struct {
	Exists    bool
	Direction types.Direction
	Ending    types.Direction //the types that the wall ends at compared to the non end point
}
