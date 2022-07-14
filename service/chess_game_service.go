package service

import (
	"domain"
	"repository"
)

type ChessGameService struct {
	Repository repository.MySqlChessGameRepository
}

func (service *ChessGameService) FetchById(id int) (*domain.ChessGame, error) {
	chessGameEntity, error := service.Repository.FindGameById(id)

	var chessGame *domain.ChessGame
	if error == nil {
		generatedChessGame := GenerateChessGame(chessGameEntity)

		// Cannot do '&' on the return value of the method.
		chessGame = &generatedChessGame
	} else {
		chessGame = nil
	}

	return chessGame, error
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

	return chessGame
}

func GenerateMoveAction(move repository.ChessGameMoveEntity) domain.PieceMoveAction {
	return domain.PieceMoveAction{FromPosition: move.FromPosition, ToPosition: move.ToPosition}
}
