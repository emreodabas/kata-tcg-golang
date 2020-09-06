package gamer

import (
	"bufio"
	"bytes"
	"github.com/fatih/color"
	"strings"
	"testing"
)

func TestConsoleGamer_GetCardDecision(t *testing.T) {

	consoleGamer := new(ConsoleGamer)
	hand := []int{0, 2, 3}
	mana := 3
	input := "3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	cardDecision := consoleGamer.GetCardDecision(hand, mana)
	if cardDecision != 3 {
		t.Errorf("Expected 3 found %v", cardDecision)
	}
}

func TestConsoleGamer_GetCardDecisionInvalid(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	consoleGamer := new(ConsoleGamer)
	hand := []int{0, 2, 3}
	mana := 2
	input := "asdasda\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	cardDecision := consoleGamer.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != -1 {
		t.Errorf("Expected -1 found %v", cardDecision)
	}
	readString, _ := rb.ReadString('9')

	if !strings.Contains(readString, "Please write valid number") {
		t.Errorf("Expection output contains Invalid but not found %s", readString)
	}
}

func TestConsoleGamer_GetCardDecisionInsufficientMana(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	consoleGamer := new(ConsoleGamer)
	hand := []int{0, 2, 3}
	mana := 2
	input := "3\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	cardDecision := consoleGamer.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != -1 {
		t.Errorf("Expected -1 found %v", cardDecision)
	}
	readString, _ := rb.ReadString('9')

	if !strings.Contains(readString, "You don't have sufficient mana") {
		t.Errorf("Expection output contains 'You don't have sufficient mana' but not found %s", readString)
	}
}

func TestConsoleGamer_GetCardDecisionSkip(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	consoleGamer := new(ConsoleGamer)
	hand := []int{0, 2, 3}
	mana := 2
	input := "skip\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	cardDecision := consoleGamer.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != -1 {
		t.Errorf("Expected -1 found %v", cardDecision)
	}
}

func TestConsoleGamer_GetCardDecisionCardAbsent(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	consoleGamer := new(ConsoleGamer)
	hand := []int{0, 2, 3}
	mana := 2
	input := "4\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	cardDecision := consoleGamer.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != -1 {
		t.Errorf("Expected -1 found %v", cardDecision)
	}
	readString, _ := rb.ReadString('9')

	if !strings.Contains(readString, "You don't have 4 card") {
		t.Errorf("Expection output contains 'You don't have 4 card' but not found %s", readString)
	}
}

func TestReadChosen(t *testing.T) {

	input := "Test\n"
	reader := bufio.NewReader(strings.NewReader(input))
	SetReader(reader)
	chosen := ReadChosen()
	if chosen != strings.TrimSpace(input) {
		t.Fatal()
	}
}

func TestConsoleGamer_GetGamerType(t *testing.T) {
	consoleGamer := new(ConsoleGamer)
	gamerType := consoleGamer.GetGamerType()

	if gamerType != Console {
		t.Errorf("Expected to %v found %v ", Console, gamerType)
	}
}
