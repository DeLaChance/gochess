package domain

type PieceAction struct {
	FromPosition int // 0 top-left, 63 bottom-right
	ToPosition   int
}
