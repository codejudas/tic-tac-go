package util

import (
	"fmt"
	"strings"
)

const (
	numCols    = 3
	numRows    = 3
	InitialVal = " "
)

var RowLabels = [3]string{"A", "B", "C"}
var ColumnLabels = [3]string{"1", "2", "3"}

type Board struct {
	Board   [3][3]string
	players []Player
}

func NewBoard(players []Player) Board {
	var board [3][3]string

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			board[r][c] = InitialVal
		}
	}
	return Board{Board: board, players: players}
}

func (b Board) Print() {
	// Print header
	fmt.Print(" ")
	for _, c := range ColumnLabels {
		fmt.Printf(" %s  ", c)
	}
	fmt.Print("\n")

	// Print board
	for r := 0; r < len(RowLabels); r++ {
		fmt.Printf("%s", RowLabels[r])
		for c := 0; c < len(ColumnLabels); c++ {
			fmt.Printf(" %s ", b.Board[r][c])

			if c != len(ColumnLabels)-1 {
				fmt.Print("|")
			}
		}

		// Print separator
		if r != len(RowLabels)-1 {
			fmt.Printf("\n %s", strings.Repeat("-", len(ColumnLabels)*3+len(RowLabels)-1))
		}

		fmt.Print("\n")
	}
}

func (b Board) WinningPlayer() *Player {
	for _, p := range b.players {
		// Check rows
		for r := 0; r < numRows; r++ {
			var allEqual = true

			for c := 0; c < numCols; c++ {
				allEqual = allEqual && b.Board[r][c] == p.Symbol
			}
			if allEqual {
				return &p
			}
		}

		// Check cols
		for c := 0; c < numCols; c++ {
			var allEqual = true

			for r := 0; r < numRows; r++ {
				allEqual = allEqual && b.Board[r][c] == p.Symbol
			}
			if allEqual {
				return &p
			}
		}

		// Check diagonal top to bottom
		var allEqual = true
		for c, r := 0, 0; c < numCols && r < numRows; c, r = c+1, r+1 {
			allEqual = allEqual && b.Board[r][c] == p.Symbol
		}
		if allEqual {
			return &p
		}

		// Check diagonal bottom to top
		allEqual = true
		for c, r := 0, numRows-1; c < numCols && r >= 0; c, r = c+1, r-1 {
			allEqual = allEqual && b.Board[r][c] == p.Symbol
		}
		if allEqual {
			return &p
		}
	}
	return nil
}

func (b Board) IsDraw() bool {
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if b.Board[r][c] == InitialVal {
				return false
			}
		}
	}
	return true
}

func (b *Board) Play(p Player, position string) Error {
	// Valiate position
	if len(position) != 2 {
		return InvalidMoveError{move: position}
	}

	row := findInArray(RowLabels[:], string(position[0]))
	col := findInArray(ColumnLabels[:], string(position[1]))

	if col < 0 || row < 0 {
		return InvalidMoveError{move: position}
	}

	if b.Board[row][col] != InitialVal {
		return PositionAlreadyOccupiedError{move: position}
	}

	// Update the board
	b.Board[row][col] = p.Symbol
	return nil
}

func findInArray(arr []string, target string) int {
	for i, e := range arr {
		if e == target {
			return i
		}
	}
	return -1
}
