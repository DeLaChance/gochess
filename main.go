package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"go.uber.org/fx"

	"adapter"
	"config"
	"repository"
)

func main() {
	fx.New(
		fx.Provide(config.GenerateDefaultConfig),
		fx.Invoke(initializeApp),
		fx.Provide(http.NewServeMux),
		fx.Invoke(adapter.New),
		fx.Invoke(repository.GenerateChessGameRepository),
		fx.Invoke(registerHooks),
	).Run()
}

func initializeApp() {
	config.InitializeLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func registerHooks(lifecycle fx.Lifecycle, mux *http.ServeMux, config *config.Config) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				httpUrl := fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort)
				go http.ListenAndServe(httpUrl, mux)
				return nil
			},
		},
	)
}

/*
func main() {

	domain.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	chessGame := domain.InitialChessGame()
	chessPlayer1 := domain.ChessAIPlayer{domain.WHITE, "Terminator"}
	chessPlayer2 := domain.ChessAIPlayer{domain.BLACK, "C3P0"}
	chessGame.StartGame(chessPlayer1, chessPlayer2)
}*/
