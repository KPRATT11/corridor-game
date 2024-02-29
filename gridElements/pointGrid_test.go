package gridElements

import (
	"corridor-game/types"
	"errors"
	"testing"
)

func TestPointGrid_CreateWall(t *testing.T) {
	cellGrid := MakeCellGrid(6)
	pointGrid := MakePointGrid(6, cellGrid)

	err := pointGrid.CreateWall([2]int{1, 1}, [2]int{1, 3})
	if err != nil {
		t.Error(err)
	}

	cell := cellGrid.GetAt([2]int{0, 1})
	if cell.SurroundingWalls[types.South].Exists != true {
		t.Error("expected south wall to be true")
	}

	if cell.SurroundingWalls[types.South].Ending != types.West {
		t.Error("expected wall ending to be west")
	}

	cell2 := cellGrid.GetAt([2]int{0, 2})
	if cell2.SurroundingWalls[types.South].Exists != true {
		t.Error("expected south wall to be true")
	}

	if cell2.SurroundingWalls[types.South].Ending != types.East {
		t.Error("expected wall ending to be east")
	}
}

// test the intersecting wall
func TestPointGrid_CreateWall2(t *testing.T) {
	cellGrid := MakeCellGrid(6)
	pointGrid := MakePointGrid(6, cellGrid)

	err := pointGrid.CreateWall([2]int{1, 1}, [2]int{1, 3})
	if err != nil {
		t.Error(err)
	}
	err2 := pointGrid.CreateWall([2]int{2, 2}, [2]int{0, 2})
	if errors.Is(err2, IntersectingWallErr) {
		t.Skip()
	}
	if err2 != nil {
		t.Error(err2)
	}
}
