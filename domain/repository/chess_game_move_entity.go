package repository

import "gorm.io/gorm"

type ChessGameMoveEntity struct {
	gorm.Model
	FromPosition int
	ToPosition   int
}
