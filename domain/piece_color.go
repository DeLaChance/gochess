package domain

type PieceColor int

const (
	WHITE PieceColor = iota
	BLACK
)

func (value PieceColor) String() string {
	if value == WHITE {
		return "W"
	} else {
		return "B"
	}
}

func pieceColorValues() []PieceColor {
	return []PieceColor{WHITE, BLACK}
}
