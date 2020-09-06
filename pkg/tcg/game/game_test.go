package game

import (
	"bytes"
	"fmt"
	mocks "github.com/emreodabas/kata-tcg-golang/mocks/tcg"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/gamer"
	player2 "github.com/emreodabas/kata-tcg-golang/pkg/tcg/player"
	"github.com/fatih/color"
	"strings"
	"testing"
)

func TestTurnDeckIsOverNoMove(t *testing.T) {

	mckPlayer := new(mocks.IPlayer)
	mckOpp := new(mocks.IPlayer)
	mckGamer := new(mocks.IGamer)
	game := NewGame(mckPlayer, mckPlayer, mckOpp)
	err := fmt.Errorf("Deck is over")
	mckPlayer.On("IncreaseManaSlot")
	mckPlayer.On("RefillMana")
	mckPlayer.On("DrawCard").Return(err)
	mckPlayer.On("GetHand").Return([]int{1, 2})
	mckPlayer.On("GetMana").Return(0)
	mckPlayer.On("GetHealth").Return(1)
	mckPlayer.On("GetGamer").Return(mckGamer)
	mckGamer.On("GetCardDecision", []int{1, 2}, 1).Return(1)
	mckGamer.On("GetGamerType").Return(gamer.Auto)
	mckPlayer.On("RemoveCardFromHand", 1)
	mckPlayer.On("EmptiesMana", 1)
	mckPlayer.On("GetName").Return("MyName")
	mckOpp.On("TakeDamage", 1)
	mckOpp.On("GetName").Return("Opponent")
	mckOpp.On("GetHealth").Return(3)
	game.Turn()

}

func TestMoveActions(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	mckOpp := new(mocks.IPlayer)
	mckGamer := new(mocks.IGamer)
	game := NewGame(mckPlayer, mckPlayer, mckOpp)
	rb := new(bytes.Buffer)
	color.Output = rb

	mckPlayer.On("GetHand").Return([]int{1, 2})
	mckPlayer.On("GetMana").Return(1)
	mckPlayer.On("GetHealth").Return(1)
	mckPlayer.On("GetGamer").Return(mckGamer)
	mckGamer.On("GetCardDecision", []int{1, 2}, 1).Return(1)
	mckGamer.On("GetGamerType").Return(gamer.Auto)
	mckPlayer.On("RemoveCardFromHand", 1)
	mckPlayer.On("EmptiesMana", 1)
	mckPlayer.On("GetName").Return("MyName")
	mckOpp.On("TakeDamage", 1)
	mckOpp.On("GetName").Return("Opponent")
	mckOpp.On("GetHealth").Return(0)
	game.MoveActions()

	readString, _ := rb.ReadString(0)
	if !strings.Contains(readString, "Health:::: 0") {
		t.Errorf("Expected to output contains attack text %v", readString)
	}
}
func TestIterateTurns(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	mckOpp := new(mocks.IPlayer)
	mckGamer := new(mocks.IGamer)
	game := NewGame(mckPlayer, mckPlayer, mckOpp)
	mckPlayer.On("GetHand").Return([]int{1, 2})
	mckPlayer.On("GetMana").Return(1)
	mckPlayer.On("GetHealth").Return(0)
	mckPlayer.On("GetGamer").Return(mckGamer)
	mckGamer.On("GetCardDecision", []int{1, 2}, 1).Return(1)
	mckGamer.On("GetGamerType").Return(gamer.Auto)
	mckPlayer.On("RemoveCardFromHand", 1)
	mckPlayer.On("EmptiesMana", 1)
	mckPlayer.On("GetName").Return("MyName")
	mckOpp.On("TakeDamage", 1)
	mckOpp.On("GetName").Return("Opponent")
	mckOpp.On("GetHealth").Return(0)
	game.IterateTurns()
}

//integration
func TestIterateTurnsIT(t *testing.T) {
	player := &player2.Player{
		Health:   10,
		Deck:     []int{3},
		Hand:     []int{3, 5},
		Mana:     8,
		ManaSlot: 8,
		Gamer:    new(gamer.AutoGamer),
	}
	cPlayer := &player2.Player{
		Health:   6,
		Deck:     []int{0},
		Hand:     []int{},
		Mana:     0,
		ManaSlot: 8,
		Gamer:    new(gamer.AutoGamer),
	}
	game := Game{
		ActivePlayer:  cPlayer,
		ConsolePlayer: cPlayer,
		Player:        player,
		round:         0,
	}
	game.IterateTurns()

	if game.round != 2 {
		t.Errorf("Expected %v found %v", 2, game.round)
	}
	if game.ConsolePlayer.GetHealth() >= 0 {
		t.Errorf("Expected %v found %v", -2, game.ConsolePlayer.GetHealth())
	}
	if !game.GameEnded() {
		t.Errorf("Expected true found false")
	}
}

