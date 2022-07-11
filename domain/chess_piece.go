package domain

type ChessPiece struct {
	Type  PieceType
	Color PieceColor
}

func (value ChessPiece) String() string {
	if value.Color == WHITE {
		switch {
		case value.Type == PAWN:
			return "♙"
		case value.Type == ROOK:
			return "♖"
		case value.Type == KNIGHT:
			return "♘"
		case value.Type == BISHOP:
			return "♙"
		case value.Type == QUEEN:
			return "♕"
		}

		return "♔"
	} else {
		switch {
		case value.Type == PAWN:
			return "♟"
		case value.Type == ROOK:
			return "♜"
		case value.Type == KNIGHT:
			return "♞"
		case value.Type == BISHOP:
			return "♝"
		case value.Type == QUEEN:
			return "♛"
		}

		return "♚"

	}
}
