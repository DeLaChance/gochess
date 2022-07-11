package domain

type ChessBoard struct {
	Cells [8][8]BoardCell
}

func InitialChessBoard() ChessBoard {

	chessBoard := ChessBoard{}

	return chessBoard
}

func (value ChessBoard) String() string {
	output := ""
	for _, rowOfCells := range value.Cells {
		for _, cell := range rowOfCells {
			output += cell.String()
		}
		output += "\n"
	}

	return output
}
