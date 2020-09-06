package main

import (
	"os"

	"github.com/emreodabas/kata-tcg-golang/pkg/cmd"
)

func main() {
	game := cmd.NewGame()
	if err := game.Execute(); err != nil {
		os.Exit(1)
	}
}
