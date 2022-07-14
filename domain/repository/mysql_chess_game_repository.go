package repository

import (
	"config"
	"domain"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlChessGameRepository struct {
	db *gorm.DB
}

// TODO: refactor these confusing methods.
func (repo *MySqlChessGameRepository) FindGameById(id int) (ChessGame, error) {
	gameEntity, error := repo.FindGameEntityById(id)

	if error == nil {
		moveEntities := repo.findMovesByGameId(id)
		whitePlayer, player1NotFoundError := repo.FindPlayerById(gameEntity.WhitePlayerID)
		blackPlayer, player2NotFoundError := repo.FindPlayerById(gameEntity.BlackPlayerID)

		if player1NotFoundError != nil {
			return ChessGame{}, player1NotFoundError
		} else if player2NotFoundError != nil {
			return ChessGame{}, player2NotFoundError
		} else {
			return ChessGame{ID: gameEntity.ID, Moves: moveEntities, WhitePlayer: whitePlayer, BlackPlayer: blackPlayer, Result: gameEntity.Result}, nil
		}

	} else {
		return ChessGame{}, error
	}
}

// TODO: refactor these confusing methods.
func (repo *MySqlChessGameRepository) FindGameEntityById(id int) (ChessGameEntity, error) {

	var chessGame ChessGameEntity

	repo.db.First(&chessGame, id)
	if chessGame.ID == 0 {
		return chessGame, errors.New("Not found")
	} else {
		return chessGame, nil
	}
}

func (repo *MySqlChessGameRepository) FindPlayerById(id uint) (ChessPlayerEntity, error) {

	var player ChessPlayerEntity
	repo.db.First(&player, id)
	if player.ID == 0 {
		return player, errors.New("Not found")
	} else {
		return player, nil
	}
}

func (repo *MySqlChessGameRepository) CreateNewGame(whitePlayer ChessPlayerEntity, blackPlayer ChessPlayerEntity) ChessGame {
	chessGameEntity := ChessGameEntity{WhitePlayerID: whitePlayer.ID, BlackPlayerID: blackPlayer.ID}
	repo.db.Create(&chessGameEntity)
	return ChessGame{ID: chessGameEntity.ID, WhitePlayer: whitePlayer, BlackPlayer: blackPlayer, Result: chessGameEntity.Result}
}

func (repo *MySqlChessGameRepository) SaveMoveEntity(moveEntity ChessGameMoveEntity) {
	repo.db.Create(&moveEntity)
}

func (repo *MySqlChessGameRepository) SaveGameEntity(gameEntity ChessGameEntity) {
	repo.db.Save(&gameEntity)
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

	game := ChessGameEntity{WhitePlayerID: player1.ID, BlackPlayerID: player2.ID, Result: domain.UNDETERMINED.String()}
	db.FirstOrCreate(&game, &game)

	firstMove := ChessGameMoveEntity{GameID: game.ID, FromPosition: 8, ToPosition: 24}
	secondMove := ChessGameMoveEntity{GameID: game.ID, FromPosition: 56, ToPosition: 32}

	db.FirstOrCreate(&firstMove, &firstMove)
	db.FirstOrCreate(&secondMove, &secondMove)
}

func (repo *MySqlChessGameRepository) findMovesByGameId(id int) []ChessGameMoveEntity {
	var moves []ChessGameMoveEntity

	repo.db.Where("game_id = ?", id).Find(&moves)
	return moves
}
