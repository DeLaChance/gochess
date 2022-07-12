package domain

type ChessGame struct {
	Board       ChessBoard
	ActiveColor PieceColor
}

func (game *ChessGame) GeneratePossibleActions() []PieceAction {
	activeCells := game.Board.FindCellsByColor(game.ActiveColor)

	var pieceActions []PieceAction
	for _, activeCell := range activeCells {
		pieceActions = game.Board.GeneratePossibleActions(activeCell)
	}

	return pieceActions
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
	if game.ActiveColor == WHITE {
		game.ActiveColor = BLACK
	} else {
		game.ActiveColor = WHITE
	}
}
