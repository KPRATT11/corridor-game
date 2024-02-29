package main

import "corridor-game/gridElements"

func main() {
	cellGrid := gridElements.MakeCellGrid(6)
	pointGrid := gridElements.MakePointGrid(6, cellGrid)
	print(pointGrid)
}
