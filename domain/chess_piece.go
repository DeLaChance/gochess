package domain

type ChessPiece struct {
	Type  PieceType
	Color PieceColor
}

func (piece ChessPiece) GenerateHumanReadableDescription() string {
	switch {
	case piece.Type == PAWN:
		return "Pawn"
	case piece.Type == ROOK:
		return "Rook"
	case piece.Type == KNIGHT:
		return "Knight"
	case piece.Type == BISHOP:
		return "Bishop"
	case piece.Type == QUEEN:
		return "Queen"
	}
	return "King"
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
