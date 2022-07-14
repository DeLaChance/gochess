package service

import (
	"config"
	"domain"
	"repository"
)

type ChessGameService struct {
	Repository repository.MySqlChessGameRepository
}

func (service *ChessGameService) StartNewGame(id int) (domain.ChessGame, error) {
	game, error := service.FetchById(id)
	if error == nil {
		startGameAsync(game, &service.Repository)
		return game, nil
	} else {
		return domain.ChessGame{}, error
	}
}

func (service *ChessGameService) FetchById(id int) (domain.ChessGame, error) {
	chessGameEntity, error := service.Repository.FindGameById(id)

	var chessGame domain.ChessGame
	if error == nil {
		generatedChessGame := GenerateChessGame(chessGameEntity)

		// Cannot do '&' on the return value of the method.
		chessGame = generatedChessGame
	} else {
		chessGame = domain.ChessGame{}
	}

	return chessGame, error
}

func (service *ChessGameService) CreateNewGame() (domain.ChessGame, error) {

	whitePlayer, whitePlayerNotFound := service.Repository.FindPlayerById(uint(1))
	blackPlayer, blackPlayerNotFound := service.Repository.FindPlayerById(uint(2))

	if whitePlayerNotFound != nil {
		return domain.ChessGame{}, whitePlayerNotFound
	} else if blackPlayerNotFound != nil {
		return domain.ChessGame{}, blackPlayerNotFound
	} else {
		chessGameEntity := service.Repository.CreateNewGame(whitePlayer, blackPlayer)
		generatedChessGame := GenerateChessGame(chessGameEntity)
		return generatedChessGame, nil
	}
}

// Static methods
func GenerateChessGameService(repository *repository.MySqlChessGameRepository) *ChessGameService {
	return &ChessGameService{Repository: *repository}
}

func GenerateChessGame(chessGameEntity repository.ChessGame) domain.ChessGame {
	whitePlayer := domain.ChessAIPlayer{ID: chessGameEntity.WhitePlayer.ID, Name: chessGameEntity.WhitePlayer.Name, Color: domain.WHITE}
	blackPlayer := domain.ChessAIPlayer{ID: chessGameEntity.BlackPlayer.ID, Name: chessGameEntity.BlackPlayer.Name, Color: domain.BLACK}
	chessGame := domain.InitialChessGame(chessGameEntity.ID, whitePlayer, blackPlayer)

	for _, move := range chessGameEntity.Moves {
		moveAction := GenerateMoveAction(move)
		chessGame.ApplyMoveAction(moveAction)
	}

	config.Info.Println(chessGameEntity.Result)
	chessGame.GameResult = domain.MapStringToGameResult(chessGameEntity.Result)
	config.Info.Println(chessGame.GameResult.String())

	return chessGame
}

func GenerateMoveAction(move repository.ChessGameMoveEntity) domain.PieceMoveAction {
	return domain.PieceMoveAction{FromPosition: move.FromPosition, ToPosition: move.ToPosition}
}

// Private methods
func startGameAsync(game domain.ChessGame, repo *repository.MySqlChessGameRepository) {
	go startGame(game, repo)
}

func startGame(game domain.ChessGame, repo *repository.MySqlChessGameRepository) {

	for game.GameResult == domain.UNDETERMINED {

		config.Info.Printf("Player %s can make a move. Turn number %d.", game.ActiveColor.String(), len(game.Actions)+1)
		config.Info.Printf("Board: \n" + game.Board.String() + "\n")

		var activePlayer domain.ChessAIPlayer

		if game.ActiveColor == domain.WHITE {
			activePlayer = game.WhitePlayer
		} else {
			activePlayer = game.BlackPlayer
		}

		chosenAction := activePlayer.ChooseAction(game)
		if chosenAction == nil {
			game.GameResult = domain.DRAW

			config.Info.Printf("Game %d ended in a draw", game.ID)

			// TODO: separate method for this logic. Create mapper Game -> GameEntity
			gameEntity, _ := repo.FindGameEntityById(int(game.ID))
			gameEntity.Result = game.GameResult.String()
			repo.SaveGameEntity(gameEntity)
		} else {
			game.ApplyMoveAction(*chosenAction)

			config.Info.Println(chosenAction.String())

			// TODO: separate method for this logic
			moveEntity := repository.ChessGameMoveEntity{GameID: game.ID, FromPosition: chosenAction.FromPosition, ToPosition: chosenAction.ToPosition}
			repo.SaveMoveEntity(moveEntity)
		}

		game.AdvanceToNextTurn()
	}
}
