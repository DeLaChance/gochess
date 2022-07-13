package repository

type ChessGameRepository interface {
	FindGameById(id int) ChessGameEntity
}
