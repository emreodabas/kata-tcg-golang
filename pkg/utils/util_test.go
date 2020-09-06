package utils

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"strings"
	"testing"
)

func TestFindMaxLengthAndIndex(t *testing.T) {
	var ints = [][]int{{1}, {2, 3}, {4}}
	index, i := FindMaxLengthAndIndex(ints)

	if index != 1 {
		t.Errorf("Expected to 1 found %v", index)
	}
	if i != 2 {
		t.Errorf("Expected to 2 found %v", i)
	}
}

func TestFindMinLengthAndIndex(t *testing.T) {
	var ints = [][]int{{1}, {2, 3}, {}}
	index, i := FindMinLengthAndIndex(ints)

	if index != 2 {
		t.Errorf("Expected to 2 found %v", index)
	}
	if i != 0 {
		t.Errorf("Expected to 0 found %v", i)
	}
}

func TestFindAndRemove(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}
	remove := FindAndRemove(ints, 3)

	if len(remove) != 4 {
		t.Errorf("Expected to 4 found %v", remove)
	}

	if ContainsValue(remove, 3) {
		t.Errorf("Expected to contains return false found true")
	}
}

func TestRemoveItem(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}
	item := RemoveItem(ints, 3)

	if len(item) != 4 {
		t.Errorf("Expected to 4 found %v", item)
	}

	if ContainsValue(item, 4) {
		t.Errorf("Expected to contains return false found true")
	}
}

func TestContainsValue(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}

	contains := ContainsValue(ints, 2)

	if !contains {
		t.Errorf("Expected to true found false")
	}

	contains = ContainsValue(ints, 6)
	if contains {
		t.Errorf("Expected to false found true")
	}
}

func TestFindItemIndex(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}

	index := FindItemIndex(ints, 1)

	if index < 0 {
		t.Errorf("Expected to true found false")
	}
	index = FindItemIndex(ints, 6)

	if index > 0 {
		t.Errorf("Expected to false found true")
	}
}

func TestPrintMessage(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false

	testTypes := []struct {
		mType MessageType
		text  string
		color color.Attribute
	}{
		{Attack, "magent", color.FgMagenta},
		{Draw, "green", color.FgGreen},
		{State, "blue", color.FgBlue},
		{Warn, "cyan", color.FgCyan},
		{Bleed, "hred", color.FgHiRed},
		{Overload, "hred", color.FgHiRed},
		{Command, "black", color.FgBlack},
		{Error, "red", color.FgRed},
		{End, "hgreen", color.FgHiGreen},
	}

	for _, c := range testTypes {
		PrintMessage(c.mType, c.text)
		line, _ := rb.ReadString('\n')
		scannedLine := fmt.Sprintf("%q", line)
		colored := fmt.Sprintf("%d", c.color)
		escapedForm := fmt.Sprintf("%q", colored)

		fmt.Printf("%s\t: %s\n", c.text, line)

		if !strings.Contains(line, c.text) || !strings.Contains(line, colored) {
			t.Errorf("Expecting %s, got '%s'\n", escapedForm, scannedLine)
		}
	}
}
