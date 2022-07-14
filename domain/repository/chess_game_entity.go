package repository

type ChessGameEntity struct {
	ID            uint   `gorm:"column:id"`
	WhitePlayerID uint   `gorm:"column:white_player"`
	BlackPlayerID uint   `gorm:"column:black_player"`
	Result        string `gorm:"column:result"`
}

func (ChessGameEntity) TableName() string {
	return "chess_game"
}
