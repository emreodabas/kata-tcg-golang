package gamer

import (
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"sort"
)

type AutoGamer struct {
}

func (c *AutoGamer) GetGamerType() GamerType {
	return Auto
}

func (c *AutoGamer) GetCardDecision(hand []int, mana int) int {

	combinations := findCombinations(hand, mana)
	//no solution -> skip hand
	if len(combinations) == 0 {
		utils.PrintMessage(utils.Warn, "Skipping hand")
		return -1
		//only one solution -> do it
	} else if len(combinations) == 1 {
		ints := combinations[0]
		utils.PrintMessage(utils.State, "One Solution to move %v", ints)
		return ints[0]
	} else {
		index, maxLen := utils.FindMaxLengthAndIndex(combinations)
		//max card over max damage strategy
		if (len(hand) - maxLen) > 0 {
			utils.PrintMessage(utils.State, "Selecting Max Card Solution %v", combinations[index])
			return combinations[index][0]
		} else {
			index, _ := utils.FindMinLengthAndIndex(combinations)
			utils.PrintMessage(utils.State, "Selecting Min Card Solution %v", combinations[index])
			return combinations[index][0]
		}
	}
}
func findCombinations(cards []int, mana int) [][]int {
	sort.Ints(cards)
	return genCombinations(cards, mana)
}
func genCombinations(cards []int, mana int) [][]int {
	solution := make([][]int, 0)
	for i, candidate := range cards {
		// if two consecutive number is same then skip
		if i != 0 && cards[i] == cards[i-1] {
			continue
		}
		if candidate <= mana {
			solution = append(solution, []int{candidate})
		}

		tmp := genCombinations(cards[i+1:], mana-candidate)
		for _, val := range tmp {
			solution = append(solution, append(val, candidate))
		}
	}
	return solution
}
