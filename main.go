package main

import (
	"domain"
	"fmt"
)

func main() {
	pieceTypeValues := domain.PieceTypeValues()
	for _, pieceTypeValue := range pieceTypeValues {
		fmt.Println(pieceTypeValue.String())
	}

	chessBoard := domain.InitialChessBoard()
	fmt.Println(chessBoard.String())

}
