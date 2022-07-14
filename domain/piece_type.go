package domain

type PieceType int

const (
	PAWN PieceType = iota
	ROOK
	KNIGHT
	BISHOP
	KING
	QUEEN
)

func (pieceType PieceType) GenerateHumanReadableDescription() string {
	switch {
	case pieceType == PAWN:
		return "Pawn"
	case pieceType == ROOK:
		return "Rook"
	case pieceType == KNIGHT:
		return "Knight"
	case pieceType == BISHOP:
		return "Bishop"
	case pieceType == QUEEN:
		return "Queen"
	}
	return "King"
}

func (value PieceType) PieceValue() int {
	switch {
	case value == PAWN:
		return 1
	case value == ROOK:
		return 3
	case value == KNIGHT:
	case value == BISHOP:
		return 2
	case value == QUEEN:
		return 8
	}

	return 100
}
