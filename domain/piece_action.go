package domain

import "fmt"

type PieceMoveAction struct {
	FromPosition int // 0 top-left, 63 bottom-right
	ToPosition   int
}

func (action PieceMoveAction) GenerateReversedAction() PieceMoveAction {
	return PieceMoveAction{action.ToPosition, action.FromPosition}
}

func GeneratePieceMoveAction(oldX int, oldY int, newX int, newY int) PieceMoveAction {
	return PieceMoveAction{convertToOneDimensionalPosition(oldX, oldY), convertToOneDimensionalPosition(newX, newY)}
}

func (action PieceMoveAction) String() string {
	return fmt.Sprintf("Move from %d to %d \n", action.FromPosition, action.ToPosition)
}

// Private methods
func convertToOneDimensionalPosition(x int, y int) int {
	return y*BOARD_DIMENSION + x
}
