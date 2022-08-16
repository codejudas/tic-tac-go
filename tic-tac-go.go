package main

import (
	"fmt"

	"github.com/efossier/tic-tac-go/util"
)

const (
	P1 = iota
	P2 = iota
)

func main() {
	fmt.Println("Welcome to Tic-Tac-Go")
	fmt.Println()
	play()
}

func play() {
	var curPlayerIdx = 0
	var board util.Board = util.NewBoard(util.Players[:])

	for {
		board.Print()
		fmt.Println()

		// Check winner
		winner := board.WinningPlayer()
		if winner != nil {
			fmt.Printf("Player %d WINS!\n", winner.Ordinal)
			return
		}

		// Else get next move
		curPlayer := util.Players[curPlayerIdx]
		for {
			fmt.Printf("Player %d's turn: ", curPlayer.Ordinal)
			var move string
			fmt.Scan(&move)

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
		curPlayerIdx = (curPlayerIdx + 1) % len(util.Players)
	}
}
