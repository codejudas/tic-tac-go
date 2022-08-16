package strategies

import (
	"fmt"

	"github.com/efossier/tic-tac-go/util"
)

const InteractiveStrategyName = "Interactive"

type InteractiveStrategy struct {
	name string
}

func NewInteractiveStrategy() InteractiveStrategy {
	return InteractiveStrategy{name: InteractiveStrategyName}
}

func (is InteractiveStrategy) Name() string {
	return is.name
}

func (is InteractiveStrategy) IsInteractive() bool {
	return true
}

func (is InteractiveStrategy) NextMove(player util.Player, board util.Board) string {
	fmt.Printf("Player %d's turn: ", player.Ordinal)
	var move string
	fmt.Scan(&move)
	return move
}
