package domain

import "math/rand"

type ChessAIPlayer struct {
	ID    uint
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
	possibleActions := shuffle(chessGame.GeneratePossibleActions())
	return player.generateBestPossibleAction(chessGame, possibleActions)
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

func shuffle(src []PieceMoveAction) []PieceMoveAction {
	dest := make([]PieceMoveAction, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}
