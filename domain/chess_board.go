package domain

const BOARD_DIMENSION = 8

type ChessBoard struct {
	Cells [BOARD_DIMENSION][BOARD_DIMENSION]BoardCell
}

func (board ChessBoard) GeneratePossibleActions(activeCell BoardCell) []PieceMoveAction {

	var actions []PieceMoveAction
	piece := *activeCell.Contents
	if piece.Type == PAWN {
		actions = board.generatePawnPossibleActions(activeCell)
	} else {
		actions = make([]PieceMoveAction, 0) // TODO: implement other pieces
	}

	return actions
}

func (chessBoard *ChessBoard) ApplyMoveAction(action PieceMoveAction) *ChessPiece {
	oldX, oldY := calculateTwoDimensionalPosition(action.FromPosition)
	newX, newY := calculateTwoDimensionalPosition(action.ToPosition)

	oldCell := chessBoard.fetchCell(oldX, oldY)
	newCell := chessBoard.fetchCell(newX, newY)

	movedPiece := oldCell.Contents
	capturedPiece := newCell.Contents

	chessBoard.Cells[newY][newX] = BoardCell{newX, newY, movedPiece}
	chessBoard.Cells[oldY][oldX] = BoardCell{oldX, oldY, nil}

	return capturedPiece
}

func (chessBoard *ChessBoard) UnapplyMoveAction(action PieceMoveAction, capturedPiece *ChessPiece) {
	oldX, oldY := calculateTwoDimensionalPosition(action.FromPosition)
	newX, newY := calculateTwoDimensionalPosition(action.ToPosition)

	newCell := chessBoard.fetchCell(newX, newY)
	movedPiece := newCell.Contents

	chessBoard.Cells[newY][newX] = BoardCell{newX, newY, capturedPiece}
	chessBoard.Cells[oldY][oldX] = BoardCell{oldX, oldY, movedPiece}
}

func (chessBoard *ChessBoard) CalculateScore(color PieceColor) int {
	playerCells := chessBoard.FindCellsByColor(color)

	score := 0
	for _, cell := range playerCells {
		score += cell.Contents.Type.PieceValue()
	}

	oponentPlayerCells := chessBoard.FindCellsByColor(color)
	for _, cell := range oponentPlayerCells {
		score -= cell.Contents.Type.PieceValue()
	}

	return score
}

func (board *ChessBoard) FindCellsByColor(color PieceColor) []BoardCell {
	var matchingCells []BoardCell
	for _, rowOfCells := range board.Cells {
		for _, cell := range rowOfCells {
			if !cell.IsEmpty() {
				if cell.Contents.Color == color {
					matchingCells = append(matchingCells, cell)
				}
			}
		}
	}

	return matchingCells
}

func (value ChessBoard) String() string {
	output := ""
	for _, rowOfCells := range value.Cells {

		output += "-----------------\n"

		for _, cell := range rowOfCells {
			output += cell.GenerateDrawableString()
		}

		output += "|\n"

	}

	output += "-----------------\n"

	return output
}

// Static methods
func InitialChessBoard() ChessBoard {

	chessBoard := ChessBoard{}

	chessBoard.fillBlackMainRow()
	chessBoard.fillRowWithPawns(1, BLACK)

	chessBoard.fillWhiteMainRow()
	chessBoard.fillRowWithPawns(6, WHITE)

	chessBoard.fillEmptyCells(2, 5)

	return chessBoard
}

// Private methods
func (chessBoard *ChessBoard) fillRowWithPawns(rowIndex int, pieceColor PieceColor) {
	for index := 0; index < BOARD_DIMENSION; index += 1 {
		chessBoard.Cells[rowIndex][index] = createCell(ChessPiece{PAWN, pieceColor}, index, rowIndex)
	}
}

func calculateTwoDimensionalPosition(position int) (int, int) {
	return position % BOARD_DIMENSION, position / BOARD_DIMENSION
}

