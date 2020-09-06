package cmd

import (
	gm "github.com/emreodabas/kata-tcg-golang/pkg/tcg/game"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/gamer"
	ply "github.com/emreodabas/kata-tcg-golang/pkg/tcg/player"
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"github.com/spf13/cobra"
	"time"
)

func NewGame() *cobra.Command {

	cmd := &cobra.Command{
		Use:          "TCG player_name ",
		SilenceUsage: true,
		Short:        "play Trading Card Game (TCG) on console",
		Args:         cobra.MaximumNArgs(1),
		RunE:         startGame,
	}

	return cmd
}

func startGame(_ *cobra.Command, args []string) error {
	var player, consolePlayer ply.IPlayer
	var name = "player"
	consolePlayer = ply.NewPlayer("TCGamer", new(gamer.AutoGamer))
	if len(args) > 0 {
		name = args[0]
	}
	player = ply.NewPlayer(name, new(gamer.ConsoleGamer))
	var game = gm.Game{
		Player:        player,
		ConsolePlayer: consolePlayer,
	}

	utils.PrintMessage(utils.State, "Welcome to Trading Card Game %s ", name)
	//Randomize Who Start Game
	game.ChooseGameStarter(time.Now().UnixNano())
	game.IterateTurns()
	return nil
}
