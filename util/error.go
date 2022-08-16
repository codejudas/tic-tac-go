package util

import (
	"fmt"

	"github.com/gookit/color"
)

var errorPrefix = color.Bold.Render(color.Red.Render("Error:"))

type Error interface {
	Message() string
}

type InvalidMoveError struct {
	move string
}

func (ime InvalidMoveError) Message() string {
	return fmt.Sprintf("%s Invalid move '%s'", errorPrefix, ime.move)
}

type PositionAlreadyOccupiedError struct {
	move string
}

func (paoe PositionAlreadyOccupiedError) Message() string {
	return fmt.Sprintf("%s Position '%s' has already been played", errorPrefix, paoe.move)
}

type UnknownStrategyError struct {
	stratName string
}

func NewUnknownStrategyError(stratName string) UnknownStrategyError {
	return UnknownStrategyError{stratName: stratName}
}

func (use UnknownStrategyError) Message() string {
	return fmt.Sprintf("%s Strategy '%s' is invalid", errorPrefix, use.stratName)
}
