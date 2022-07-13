package repository

import (
	"config"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlChessGameRepository struct {
}

func (repo MySqlChessGameRepository) A() string {
	return "hello_world"
}

func GenerateMySQLChessGameRepository(aConfig *config.Config) *MySqlChessGameRepository {

	databaseUrl := aConfig.DatabaseUrl

	db, err := sql.Open("mysql", databaseUrl)
	if err == nil {
		config.Info.Printf("Successfully connected to %s", databaseUrl)
	} else {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	chessGameRepository := MySqlChessGameRepository{}
	return &chessGameRepository
}
