package player

import (
	mocks "github.com/emreodabas/kata-tcg-golang/mocks/tcg"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/gamer"
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestPlayer_DrawCard(t *testing.T) {

	var value = getRandomInt()
	player := &Player{
		Deck: []int{value},
		Hand: []int{},
	}
	err := player.DrawCard()
	if err != nil {
		t.Fatal(err)
	}
	if len(player.Deck) > 0 {
		t.Errorf("Expected to 1, found %v", len(player.Deck))
	}
	if len(player.Hand) != 1 {
		t.Errorf("Expected to 0, found %v", len(player.Hand))
	}
	if player.Hand[0] != value {
		t.Errorf("Expected to %v, found %v", value, player.Deck[0])
	}
}

func TestPlayer_NoDrawCard(t *testing.T) {
	var value = getRandomInt()
	player := &Player{
		// size MaxHandSize
		Deck: []int{value},
		Hand: []int{0, 0, 0, 0, 0},
	}
	err := player.DrawCard()

	if err != nil {
		t.Errorf("Expected no error,found %s", err)
	}
	// card need to be discarded
	if len(player.Deck) > 0 {
		t.Errorf("Expected to 1, found %v", len(player.Deck))
	}
	if len(player.Hand) != MaxHandSize {
		t.Errorf("Expected to %v, found %v", MaxHandSize, len(player.Hand))
	}
	if utils.ContainsValue(player.Hand, value) {
		t.Errorf("Expected to %v not found in %v", value, player.Hand)
	}
}

func TestPlayer_DrawCardWithEmptyDeck(t *testing.T) {
	player := &Player{
		Deck: []int{},
		Hand: []int{},
	}
	err := player.DrawCard()
	if err == nil {
		t.Errorf("Expected error,found %s", err)
	}
	if len(player.Deck) != 0 {
		t.Errorf("Expected to 0, found %v", len(player.Deck))
	}
	if len(player.Hand) != 0 {
		t.Errorf("Expected to 0, found %v", len(player.Hand))
	}
}

func TestPlayer_TakeDamage(t *testing.T) {
	var value = getRandomInt() % InitialHealth
	player := &Player{
		Health: InitialHealth,
	}
	t.Log("Random Damage", value)
	player.TakeDamage(value)

	if player.Health != InitialHealth-value {
		t.Errorf("Expected to %v, found %v", InitialHealth-value, player.Health)
	}
}

func TestPlayer_RefillMana(t *testing.T) {
	var value = 1
	player := &Player{
		ManaSlot: value,
	}
	player.RefillMana()

	if player.ManaSlot != player.Mana {
		t.Errorf("Expected to %v, found %v", player.ManaSlot, player.Mana)
	}
}

func TestPlayer_RefillManaMaxed(t *testing.T) {

	player := &Player{
		ManaSlot: MaxMana,
	}
	player.RefillMana()

	if player.ManaSlot != MaxMana {
		t.Errorf("Expected to %v, found %v", MaxMana, player.ManaSlot)
	}
}

func TestPlayer_GetRandomDeck(t *testing.T) {
	player := &Player{
		Deck: getRandomDeck(),
	}
	if len(player.Deck) != len(DECK) {
		t.Errorf("Expected to %v deck size found %v", len(DECK), len(player.Deck))
	}
}

func TestNewPlayer(t *testing.T) {

	player := NewPlayer("test", new(gamer.AutoGamer))

	if player.GetName() != "test" {
		t.Errorf("Expected to test found %v", player.GetName())
	}
	if player.GetHealth() != InitialHealth {
		t.Errorf("Expected to %v health found %v", InitialHealth, player.GetHealth())
	}
	if player.GetMana() != InitialMana {
		t.Errorf("Expected to %v mana found %v", InitialMana, player.GetMana())
	}
	if len(player.GetDeck()) != 17 {
		t.Errorf("Expected to 20 found %v", len(player.GetDeck()))
	}
	if len(player.GetHand()) != 3 {
		t.Errorf("Expected to 3 found %v", len(player.GetHand()))
	}
}

func TestPlayer_EmptiesMana(t *testing.T) {
	player := &Player{
		Mana: 5,
	}
	player.EmptiesMana(5)

	if player.GetMana() != 0 {
		t.Errorf("Expected to 0 found %v", player.GetMana())
	}
}

func TestPlayer_GetDeck(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}
	player := &Player{
		Deck: ints,
	}
	deck := player.GetDeck()

	if len(deck) != 5 {
		t.Errorf("Expected to 5 found %v", len(deck))
	}

	if !reflect.DeepEqual(deck, ints) {
		t.Errorf("Expected to true found false")
	}

}

func TestPlayer_GetGamer(t *testing.T) {
	gamer := new(mocks.IGamer)
	player := &Player{
		Gamer: gamer,
	}
	getGamer := player.GetGamer()
	if gamer != getGamer {
		t.Errorf("Expected to both are same")
	}
}

func TestPlayer_IncreaseManaSlot(t *testing.T) {
	player := &Player{
		ManaSlot: 3,
	}
	player.IncreaseManaSlot()

	if player.ManaSlot != 3+ManaSlotIncrement {
		t.Errorf("Expected to %v found %v", 3+ManaSlotIncrement, player.ManaSlot)
	}
}

func TestPlayer_RemoveCardFromHand(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}
	player := &Player{
		Hand: ints,
	}

	player.RemoveCardFromHand(2)
	if len(player.Hand) != 4 {
		t.Errorf("Expected to 4 found %v", len(player.Hand))
	}
	if utils.ContainsValue(player.Hand, 2) {
		t.Errorf("Expected to false found true")
	}
}

func getRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
