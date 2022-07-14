package repository

type ChessPlayerEntity struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Type string `gorm:"column:type"`
}

func (ChessPlayerEntity) TableName() string {
	return "chess_player"
}
