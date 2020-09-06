package game

import (
	"fmt"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/gamer"
	"github.com/emreodabas/kata-tcg-golang/pkg/tcg/player"
	"github.com/emreodabas/kata-tcg-golang/pkg/utils"
	"sort"
)

type Game struct {
	ActivePlayer          player.IPlayer
	Player, ConsolePlayer player.IPlayer
	round                 int
}

type IGame interface {
	ChooseGameStarter(random int64)
	IterateTurns()
	PrintRoundInfo()
	Turn()
	MoveActions()
	Move(card int)
	SwitchPlayer()
	GameEnded() bool
	GetOpponent() player.IPlayer
	CheckBleedingOut()
	CanPlayerMove() bool
	PauseGame()
	PrintActivePlayerState()
	GetRound() int
	SetRound(round int)
}

func NewGame(activePlayer player.IPlayer, player player.IPlayer, console player.IPlayer) IGame {
	return &Game{
		ActivePlayer:  activePlayer,
		Player:        player,
		ConsolePlayer: console,
		round:         0,
	}
}

func (g *Game) ChooseGameStarter(random int64) {
	if random%2 == 1 {
		g.ActivePlayer = g.Player
		utils.PrintMessage(utils.State, "You are the chosen one. Do your best ...")
	} else {
		g.ActivePlayer = g.ConsolePlayer
		utils.PrintMessage(utils.State, "%s is lucky to start", g.ActivePlayer.GetName())
	}
}
func (g *Game) IterateTurns() {
	for !g.GameEnded() {
		g.round = g.round + 1
		g.PrintRoundInfo()
		g.CheckBleedingOut()
		g.Turn()
		g.SwitchPlayer()
	}
	if g.Player.GetHealth() > 0 {
		utils.PrintMessage(utils.End, "Congratulations, YOU WIN ")
	} else {
		utils.PrintMessage(utils.End, "Congratulations to my DUMMY Player WIN :) ")

	}
}
func (g *Game) PrintRoundInfo() {
	if g.GetRound()%2 == 1 {
		utils.PrintMessage(utils.State, "-------------------------------")
		utils.PrintMessage(utils.State, "\t  Round %v \t", g.GetRound()/2+1)
		utils.PrintMessage(utils.State, "-------------------------------")
	}
}
func (g *Game) Turn() {
	g.ActivePlayer.IncreaseManaSlot()
	g.ActivePlayer.RefillMana()
	err := g.ActivePlayer.DrawCard()
	if err != nil {
		utils.PrintMessage(utils.Error, "Deck is over")
	}
	g.MoveActions()
}
func (g *Game) MoveActions() {
	if !g.CanPlayerMove() {
		g.PrintActivePlayerState()
		utils.PrintMessage(utils.Warn, "No Card could be Move")
		utils.PrintMessage(utils.Warn, "Skipping Hand ")
		g.PauseGame()
	} else {
		i := 0
		for g.CanPlayerMove() && i <= player.MaxHandSize {
			i += 1
			g.PrintActivePlayerState()
			card := g.ActivePlayer.GetGamer().GetCardDecision(g.ActivePlayer.GetHand(), g.ActivePlayer.GetMana())
			g.Move(card)
			if g.ActivePlayer.GetGamer().GetGamerType() == gamer.Auto {
				g.PauseGame()
			}
		}
	}
}
func (g *Game) Move(card int) {
	if card >= 0 {
		g.ActivePlayer.RemoveCardFromHand(card)
		g.ActivePlayer.EmptiesMana(card)
		g.GetOpponent().TakeDamage(card)
		if card == 0 {
			utils.PrintMessage(utils.Attack, "Dud Move!!!, No card")
		} else {
			utils.PrintMessage(utils.Attack, "%s  attack %v ", g.ActivePlayer.GetName(), card)
		}
	}
	utils.PrintMessage(utils.Attack, "%s Remaining Health:::: %v", g.GetOpponent().GetName(), g.GetOpponent().GetHealth())
}
func (g *Game) SwitchPlayer() {
	if g.ActivePlayer == g.ConsolePlayer {
		g.ActivePlayer = g.Player
	} else {
		g.ActivePlayer = g.ConsolePlayer
	}
}
func (g *Game) GameEnded() bool {
	if g.Player.GetHealth() <= 0 || g.ConsolePlayer.GetHealth() <= 0 {
		return true
	}
	return false
}
func (g *Game) GetOpponent() player.IPlayer {
	if g.ActivePlayer == g.ConsolePlayer {
		return g.Player
	} else {
		return g.ConsolePlayer
	}
}
func (g *Game) CheckBleedingOut() {
	if len(g.ActivePlayer.GetDeck()) == 0 {
		utils.PrintMessage(utils.Bleed, "Bleeding Out !!! Cause of empty deck ")
		g.ActivePlayer.TakeDamage(1)
	}
}
func (g *Game) CanPlayerMove() bool {
	for i := 0; i < len(g.ActivePlayer.GetHand()); i++ {
		if g.ActivePlayer.GetHand()[i] <= g.ActivePlayer.GetMana() {
			return true
		}
	}
	return false
}
func (g *Game) PauseGame() {
	utils.PrintMessage(utils.Warn, "(press Enter to continue )")
	fmt.Scanln()
}
func (g *Game) PrintActivePlayerState() {
	sort.Ints(g.ActivePlayer.GetHand())
	utils.PrintMessage(utils.State, "HEALTH:: %v ", g.ActivePlayer.GetHealth())
	utils.PrintMessage(utils.State, "MANA:: %v ", g.ActivePlayer.GetMana())
	utils.PrintMessage(utils.State, "HAND:: %v ", g.ActivePlayer.GetHand())
}
func (g *Game) GetRound() int {
	return g.round
}
func (g *Game) SetRound(round int) {
	g.round = round
}
