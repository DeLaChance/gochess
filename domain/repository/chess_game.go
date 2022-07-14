package repository

type ChessGame struct {
	ID          uint
	Moves       []ChessGameMoveEntity
	WhitePlayer ChessPlayerEntity
	BlackPlayer ChessPlayerEntity
}
