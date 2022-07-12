package domain

type ChessAIPlayer struct {
	Color PieceColor
	Name  string
}

func (player ChessAIPlayer) PlayerColor() PieceColor {
	return player.Color
}

func (player ChessAIPlayer) PlayerName() string {
	return player.Name
}

func (player ChessAIPlayer) ChooseAction(chessGame ChessGame) *PieceMoveAction {
	return player.generateBestPossibleAction(chessGame, chessGame.GeneratePossibleActions())
}

// private method
func (player ChessAIPlayer) generateBestPossibleAction(chessGame ChessGame, actions []PieceMoveAction) *PieceMoveAction {

	maxScore := MINIMUM_SCORE
	var chosenAction *PieceMoveAction

	for _, action := range actions {
		chessGame.ApplyMoveAction(action)
		score := chessGame.CalculateScore(player.Color)

		if score > maxScore {
			chosenAction = &action
			maxScore = score
		}
		chessGame.UnapplyLastMoveAction()

	}

	return chosenAction
}
