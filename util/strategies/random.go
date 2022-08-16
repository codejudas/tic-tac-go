package strategies

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/efossier/tic-tac-go/util"
)

const RandomStrategyName = "Random"

type RandomStrategy struct {
	name string
}

func NewRandomStrategy() RandomStrategy {
	rand.Seed(time.Now().UnixNano())
	return RandomStrategy{name: RandomStrategyName}
}

func (rs RandomStrategy) Name() string {
	return rs.name
}

func (rs RandomStrategy) IsInteractive() bool {
	return false
}

func (rs RandomStrategy) NextMove(player util.Player, board util.Board) string {
	// Find all open spots on the board
	var moves []string

	for r := 0; r < len(board.Board); r++ {
		for c := 0; c < len(board.Board[r]); c++ {
			if board.Board[r][c] == util.InitialVal {
				moves = append(moves, util.RowLabels[r]+util.ColumnLabels[c])
			}
		}
	}

	randomMove := moves[rand.Intn(len(moves))]

	fmt.Printf("Player %d randomly picks: %s\n", player.Ordinal, randomMove)
	return randomMove
}
