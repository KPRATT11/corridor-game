package playerElements

import (
	"corridor-game/gridElements"
	"corridor-game/types"
	"errors"
)

type Player struct {
	Position [2]int
	Walls int
}

func NewPlayer(boardSize int, player1 bool, walls int) Player {
	if player1 {
		return Player{Position: [2]int{0, boardSize / 2}, Walls: walls}
	} else {
		return Player{Position: [2]int{boardSize - 1, boardSize / 2}, Walls: walls}
	}
}

func (player *Player) UpdatePosition(direction types.Direction, cellGrid gridElements.CellGrid) error {
	switch direction {
    case types.North:
		if cellGrid[player.Position[0]][player.Position[1]].SurroundingWalls[types.North].Exists {
            return errors.New("can't go north")
        } else {
            player.Position[0]--
        }
    case types.East:
		if cellGrid[player.Position[0]][player.Position[1]].SurroundingWalls[types.East].Exists {
			return errors.New("can't go east")
		} else {
			player.Position[1]++
		}
    case types.South:
		if cellGrid[player.Position[0]][player.Position[1]].SurroundingWalls[types.South].Exists {
            return errors.New("can't go south")
        } else {
            player.Position[0]++
        }
    case types.West:
		if cellGrid[player.Position[0]][player.Position[1]].SurroundingWalls[types.West].Exists {
            return errors.New("can't go west")
        } else {
            player.Position[1]--
        }
    default:
        return errors.New("unable to update player position")
    }
    return nil
}