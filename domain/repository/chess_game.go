package repository

type ChessGame struct {
	ID    uint
	Moves []ChessGameMoveEntity
}

func GenerateChessGame(chessGameEntity ChessGameEntity, moves []ChessGameMoveEntity) ChessGame {
	return ChessGame{ID: chessGameEntity.ID, Moves: moves}
}
