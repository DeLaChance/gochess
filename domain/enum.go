package domain

type Enum struct {
}

func PieceTypeValues() []PieceType {
	return []PieceType{PAWN, ROOK, KNIGHT, BISHOP, KING, QUEEN}
}
