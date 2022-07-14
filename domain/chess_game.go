package domain

const MINIMUM_SCORE = -1000

type ChessGame struct {
	ID                  uint
	Board               ChessBoard
	ActiveColor         PieceColor
	Actions             []PieceMoveAction
	CapturedWhitePieces []ChessPiece
	CapturedBlackPieces []ChessPiece
	GameResult          GameResult
	WhitePlayer         ChessAIPlayer
	BlackPlayer         ChessAIPlayer
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

func (game *ChessGame) AdvanceToNextTurn() {
	if game.ActiveColor == WHITE {
		game.ActiveColor = BLACK
	} else {
		game.ActiveColor = WHITE
	}
}

// Static methods
func InitialChessGame(id uint, whitePlayer ChessAIPlayer, blackPlayer ChessAIPlayer) ChessGame {
	return ChessGame{
		ID:                  id,
		Board:               InitialChessBoard(),
		ActiveColor:         WHITE,
		Actions:             make([]PieceMoveAction, 0),
		CapturedWhitePieces: make([]ChessPiece, 0),
		CapturedBlackPieces: make([]ChessPiece, 0),
		GameResult:          UNDETERMINED,
		WhitePlayer:         whitePlayer,
		BlackPlayer:         blackPlayer,
	}
}
