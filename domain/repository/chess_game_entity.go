package repository

import "gorm.io/gorm"

type ChessGameEntity struct {
	gorm.Model
	WhitePlayer ChessPlayerEntity     `gorm:"foreignKey:ID"`
	BlackPlayer ChessPlayerEntity     `gorm:"foreignKey:ID"`
	Moves       []ChessGameMoveEntity `gorm:"foreignKey:ID"`
}
