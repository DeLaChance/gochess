package main

import (
	"domain"
	"io/ioutil"
	"os"
)

func main() {

	domain.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	chessGame := domain.InitialChessGame()
	chessPlayer1 := domain.ChessAIPlayer{domain.WHITE, "Terminator"}
	chessPlayer2 := domain.ChessAIPlayer{domain.BLACK, "C3P0"}
	chessGame.StartGame(chessPlayer1, chessPlayer2)
}
