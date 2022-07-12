package domain

const MINIMUM_SCORE = -1000

type ChessGame struct {
	Board               ChessBoard
	ActiveColor         PieceColor
	Actions             []PieceMoveAction
	CapturedWhitePieces []ChessPiece
	CapturedBlackPieces []ChessPiece
	GameResult          GameResult
}

func (game *ChessGame) StartGame(whitePlayer ChessPlayer, blackPlayer ChessPlayer) {
	for game.GameResult == UNDETERMINED {

		Info.Printf("Player turn is %s", game.ActiveColor.String())
		Info.Printf("Board: \n" + game.Board.String() + "\n")

		var activePlayer ChessPlayer

		if game.ActiveColor == WHITE {
			activePlayer = whitePlayer
		} else {
			activePlayer = blackPlayer
		}

		chosenAction := activePlayer.ChooseAction(*game)
		if chosenAction == nil {
			game.GameResult = DRAW
		} else {
			Info.Println(chosenAction.String())

			game.ApplyMoveAction(*chosenAction)
		}

		game.swapTurns()
	}
}

func (game *ChessGame) GeneratePossibleActions() []PieceMoveAction {
	activeCells := game.Board.FindCellsByColor(game.ActiveColor)

	var pieceActions []PieceMoveAction
	for _, activeCell := range activeCells {
		pieceActions = append(pieceActions, game.Board.GeneratePossibleActions(activeCell)...)
	}

	return pieceActions
}

func (game *ChessGame) ApplyMoveAction(action PieceMoveAction) {
	game.Actions = append(game.Actions, action)

	capturedPiece := game.Board.ApplyMoveAction(action)

	Info.Printf("Board: \n" + game.Board.String() + "\n")

	if capturedPiece != nil {
		if capturedPiece.Color == BLACK {
			game.CapturedWhitePieces = append(game.CapturedWhitePieces, *capturedPiece)
		} else {
			game.CapturedBlackPieces = append(game.CapturedBlackPieces, *capturedPiece)
		}
	}
}

func (game *ChessGame) UnapplyLastMoveAction() {

	// Pop element of list: x, a = a[len(a)-1], a[:len(a)-1]
	lastAction := game.Actions[len(game.Actions)-1]
	game.Actions = game.Actions[:len(game.Actions)-1]

	var lastCapturedPiece *ChessPiece
	if game.ActiveColor == BLACK && len(game.CapturedWhitePieces) > 0 {
		lastCapturedPiece = &game.CapturedWhitePieces[len(game.CapturedWhitePieces)-1]
		game.CapturedWhitePieces = game.CapturedWhitePieces[:len(game.CapturedWhitePieces)-1]
	} else if game.ActiveColor == WHITE && len(game.CapturedBlackPieces) > 0 {
		lastCapturedPiece = &game.CapturedBlackPieces[len(game.CapturedBlackPieces)-1]
		game.CapturedBlackPieces = game.CapturedBlackPieces[:len(game.CapturedBlackPieces)-1]
	} else {
		lastCapturedPiece = nil
	}

	game.Board.UnapplyMoveAction(lastAction, lastCapturedPiece)
}

func (game *ChessGame) CalculateScore(color PieceColor) int {
	return game.Board.CalculateScore(color)
}

// Static methods
func InitialChessGame() ChessGame {
	return ChessGame{
		Board:               InitialChessBoard(),
		ActiveColor:         WHITE,
		Actions:             make([]PieceMoveAction, 8),
		CapturedWhitePieces: make([]ChessPiece, 8),
		CapturedBlackPieces: make([]ChessPiece, 8),
		GameResult:          UNDETERMINED}
}

// Private methods
func (game *ChessGame) swapTurns() {
	if game.ActiveColor == WHITE {
		game.ActiveColor = BLACK
	} else {
		game.ActiveColor = WHITE
	}
}
