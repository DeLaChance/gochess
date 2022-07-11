package domain

import "fmt"

type BoardCell struct {
	Contents *ChessPiece
}

func (value BoardCell) IsEmpty() bool {
	return value.Contents == nil
}

func (value BoardCell) String() string {

	stringValue := ""
	if value.IsEmpty() {
		stringValue = " "
	} else {
		chessPiece := *value.Contents
		stringValue = chessPiece.String()
	}

	return "|" + fmt.Sprintf("%-12s", stringValue) + "|"
}
