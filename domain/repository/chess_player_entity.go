package repository

import "gorm.io/gorm"

type ChessPlayerEntity struct {
	gorm.Model
	Name string
	Type string
}
