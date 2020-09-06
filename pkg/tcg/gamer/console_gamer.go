package gamer

import (
	"bufio"
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"os"
	"strconv"
	"strings"
)

type ConsoleGamer struct {
}

func (c *ConsoleGamer) GetGamerType() GamerType {
	return Console
}
func (c *ConsoleGamer) GetCardDecision(hand []int, mana int) int {

	utils.PrintMessage(utils.Command, "Write skip or press Enter for Skipping Hand")
	utils.PrintMessage(utils.Command, "Choose Card to Move:")
	for {
		var chosen string
		chosen = ReadChosen()
		if strings.Contains(chosen, "skip") || chosen == "" {
			return -1
		}
		damage, err := strconv.Atoi(chosen)
		if err != nil {
			utils.PrintMessage(utils.Error, "Please write valid number")
			break
		}
		if utils.ContainsValue(hand, damage) {
			if damage > mana {
				utils.PrintMessage(utils.Error, "You don't have sufficient mana %v < %v", mana, damage)
				break
			} else {
				return damage
				break
			}
		} else {
			utils.PrintMessage(utils.Error, "You don't have %v card ", damage)
			break
		}
	}
	return -1
}

func ReadChosen() string {
	line, _, _ := GetReader().ReadLine()
	return string(line)
}

func SetReader(in *bufio.Reader) {
	Reader = in
}

func GetReader() *bufio.Reader {
	return Reader
}

var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
