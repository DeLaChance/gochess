package domain

type ChessPlayer interface {
	PlayerColor() PieceColor
	PlayerName() string

	ChooseAction(chessGame ChessGame) *PieceMoveAction
}
