package repository

import (
	"config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlChessGameRepository struct {
	db *gorm.DB
}

func (repo MySqlChessGameRepository) A() string {
	return "hello_world"
}

func GenerateMySQLChessGameRepository(aConfig *config.Config) *MySqlChessGameRepository {

	db, err := gorm.Open(mysql.Open(aConfig.DatabaseUrl), &gorm.Config{})

	if err == nil {
		config.Info.Printf("Successfully connected to %s", aConfig.DatabaseUrl)
	} else {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&ChessPlayerEntity{})
	db.AutoMigrate(&ChessGameEntity{})
	db.AutoMigrate(&ChessGameMoveEntity{})

	insertInitialData(db)

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	chessGameRepository := MySqlChessGameRepository{}
	return &chessGameRepository
}

func insertInitialData(db *gorm.DB) {

	sessionDB := db.Session(&gorm.Session{CreateBatchSize: 1000})

	count := int64(0)
	sessionDB.Model(&ChessGameEntity{}).Count(&count)

	if count == 0 {

		player1 := ChessPlayerEntity{Name: "C3P0", Type: "AI"}
		player2 := ChessPlayerEntity{Name: "R2-D2", Type: "AI"}

		firstMove := ChessGameMoveEntity{FromPosition: 8, ToPosition: 24}
		secondMove := ChessGameMoveEntity{FromPosition: 56, ToPosition: 32}

		game := ChessGameEntity{WhitePlayer: player1, BlackPlayer: player2, Moves: []ChessGameMoveEntity{firstMove, secondMove}}
		sessionDB.Save(&game)
	}
}
