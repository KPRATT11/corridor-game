package pathFinding

import "errors"

type AnalyzedCellTree map[[2]int]analyzedCell

type analyzedCell struct {
	gCost float64
	hCost float64
	childNodeKey [2]int
}

func (analyzedCellTree AnalyzedCellTree) addCell(cellKey [2]int, childNode [2]int, gCost float64, hCost float64) (AnalyzedCellTree, error) {
	analyzedCellTree[cellKey] = analyzedCell{gCost, hCost, childNode}
	return analyzedCellTree, nil
}

func (analyzedCellTree AnalyzedCellTree) getChildNode(cellKey [2]int) (analyzedCell, error) {
	value, found :=  analyzedCellTree[cellKey]
	if !found {
		return analyzedCell{}, errors.New("cell not found")
	}
	return value, nil
}

func (analyzedCellTree AnalyzedCellTree) getBestChildNode() analyzedCell {
	var bestChildNode analyzedCell
    for _, value := range analyzedCellTree {
        if value.gCost + value.hCost < bestChildNode.gCost + bestChildNode.hCost {
            bestChildNode = value
        }
    }
    return bestChildNode
}