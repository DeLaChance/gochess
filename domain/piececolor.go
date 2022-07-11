package domain

type PieceColor int

const (
	WHITE PieceColor = iota
	BLACK
)

func pieceColorValues() []PieceColor {
	return []PieceColor{WHITE, BLACK}
}
