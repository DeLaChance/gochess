package adapter

import "domain"

type ChessGameDto struct {
	ID    uint               `json:"id"`
	Moves []ChessGameMoveDto `json:"moves"`
}

type ChessGameMoveDto struct {
	FromPosition uint `json:"fromPosition"`
	ToPosition   uint `json:"toPosition"`
}

func GenerateChessGameDto(game *domain.ChessGame) ChessGameDto {
	return ChessGameDto{ID: game.ID, Moves: GenerateChessGameMoveDtos(game)}
}

func GenerateChessGameMoveDtos(game *domain.ChessGame) []ChessGameMoveDto {
	var moveDtos []ChessGameMoveDto = make([]ChessGameMoveDto, 0)
	for _, move := range game.Actions {
		moveDto := ChessGameMoveDto{FromPosition: uint(move.FromPosition), ToPosition: uint(move.ToPosition)}
		moveDtos = append(moveDtos, moveDto)
	}

	return moveDtos
}