func (chessBoard *ChessBoard) fillWhiteMainRow() {

	chessBoard.Cells[7][0] = createCell(ChessPiece{ROOK, WHITE}, 0, 7)
	chessBoard.Cells[7][1] = createCell(ChessPiece{KNIGHT, WHITE}, 1, 7)
	chessBoard.Cells[7][2] = createCell(ChessPiece{BISHOP, WHITE}, 2, 7)
	chessBoard.Cells[7][3] = createCell(ChessPiece{QUEEN, WHITE}, 3, 7)
	chessBoard.Cells[7][4] = createCell(ChessPiece{KING, WHITE}, 4, 7)
	chessBoard.Cells[7][5] = createCell(ChessPiece{BISHOP, WHITE}, 5, 7)
	chessBoard.Cells[7][6] = createCell(ChessPiece{KNIGHT, WHITE}, 6, 7)
	chessBoard.Cells[7][7] = createCell(ChessPiece{ROOK, WHITE}, 7, 7)
}

func (chessBoard *ChessBoard) fillBlackMainRow() {

	chessBoard.Cells[0][0] = createCell(ChessPiece{ROOK, BLACK}, 0, 0)
	chessBoard.Cells[0][1] = createCell(ChessPiece{KNIGHT, BLACK}, 1, 0)
	chessBoard.Cells[0][2] = createCell(ChessPiece{BISHOP, BLACK}, 2, 0)
	chessBoard.Cells[0][3] = createCell(ChessPiece{QUEEN, BLACK}, 3, 0)
	chessBoard.Cells[0][4] = createCell(ChessPiece{KING, BLACK}, 4, 0)
	chessBoard.Cells[0][5] = createCell(ChessPiece{BISHOP, BLACK}, 5, 0)
	chessBoard.Cells[0][6] = createCell(ChessPiece{KNIGHT, BLACK}, 6, 0)
	chessBoard.Cells[0][7] = createCell(ChessPiece{ROOK, BLACK}, 7, 0)
}

func (chessBoard *ChessBoard) fillEmptyCells(startRow int, endRow int) {
	for y := startRow; y < endRow; y += 1 {
		for x := 0; x < BOARD_DIMENSION; x += 1 {
			chessBoard.Cells[y][x] = createEmptyCell(x, y)
		}
	}

}

func (board ChessBoard) generatePawnPossibleActions(activeCell BoardCell) []PieceMoveAction {

	var beginRow int
	var direction int

	piece := *activeCell.Contents
	if piece.Color == WHITE {
		beginRow = 6
		direction = -1
	} else {
		beginRow = 1
		direction = 1
	}

	isInitialMove := activeCell.y == beginRow

	x := activeCell.x
	y := activeCell.y

	var actions []PieceMoveAction = make([]PieceMoveAction, 0)

	// A pawn can move forward 2 if it's the initial move
	if isInitialMove && board.canMoveInEmptyCell(x, y+direction*2) {
		actions = append(actions, GeneratePieceMoveAction(x, y, x, y+direction*2))
	}

	// A pawn can always forward 1 into an empty cell
	if board.canMoveInEmptyCell(x, y+direction*1) {
		actions = append(actions, GeneratePieceMoveAction(x, y, x, y+direction*1))
	}

	// A pawn can capture a diagonal enemy piece
	if board.hasEnemyPiece(x+1, y+direction*1, piece.Color) {
		actions = append(actions, GeneratePieceMoveAction(x, y, x+1, y+direction*1))
	}

	if board.hasEnemyPiece(x-1, y+direction*1, piece.Color) {
		actions = append(actions, GeneratePieceMoveAction(x, y, x-1, y+direction*1))
	}

	return actions
}

func (board *ChessBoard) hasEnemyPiece(x int, y int, ownColor PieceColor) bool {

	var hasEnemyPiece bool
	if x >= 0 && x < BOARD_DIMENSION && y >= 0 && y < BOARD_DIMENSION {
		otherPiece := board.Cells[y][x].Contents
		hasEnemyPiece = otherPiece != nil && otherPiece.Color != ownColor
	} else {
		hasEnemyPiece = false
	}

	return hasEnemyPiece
}

func (board *ChessBoard) canMoveInEmptyCell(x int, y int) bool {
	var canMoveInEmptyCell bool
	if x >= 0 && x < BOARD_DIMENSION && y >= 0 && y < BOARD_DIMENSION {
		cell := board.fetchCell(x, y)
		cellContents := cell.Contents
		canMoveInEmptyCell = cellContents == nil
	} else {
		canMoveInEmptyCell = false
	}

	return canMoveInEmptyCell
}

func (board *ChessBoard) fetchCell(x int, y int) *BoardCell {
	return &board.Cells[y][x]
}
