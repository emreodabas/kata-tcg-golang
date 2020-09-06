package gamer

import (
	"bytes"
	"github.com/fatih/color"
	"testing"
)

func TestAutoGamer_GetCardDecision(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	auto := new(AutoGamer)
	hand := []int{0, 0, 2, 3}
	mana := 2
	cardDecision := auto.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != 2 {
		t.Errorf("Expected 2 found %v", cardDecision)
	}
}
func TestAutoGamer_GetCardDecisionSkip(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	auto := new(AutoGamer)
	hand := []int{7, 2, 3}
	mana := 0
	cardDecision := auto.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != -1 {
		t.Errorf("Expected -1 found %v", cardDecision)
	}
}
func TestAutoGamer_GetCardDecisionMinLengthStrategy(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	auto := new(AutoGamer)
	hand := []int{2, 1}
	mana := 3
	cardDecision := auto.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != 2 {
		t.Errorf("Expected 2 found %v", cardDecision)
	}
}

func TestAutoGamer_GetCardDecisionSingleSolution(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	auto := new(AutoGamer)
	hand := []int{2, 3}
	mana := 2
	cardDecision := auto.GetCardDecision(hand, mana)
	//skipped all
	if cardDecision != 2 {
		t.Errorf("Expected 2 found %v", cardDecision)
	}
}

func TestAutoGamer_GetGamerType(t *testing.T) {
	consoleGamer := new(AutoGamer)
	gamerType := consoleGamer.GetGamerType()

	if gamerType != Auto {
		t.Errorf("Expected to %v found %v ", Auto, gamerType)
	}
}
