package gridElements

import (
	"corridor-game/types"
	"errors"
)

//Cell is the actual grid spot that the piece would go on, it only had knowledge
//It should be oblivious of anything other than the cell itself (including player position)

// NewCell is the function called to create a new cell with no walls surrounding it
func NewCell() Cell {
	var returnWalls [4]SurroundingWall
	for index := range returnWalls {
		returnWalls[index] = SurroundingWall{
			Exists: false,
            Ending: types.None,
		}
	}
	return Cell{SurroundingWalls: returnWalls}
}

type Cell struct {
	SurroundingWalls [4]SurroundingWall
}

var IntersectingWallErr = errors.New("wall is intersecting")

func (cell *Cell) updateWall(wallDirection types.Direction, ending types.Direction) error {
	isIntersecting, err := cell.isWallIntersecting(wallDirection)
	if err != nil {
		return err
	} else if isIntersecting {
		return IntersectingWallErr
	}
	cell.SurroundingWalls[wallDirection].Exists = true
	cell.SurroundingWalls[wallDirection].Ending = ending
	return nil
}

type SurroundingWall struct {
	Exists bool
	Ending types.Direction //the types that the wall ends at compared to the non end point
}

func (cell *Cell) isWallIntersecting(targetDirection types.Direction) (bool, error) {
	switch targetDirection {
	case types.North:
		if cell.SurroundingWalls[types.West].Ending != types.North && cell.SurroundingWalls[types.East].Ending != types.North {
			return false, nil
		} else {
			return true, nil
		}
	case types.East:
		if cell.SurroundingWalls[types.North].Ending != types.East && cell.SurroundingWalls[types.South].Ending != types.East {
			return false, nil
		} else {
			return true, nil
		}
	case types.South:
		if cell.SurroundingWalls[types.East].Ending != types.South && cell.SurroundingWalls[types.West].Ending != types.South {
			return false, nil
		} else {
			return true, nil
		}
	case types.West:
		if cell.SurroundingWalls[types.South].Ending != types.West && cell.SurroundingWalls[types.North].Ending != types.West {
			return false, nil
		} else {
			return true, nil
		}
	default:
		return false, errors.New("unable to tell if wall is intersecting")
	}
}
