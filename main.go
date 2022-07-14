package main

import (
	"io/ioutil"
	"os"

	"go.uber.org/fx"

	"adapter"
	"config"
	"repository"
	"service"
)

func main() {
	fx.New(
		fx.Invoke(initializeApp),
		fx.Provide(config.GenerateDefaultConfig),
		fx.Provide(repository.GenerateChessGameRepository),
		fx.Provide(service.GenerateChessGameService),
		fx.Invoke(adapter.StartChessController),
	).Run()
}

func initializeApp() {
	config.InitializeLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

/*
func main() {

	domain.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	chessGame := domain.InitialChessGame()
	chessPlayer1 := domain.ChessAIPlayer{domain.WHITE, "Terminator"}
	chessPlayer2 := domain.ChessAIPlayer{domain.BLACK, "C3P0"}
	chessGame.StartGame(chessPlayer1, chessPlayer2)
}*/
