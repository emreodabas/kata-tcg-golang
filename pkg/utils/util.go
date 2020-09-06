package utils

import (
	"github.com/fatih/color"
)

func ContainsValue(ints []int, value int) bool {
	for _, a := range ints {
		if a == value {
			return true
		}
	}
	return false
}
func FindAndRemove(s []int, i int) []int {
	return RemoveItem(s, FindItemIndex(s, i))
}
func FindMaxLengthAndIndex(ints [][]int) (int, int) {
	var maxLen, index = -1, -1
	var max = 0
	for k, v := range ints {
		if len(v) >= maxLen {
			if Sum(v) >= max {
				max = Sum(v)
				index = k
				maxLen = len(v)
			}
		}
	}
	return index, maxLen
}
func FindMinLengthAndIndex(ints [][]int) (int, int) {
	var minLen, index = 100, -1
	for k, v := range ints {
		if len(v) <= minLen {
			index = k
			minLen = len(v)
		}
	}
	return index, minLen
}
func FindItemIndex(s []int, i int) int {
	for k, v := range s {
		if v == i {
			return k
		}
	}
	return -1
}
func RemoveItem(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
func Sum(s []int) int {
	var sum = 0
	for _, v := range s {
		sum = sum + v
	}
	return sum
}

type MessageType int

const (
	Attack MessageType = iota
	Draw
	State
	Warn
	Bleed
	Overload
	Command
	Error
	End
)

func PrintMessage(messageType MessageType, message string, a ...interface{}) {
	switch messageType {
	case Attack:
		color.Magenta(message, a...)
	case Draw:
		color.Green(message, a...)
	case End:
		color.HiGreen(message, a...)
	case State:
		color.Blue(message, a...)
	case Warn:
		color.Cyan(message, a...)
	case Bleed:
		color.HiRed(message, a...)
	case Overload:
		color.HiRed(message, a...)
	case Command:
		color.Black(message, a...)
	case Error:
		color.Red(message, a...)
	}
}
