package domain

import "fmt"

type BoardCell struct {
	x        int
	y        int
	Contents *ChessPiece
}

func (value BoardCell) IsEmpty() bool {
	return value.Contents == nil
}

func (value BoardCell) GenerateDrawableString() string {

	stringValue := ""
	if value.IsEmpty() {
		stringValue = " "
	} else {
		chessPiece := *value.Contents
		stringValue = chessPiece.String()
	}

	return "|" + fmt.Sprintf("%-1s", stringValue)
}

func (cell *BoardCell) fillCell(piece ChessPiece) {
	cell.Contents = &piece
}

func (cell *BoardCell) emptyCell() {
	cell.Contents = nil
}

func createCell(value ChessPiece, x int, y int) BoardCell {
	return BoardCell{x, y, &value}
}

func createEmptyCell(x int, y int) BoardCell {
	return BoardCell{x, y, nil}
}
