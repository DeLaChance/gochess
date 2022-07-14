package domain

type ChessPiece struct {
	Type  PieceType
	Color PieceColor
}

func (piece ChessPiece) String() string {
	if piece.Color == WHITE {
		switch {
		case piece.Type == PAWN:
			return "♙"
		case piece.Type == ROOK:
			return "♖"
		case piece.Type == KNIGHT:
			return "♘"
		case piece.Type == BISHOP:
			return "♙"
		case piece.Type == QUEEN:
			return "♕"
		}

		return "♔"
	} else {
		switch {
		case piece.Type == PAWN:
			return "♟"
		case piece.Type == ROOK:
			return "♜"
		case piece.Type == KNIGHT:
			return "♞"
		case piece.Type == BISHOP:
			return "♝"
		case piece.Type == QUEEN:
			return "♛"
		}

		return "♚"

	}
}
