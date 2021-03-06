package domain

type GameResult int

const (
	DRAW GameResult = iota
	WIN_WHITE
	WIN_BLACK
	UNDETERMINED
)

func (result GameResult) String() string {
	switch {
	case result == DRAW:
		return "DRAW"
	case result == WIN_WHITE:
		return "WIN_WHITE"
	case result == WIN_BLACK:
		return "WIN_BLACK"
	}

	return "UNDETERMINED"
}

func MapStringToGameResult(stringValue string) GameResult {
	if stringValue == "DRAW" {
		return DRAW
	} else if stringValue == "WIN_WHITE" {
		return WIN_WHITE
	} else if stringValue == "WIN_BLACK" {
		return WIN_BLACK
	} else {
		return UNDETERMINED
	}
}
