package repository

import (
	"config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlChessGameRepository struct {
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

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	chessGameRepository := MySqlChessGameRepository{}
	return &chessGameRepository
}
