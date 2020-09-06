package player

import (
	"fmt"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/gamer"
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"math/rand"
	"time"
)

type Player struct {
	Deck     []int
	Hand     []int
	Name     string
	Health   int
	ManaSlot int
	Mana     int
	Gamer    gamer.IGamer
}

//go:generate mockery --name=IPlayer --output=../../../mocks/tcg
type IPlayer interface {
	DrawCard() error
	RefillMana()
	IncreaseManaSlot()
	TakeDamage(damage int)
	EmptiesMana(damage int)
	GetName() string
	GetDeck() []int
	GetHand() []int
	GetHealth() int
	GetGamer() gamer.IGamer
	GetMana() int
	RemoveCardFromHand(card int)
}

// initiate Player with randomized DECK and hand
func NewPlayer(name string, gamer gamer.IGamer) IPlayer {
	randomDeck := getRandomDeck()
	hand := make([]int, 3)
	copy(hand, randomDeck[:3])
	return &Player{randomDeck[3:], hand, name, InitialHealth, InitialMana, InitialMana, gamer}
}

const (
	InitialHealth     = 30
	InitialMana       = 0
	ManaSlotIncrement = 1
	MaxMana           = 10
	MaxHandSize       = 5
)

var DECK = []int{0, 0, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 6, 6, 7, 8}

func getRandomDeck() []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(DECK), func(i, j int) {
		DECK[i], DECK[j] = DECK[j], DECK[i]
	})
	return DECK
}
func (p *Player) DrawCard() error {
	if len(p.Deck) == 0 {
		return fmt.Errorf("Deck is over")
	}
	card := p.Deck[0]
	p.Deck = p.Deck[1:]
	// hand size 5 means discard that card
	if len(p.Hand) != MaxHandSize {
		utils.PrintMessage(utils.Draw, " %s draw %v ", p.Name, card)
		p.Hand = append(p.Hand, card)
	} else {
		utils.PrintMessage(utils.Overload, "Overload!!! %s  No Card could be drawn ", p.Name)
	}
	return nil
}
func (p *Player) RefillMana() {
	p.Mana = p.ManaSlot
}
func (p *Player) IncreaseManaSlot() {
	if p.ManaSlot != MaxMana {
		p.ManaSlot = p.ManaSlot + ManaSlotIncrement
	}
}
func (p *Player) RemoveCardFromHand(card int) {
	p.Hand = utils.FindAndRemove(p.Hand, card)
}

func (p *Player) EmptiesMana(damage int) {
	p.Mana = p.Mana - damage
}

func (p *Player) TakeDamage(damage int) {
	p.Health = p.Health - damage
}
func (p *Player) GetName() string {
	return p.Name
}
func (p *Player) GetDeck() []int {
	return p.Deck
}
func (p *Player) GetHand() []int {
	return p.Hand
}
func (p *Player) GetHealth() int {
	return p.Health
}
func (p *Player) GetGamer() gamer.IGamer {
	return p.Gamer
}

func (p *Player) GetMana() int {
	return p.Mana
}
