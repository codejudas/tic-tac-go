package util

import (
	"github.com/gookit/color"
)

type Player struct {
	Ordinal  int
	Symbol   string
	Strategy Strategy
}

func NewPlayer(ord int, strat Strategy) Player {
	symbol := color.Cyan.Render("X")
	if ord != 1 {
		symbol = color.Yellow.Render("O")
	}

	return Player{
		Ordinal:  ord,
		Symbol:   color.Bold.Render(symbol),
		Strategy: strat,
	}
}

func (p Player) NextMove(board Board) string {
	return p.Strategy.NextMove(p, board)
}
