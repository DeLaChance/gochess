package repository

import "config"

func GenerateChessGameRepository(config *config.Config) *MySqlChessGameRepository {

	return GenerateMySQLChessGameRepository(config)
}
