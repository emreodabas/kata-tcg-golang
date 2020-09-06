package gamer

//go:generate mockery --name=IGamer --output=../../../mocks/tcg --keeptree
type IGamer interface {
	GetCardDecision(hands []int, mana int) int
	GetGamerType() GamerType
}
type GamerType int

const (
	Auto GamerType = iota
	Console
)
