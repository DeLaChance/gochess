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
