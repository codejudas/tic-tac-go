package util

type Strategy interface {
	Name() string
	NextMove(p Player, b Board) (position string)
	IsInteractive() bool
}