func TestChooseGameStarter(t *testing.T) {
	player := player2.NewPlayer("P", new(gamer.AutoGamer))
	cPlayer := player2.NewPlayer("C", new(gamer.ConsoleGamer))
	game := &Game{
		Player:        player,
		ConsolePlayer: cPlayer,
	}
	game.ChooseGameStarter(int64(0))
	if game.ActivePlayer != game.ConsolePlayer {
		t.Errorf("Expected active player is Player, found ConsolePlayer")
	}
	game.ChooseGameStarter(int64(1))
	if game.ActivePlayer != game.Player {
		t.Errorf("Expected active ConsolePlayer is chosed found Player ")
	}
}
func TestGameNotEnded(t *testing.T) {

	player := player2.NewPlayer(" ", nil)

	game := Game{
		ConsolePlayer: player,
		Player:        player,
	}
	ended := game.GameEnded()

	if ended {
		t.Errorf("Expected false found true")
	}
}

func TestGameEnded(t *testing.T) {

	player := player2.NewPlayer("", nil)
	cPlayer := &player2.Player{
		Health: 0,
	}
	game := Game{
		ConsolePlayer: cPlayer,
		Player:        player,
	}
	ended := game.GameEnded()

	if !ended {
		t.Errorf("Expected true found false")
	}
}
func TestGetOpponent(t *testing.T) {

	player := player2.NewPlayer("", nil)
	cPlayer := &player2.Player{
		Health: 0,
	}
	game := Game{
		ActivePlayer:  player,
		ConsolePlayer: cPlayer,
		Player:        player,
	}
	opponent := game.GetOpponent()
	if opponent != cPlayer {
		t.Errorf("Expected ConsolePlayer found Player")
	}
	game.SwitchPlayer()
	opponent = game.GetOpponent()
	if opponent != player {
		t.Errorf("Expected ConsolePlayer found Player")
	}

}

func TestCheckBleeding(t *testing.T) {

	cPlayer := &player2.Player{
		Health: 1,
		Deck:   []int{},
	}
	game := Game{
		ActivePlayer:  cPlayer,
		ConsolePlayer: cPlayer,
	}
	game.CheckBleedingOut()

	if cPlayer.Health != 0 {
		t.Errorf("Expected 0 health found %v", cPlayer.Health)
	}
}

//
func TestSwitchPlayer(t *testing.T) {

	player := player2.NewPlayer("", nil)
	cPlayer := &player2.Player{
		Health: 0,
	}
	game := Game{
		ActivePlayer:  player,
		ConsolePlayer: cPlayer,
		Player:        player,
	}
	game.SwitchPlayer()

	if game.ActivePlayer != cPlayer {
		t.Errorf("Expected ConsolePlayer found Player")
	}
	game.SwitchPlayer()
	if game.ActivePlayer != player {
		t.Errorf("Expected Player found ConsolePlayer")
	}
}

func TestCanMove(t *testing.T) {

	mckPlayer := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckPlayer)
	mckPlayer.On("GetHand").Return([]int{1, 2, 3})
	mckPlayer.On("GetMana").Return(1)
	play := newGame.CanPlayerMove()

	if !play {
		t.Errorf("Expected true found false")
	}

}

func TestCannotMove(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckPlayer)
	mckPlayer.On("GetHand").Return([]int{1, 2, 3})
	mckPlayer.On("GetMana").Return(0)
	play := newGame.CanPlayerMove()

	if play {
		t.Errorf("Expected false found true")
	}

}

func TestDudPlay(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	mckOpp := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckOpp)
	mckPlayer.On("RemoveCardFromHand", 0)
	mckPlayer.On("EmptiesMana", 0)
	mckOpp.On("TakeDamage", 0)
	mckOpp.On("GetName").Return("Oppa")
	mckOpp.On("GetHealth").Return(1)
	newGame.Move(0)

}
func TestPlay(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	mckOpp := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckOpp)
	mckPlayer.On("RemoveCardFromHand", 4)
	mckPlayer.On("EmptiesMana", 4)
	mckOpp.On("TakeDamage", 4)
	mckPlayer.On("GetName").Return("Emre")
	mckOpp.On("GetName").Return("Oppa")
	mckOpp.On("GetHealth").Return(1)
	newGame.Move(4)
}

func TestPrintState(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckPlayer)
	mckPlayer.On("GetHand").Return([]int{1, 2, 3})
	mckPlayer.On("GetHealth").Return(1)
	mckPlayer.On("GetMana").Return(3)
	newGame.PrintActivePlayerState()
}

func TestPauseGame(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckPlayer)
	rb := new(bytes.Buffer)
	color.NoColor = false
	color.Output = rb
	newGame.PauseGame()
	line, _ := rb.ReadString('\n')

	if !strings.Contains(line, "press Enter to continue ") {
		t.Errorf("Expecting to contains 'press Enter to continue' got '%s'\n", line)
	}
}
func TestPrintRoundInfoOutput(t *testing.T) {
	mckPlayer := new(mocks.IPlayer)
	newGame := NewGame(mckPlayer, mckPlayer, mckPlayer)
	newGame.SetRound(3)
	rb := new(bytes.Buffer)
	color.NoColor = false
	color.Output = rb
	newGame.PrintRoundInfo()
	line, _ := rb.ReadString('2')

	if !strings.Contains(line, " Round 2") {
		t.Errorf("Expecting to contains 'Round 2' got '%s'\n", line)
	}
}
