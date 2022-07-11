package domain

type ChessGame struct {
	Board ChessBoard
	Turn  PieceColor
}

func (game *ChessGame) ApplyAction(action PieceAction) {
	game.Board.ApplyAction(action)
	game.swapTurns()
}

// Static methods
func InitialChessGame() ChessGame {
	return ChessGame{InitialChessBoard(), WHITE}
}

// Private methods
func (game *ChessGame) swapTurns() {
	if game.Turn == WHITE {
		game.Turn = BLACK
	} else {
		game.Turn = WHITE
	}
}
