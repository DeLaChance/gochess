package domain

type PieceColor int

const (
	WHITE PieceColor = iota
	BLACK
)

func (value PieceColor) generateOpposite() PieceColor {
	if value == WHITE {
		return BLACK
	} else {
		return WHITE
	}
}

func (value PieceColor) String() string {
	if value == WHITE {
		return "White"
	} else {
		return "Black"
	}
}

func GeneratePieceColorValues() []PieceColor {
	return []PieceColor{WHITE, BLACK}
}
