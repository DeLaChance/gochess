package domain

type ChessBoard struct {
	Cells [8][8]BoardCell
}

func InitialChessBoard() ChessBoard {

	chessBoard := ChessBoard{}

	chessBoard.fillBlackMainRow()
	chessBoard.fillRowWithPawns(1, BLACK)

	chessBoard.fillWhiteMainRow()
	chessBoard.fillRowWithPawns(6, WHITE)

	return chessBoard
}

func (value ChessBoard) String() string {
	output := ""
	for _, rowOfCells := range value.Cells {

		output += "-----------------\n"

		for _, cell := range rowOfCells {
			output += cell.String()
		}

		output += "|\n"

	}

	output += "-----------------\n"

	return output
}

// Private methods
func (chessBoard *ChessBoard) fillRowWithPawns(rowIndex int, pieceColor PieceColor) {
	for index := 0; index < 8; index += 1 {
		chessBoard.Cells[rowIndex][index] = createCell(ChessPiece{PAWN, pieceColor})
	}
}

func (chessBoard *ChessBoard) fillWhiteMainRow() {

	chessBoard.Cells[7][0] = createCell(ChessPiece{ROOK, WHITE})
	chessBoard.Cells[7][1] = createCell(ChessPiece{KNIGHT, WHITE})
	chessBoard.Cells[7][2] = createCell(ChessPiece{BISHOP, WHITE})
	chessBoard.Cells[7][3] = createCell(ChessPiece{QUEEN, WHITE})
	chessBoard.Cells[7][4] = createCell(ChessPiece{KING, WHITE})
	chessBoard.Cells[7][5] = createCell(ChessPiece{BISHOP, WHITE})
	chessBoard.Cells[7][6] = createCell(ChessPiece{KNIGHT, WHITE})
	chessBoard.Cells[7][7] = createCell(ChessPiece{ROOK, WHITE})
}

func (chessBoard *ChessBoard) fillBlackMainRow() {

	chessBoard.Cells[0][0] = createCell(ChessPiece{ROOK, BLACK})
	chessBoard.Cells[0][1] = createCell(ChessPiece{KNIGHT, BLACK})
	chessBoard.Cells[0][2] = createCell(ChessPiece{BISHOP, BLACK})
	chessBoard.Cells[0][3] = createCell(ChessPiece{QUEEN, BLACK})
	chessBoard.Cells[0][4] = createCell(ChessPiece{KING, BLACK})
	chessBoard.Cells[0][5] = createCell(ChessPiece{BISHOP, BLACK})
	chessBoard.Cells[0][6] = createCell(ChessPiece{KNIGHT, BLACK})
	chessBoard.Cells[0][7] = createCell(ChessPiece{ROOK, BLACK})
}
