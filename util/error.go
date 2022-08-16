package util

import "fmt"

type Error interface {
	Message() string
}

type InvalidMoveError struct {
	move string
}

func (ime InvalidMoveError) Message() string {
	return fmt.Sprintf("Error: Invalid move '%s'", ime.move)
}

type PositionAlreadyOccupiedError struct {
	move string
}

func (paoe PositionAlreadyOccupiedError) Message() string {
	return fmt.Sprintf("Error: Position '%s' has already been played", paoe.move)
}
