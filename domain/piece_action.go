package domain

type PieceAction struct {
	FromPosition int // 0 top-left, 63 bottom-right
	ToPosition   int
}

func GeneratePieceAction(oldX int, oldY int, newX int, newY int) PieceAction {
	return PieceAction{convertToOneDimensionalPosition(oldX, oldY), convertToOneDimensionalPosition(newX, newY)}
}

// Private methods
func convertToOneDimensionalPosition(x int, y int) int {
	return y*BOARD_DIMENSION + x
}
