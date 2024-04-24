package pathFinding

import "testing"

func Test_findDistanceBetweenPoints(t *testing.T) {
	startPosition := [2]int{1, 1}
    endPosition := [2]int{3, 8}
    distance := findDistanceBetweenPoints(startPosition, endPosition)
    if distance!= 7.28 {
        t.Error("expected distance to be 2 but distance was", distance)
    }
}