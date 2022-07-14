package repository

type ChessGameMoveEntity struct {
	ID           uint `gorm:"column:id"`
	GameID       uint `gorm:"column:game_id"`
	FromPosition int  `gorm:"column:from_position"`
	ToPosition   int  `gorm:"column:to_position"`
}

func (ChessGameMoveEntity) TableName() string {
	return "chess_game_moves"
}
