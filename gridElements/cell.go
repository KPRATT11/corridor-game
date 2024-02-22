package gridElements

import "corridor-game/types"

//Cell is the actual grid spot that the piece would go on, it only had knowledge
//It should be oblivious of anything other than the cell itself (including player position)

// NewCell is the function called to create a new cell with no walls surrounding it
func NewCell() Cell {
	var returnWalls [4]SurroundingWall
	for _, value := range returnWalls {
		value.Ending = types.None
		value.Exists = false
	}
	return Cell{SurroundingWalls: returnWalls}
}

type Cell struct {
	SurroundingWalls [4]SurroundingWall
}

// UpdateWall takes in a direction (the side that the wall should be updated) and an ending direction
func UpdateWall(cell *Cell, direction types.Direction, ending types.Direction) {
	cell.SurroundingWalls[direction].Exists = true
	cell.SurroundingWalls[direction].Ending = ending
}

type SurroundingWall struct {
	Exists bool
	Ending types.Direction //the types that the wall ends at compared to the non end point
}
