package util

type Player struct {
	Ordinal int
	Symbol  string
}

func newPlayer(ord int) Player {
	symbol := "X"
	if ord != 0 {
		symbol = "O"
	}

	return Player{
		Ordinal: ord,
		Symbol:  symbol,
	}
}

var Players = [...]Player{newPlayer(0), newPlayer(1)}
