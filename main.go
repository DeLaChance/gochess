package main

import (
	"domain"
	"fmt"
)

func main() {

	chessBoard := domain.InitialChessBoard()
	fmt.Println(chessBoard.String())

}
