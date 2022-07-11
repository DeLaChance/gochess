package main

import (
	"domain"
	"fmt"
)

func main() {

	chessBoard := domain.InitialChessBoard()
	fmt.Println(chessBoard.String())

	chessBoard.ApplyAction(domain.PieceAction{FromPosition: 8, ToPosition: 16})
	fmt.Println(chessBoard.String())
}
