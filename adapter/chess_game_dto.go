package adapter

import "domain"

type ChessGameDto struct {
	ID uint `json:"id"`
}

func GenerateChessGameDto(game *domain.ChessGame) ChessGameDto {
	return ChessGameDto{ID: game.ID}
}
