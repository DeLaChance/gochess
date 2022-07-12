package domain

type GameResult int

const (
	DRAW GameResult = iota
	WIN_WHITE
	WIN_BLACK
	UNDETERMINED
)
