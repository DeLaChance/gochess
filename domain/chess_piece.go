package domain

import "fmt"

type ChessPiece struct {
	Type  PieceType
	Color PieceColor
}

func (value ChessPiece) String() string {
	return fmt.Sprintf("%s %s", value.Color.String(), value.Type.String())
}
