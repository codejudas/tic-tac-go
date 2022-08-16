package util

import (
	"fmt"
	"strings"
)

const (
	numCols    = 3
	numRows    = 3
	initialVal = "."
)

var columns = [3]string{"A", "B", "C"}
var rows = [3]string{"1", "2", "3"}

type Board struct {
	board   [3][3]string
	players []Player
}

func NewBoard(players []Player) Board {
	var board [3][3]string

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			board[r][c] = initialVal
		}
	}
	return Board{board: board, players: players}
}

func (b Board) Print() {

	// Print header
	fmt.Print(" ")
	for _, c := range columns {
		fmt.Printf(" %s  ", c)
	}
	fmt.Print("\n")

	// Print board
	for r := 0; r < len(rows); r++ {
		fmt.Printf("%s", rows[r])
		for c := 0; c < len(columns); c++ {
			fmt.Printf(" %s ", b.board[r][c])

			if c != len(columns)-1 {
				fmt.Print("|")
			}
		}

		// Print separator
		if r != len(rows)-1 {
			fmt.Printf("\n %s", strings.Repeat("-", len(columns)*3+len(rows)-1))
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
				allEqual = allEqual && b.board[r][c] == p.Symbol
			}
			if allEqual {
				return &p
			}
		}

		// Check cols
		for c := 0; c < numCols; c++ {
			var allEqual = true

			for r := 0; r < numRows; r++ {
				allEqual = allEqual && b.board[r][c] == p.Symbol
			}
			if allEqual {
				return &p
			}
		}

		// Check diagonal top to bottom
		var allEqual = true
		for c, r := 0, 0; c < numCols && r < numRows; c, r = c+1, r+1 {
			allEqual = allEqual && b.board[r][c] == p.Symbol
		}
		if allEqual {
			return &p
		}

		// Check diagonal bottom to top
		allEqual = true
		for c, r := 0, numRows-1; c < numCols && r >= 0; c, r = c+1, r-1 {
			allEqual = allEqual && b.board[r][c] == p.Symbol
		}
		if allEqual {
			return &p
		}
	}
	return nil
}

func (b *Board) Play(p Player, position string) Error {
	// Valiate position
	if len(position) != 2 {
		return InvalidMoveError{move: position}
	}

	col := findInArray(columns[:], string(position[0]))
	row := findInArray(rows[:], string(position[1]))

	if col < 0 || row < 0 {
		return InvalidMoveError{move: position}
	}

	if b.board[row][col] != initialVal {
		return PositionAlreadyOccupiedError{move: position}
	}

	// Update the board
	b.board[row][col] = p.Symbol
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
