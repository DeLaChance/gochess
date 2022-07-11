package domain

type BoardCell struct {
	Contents *ChessPiece
}

func (value BoardCell) isEmpty() bool {
	return value.Contents != nil
}
