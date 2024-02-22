package main

import "corridor-game/gridElements"

func main() {
	cellGrid := gridElements.MakeCellGrid(6)
	pointGrid := gridElements.MakePointGrid(6, cellGrid) //todo we need to increase how much we are getting for points I think there needs to be 1 more point than there is in the cell grid in each direction
	print(pointGrid)
}
