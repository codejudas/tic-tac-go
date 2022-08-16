package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/efossier/tic-tac-go/util"
	"github.com/efossier/tic-tac-go/util/strategies"
)

func main() {

	p1StratName := flag.String("p1", "interactive", "Strategy to use for Player 1")
	p2StratName := flag.String("p2", "interactive", "Strategy to use for Player 2")
	flag.Parse()

	p1Strat, err := chooseStrategy(*p1StratName)
	if err != nil {
		fmt.Println(err.Message())
		return
	}

	p2Strat, err := chooseStrategy(*p2StratName)
	if err != nil {
		fmt.Println(err.Message())
		return
	}

	fmt.Println("Welcome to Tic-Tac-Go")
	fmt.Println()

	player1 := util.NewPlayer(1, p1Strat)
	player2 := util.NewPlayer(2, p2Strat)

	var players = []util.Player{player1, player2}

	play(players)
}

func chooseStrategy(name string) (util.Strategy, util.Error) {
	if name == strings.ToLower(strategies.InteractiveStrategyName) {
		return strategies.NewInteractiveStrategy(), nil
	} else if name == strings.ToLower(strategies.RandomStrategyName) {
		return strategies.NewRandomStrategy(), nil
	}

	return nil, util.NewUnknownStrategyError(name)
}

func play(players []util.Player) {
	var curPlayerIdx = 0
	var board util.Board = util.NewBoard(players)

	for {
		board.Print()
		fmt.Println()

		// Check winner
		winner := board.WinningPlayer()
		if winner != nil {
			fmt.Printf("Player %d WINS!\n", winner.Ordinal)
			return
		}

		// Check draw
		if board.IsDraw() {
			fmt.Println("Game Ends in DRAW")
			return
		}

		// Else get next move
		curPlayer := players[curPlayerIdx]
		for {
			// Don't spam the terminal if playing between two non-interactive players or there is a bug causing infinite bad choices
			if !curPlayer.Strategy.IsInteractive() {
				time.Sleep(1 * time.Second)
			}

			move := curPlayer.NextMove(board)

			// Try to update the board
			res := board.Play(curPlayer, move)
			if res != nil {
				fmt.Println(res.Message())
			} else {
				break
			}
		}

		fmt.Println()

		// Next player
		curPlayerIdx = (curPlayerIdx + 1) % len(players)
	}
}
