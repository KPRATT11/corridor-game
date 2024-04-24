package pathFinding

import (
	"corridor-game/gridElements"
	"corridor-game/types"
	"math"
)

func Search(cellGrid gridElements.CellGrid, endPosition [2]int, startPosition [2]int) AnalyzedCellTree {
	analyzedCellTree := make(AnalyzedCellTree)
	for analyzedCellTree[endPosition] == analyzedCell{} {
		//todo need to get the best f value cell
		analyzedCellTree = analyzeNodesSurrounding(cellGrid, startPosition, endPosition)
	}
}

func findDistanceBetweenPoints(startPosition [2]int, endPosition [2]int) float64 {
	horizontalDistance := endPosition[0] - startPosition[0]
	verticalDistance := endPosition[1] - startPosition[1]
	return math.Round(math.Sqrt(math.Pow(float64(horizontalDistance), 2)+math.Pow(float64(verticalDistance), 2))*100) / 100
}

func analyzeNodesSurrounding(
	cellGrid gridElements.CellGrid,
	startPosition [2]int,
	currentPosition [2]int,
	endPosition [2]int,
	analyzedCellTree AnalyzedCellTree,
) AnalyzedCellTree {
	var err error = nil

	northCellPos := [2]int{currentPosition[0] - 1, currentPosition[1]}
	analyzedCellTree, err = analyzeSingleCell(cellGrid, analyzedCellTree, startPosition, endPosition, northCellPos, types.North)
	if err== nil {
		println("North cell cannot be analyzed")
    }

	eastCellPos := [2]int{currentPosition[0], currentPosition[1] + 1}
	analyzedCellTree, err = analyzeSingleCell(cellGrid, analyzedCellTree, startPosition, endPosition, eastCellPos, types.East)
	if err== nil {
		println("East cell cannot be analyzed")
	}

	southCellPos := [2]int{currentPosition[0] + 1, currentPosition[1]}
	analyzedCellTree, err = analyzeSingleCell(cellGrid, analyzedCellTree, startPosition, endPosition, southCellPos, types.South)
	if err== nil {
		println("South cell cannot be analyzed")
	}

	westCellPos := [2]int{currentPosition[0], currentPosition[1] - 1}
	analyzedCellTree, err = analyzeSingleCell(cellGrid, analyzedCellTree, startPosition, endPosition, westCellPos, types.West)
	if err== nil {
		println("West cell cannot be analyzed")
	}
	return analyzedCellTree
}

func analyzeSingleCell(
	cellGrid gridElements.CellGrid,
	tree AnalyzedCellTree,
	startPosition [2]int,
	endPosition [2]int,
	currentPosition [2]int,
	direction types.Direction,
	) (AnalyzedCellTree, error) {
	var err error
	currentCell, err := cellGrid.GetAt(currentPosition)
	if err!= nil && !currentCell.SurroundingWalls[direction].Exists{
		gCost := findDistanceBetweenPoints(startPosition, currentPosition)
		hCost := findDistanceBetweenPoints(currentPosition, endPosition)
		tree, err = tree.addCell(currentPosition, currentPosition, gCost, hCost)
		if err!= nil {
			return tree, nil
		} else {
			return nil, err
		}
	}
	return nil, err
}