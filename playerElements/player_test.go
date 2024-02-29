package playerElements

import (
    "corridor-game/gridElements"
    "corridor-game/types"
    "testing"
)

func TestPlayer_UpdatePosition(t *testing.T) {
	player := NewPlayer(6, true, 9)
    cellGrid := gridElements.MakeCellGrid(6)

    err := player.UpdatePosition(types.South, cellGrid)
    if err!= nil {
        t.Error(err)
    }

    if player.Position[0]!= 1 {
        t.Error("expected position to be 0")
    }

    if player.Position[1]!= 3 {
        t.Error("expected position to be 1")
    }

}

func TestPlayer_UpdatePosition2(t *testing.T) {
    player := NewPlayer(6, true, 9)
    cellGrid := gridElements.MakeCellGrid(6)
    pointGrid := gridElements.MakePointGrid(6, cellGrid)

    err := pointGrid.CreateWall([2]int{1, 3}, [2]int{1, 5})
    if err!= nil {
        t.Error(err)
    }

    err = player.UpdatePosition(types.South, cellGrid)
    if err == nil {
        t.Error("expected an error")
    }
    t.Skip()
}