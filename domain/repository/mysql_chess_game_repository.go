package repository

import (
	"config"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlChessGameRepository struct {
	db *gorm.DB
}

func (repo *MySqlChessGameRepository) FindGameById(id int) (ChessGame, error) {
	gameEntity, error := repo.findGameEntityById(id)

	if error == nil {
		moveEntities := repo.findMovesByGameId(id)
		return GenerateChessGame(gameEntity, moveEntities), nil
	} else {
		return ChessGame{}, error
	}
}

func GenerateMySQLChessGameRepository(aConfig *config.Config) *MySqlChessGameRepository {

	db, err := gorm.Open(mysql.Open(aConfig.DatabaseUrl), &gorm.Config{})

	if err == nil {
		config.Info.Printf("Successfully connected to %s", aConfig.DatabaseUrl)
	} else {
		panic(err)
	}

	db = db.Debug()

	insertInitialData(db)

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	chessGameRepository := MySqlChessGameRepository{db: db}
	return &chessGameRepository
}

func insertInitialData(db *gorm.DB) {

	player1 := ChessPlayerEntity{Name: "C3P0", Type: "AI"}
	player2 := ChessPlayerEntity{Name: "R2-D2", Type: "AI"}

	db.FirstOrCreate(&player1, &player1)
	db.FirstOrCreate(&player2, &player2)

	game := ChessGameEntity{WhitePlayerID: player1.ID, BlackPlayerID: player2.ID}
	db.FirstOrCreate(&game, &game)

	firstMove := ChessGameMoveEntity{GameID: game.ID, FromPosition: 8, ToPosition: 24}
	secondMove := ChessGameMoveEntity{GameID: game.ID, FromPosition: 56, ToPosition: 32}

	db.FirstOrCreate(&firstMove, &firstMove)
	db.FirstOrCreate(&secondMove, &secondMove)
}

func (repo *MySqlChessGameRepository) findMovesByGameId(id int) []ChessGameMoveEntity {
	var moves []ChessGameMoveEntity

	repo.db.Find(&moves).Where("game_id = ?", id)
	return moves
}

func (repo *MySqlChessGameRepository) findGameEntityById(id int) (ChessGameEntity, error) {
	var chessGame ChessGameEntity

	repo.db.First(&chessGame, id)
	if chessGame.ID == 0 {
		return chessGame, errors.New("Not found")
	} else {
		return chessGame, nil
	}
}
