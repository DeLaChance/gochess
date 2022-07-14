package repository

type ChessGame struct {
	ID          uint
	Moves       []ChessGameMoveEntity
	Result      string
	WhitePlayer ChessPlayerEntity
	BlackPlayer ChessPlayerEntity
}
