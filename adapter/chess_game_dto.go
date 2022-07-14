package adapter

import "domain"

type ChessGameDto struct {
	ID    uint               `json:"id"`
	Moves []ChessGameMoveDto `json:"moves"`
	State ChessGameStateDto  `json:"state"`
}

type ChessGameMoveDto struct {
	FromPosition uint `json:"fromPosition"`
	ToPosition   uint `json:"toPosition"`
}

type ChessGameStateDto struct {
	Board       []CellDto `json:"board"`
	ActiveColor string    `json:"activeColor"`
	Result      string    `json:"result"`
	WhitePlayer PlayerDto `json:"whitePlayer"`
	BlackPlayer PlayerDto `json:"blackPlayer"`
}

type CellDto struct {
	State string    `json:"state"`
	Piece *PieceDto `json:"piece"`
}

type PieceDto struct {
	Type  string `json:"type"`
	Color string `json:"color"`
}

type PlayerDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GenerateChessGameDto(game domain.ChessGame) ChessGameDto {
	return ChessGameDto{ID: game.ID, Moves: GenerateChessGameMoveDtos(game), State: GenerateState(game)}
}

func GenerateChessGameMoveDtos(game domain.ChessGame) []ChessGameMoveDto {
	var moveDtos []ChessGameMoveDto = make([]ChessGameMoveDto, 0)
	for _, move := range game.Actions {
		moveDto := ChessGameMoveDto{FromPosition: uint(move.FromPosition), ToPosition: uint(move.ToPosition)}
		moveDtos = append(moveDtos, moveDto)
	}

	return moveDtos
}

func GenerateState(game domain.ChessGame) ChessGameStateDto {
	var cellDtos []CellDto

	board := game.Board
	for _, rowOfCells := range board.Cells {
		for _, cell := range rowOfCells {
			var cellDto CellDto
			if cell.IsEmpty() {
				cellDto = CellDto{State: "empty"}
			} else {
				chessPiece := cell.Contents
				cellDto = CellDto{State: "occupied", Piece: &PieceDto{Type: chessPiece.Type.GenerateHumanReadableDescription(), Color: chessPiece.Color.String()}}
			}

			cellDtos = append(cellDtos, cellDto)
		}
	}

	return ChessGameStateDto{
		Board:       cellDtos,
		ActiveColor: game.ActiveColor.String(),
		Result:      game.GameResult.String(),
		WhitePlayer: PlayerDto{ID: game.WhitePlayer.ID, Name: game.WhitePlayer.Name},
		BlackPlayer: PlayerDto{ID: game.BlackPlayer.ID, Name: game.BlackPlayer.Name},
	}
}
