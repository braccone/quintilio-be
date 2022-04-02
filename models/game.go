package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Players    [5]*Player
	Team1      [2]*Player
	Team2      [3]*Player
	PlayerTurn int
	Dealer     int
}

func NewGame(players [5]*Player) *Game {
	// Create a new deck of cards
	deck := NewDeck()
	// Shuffle the deck
	deck.Shuffle()
	// while deck is not empty draw a card for each player
	for len(deck.Cards) > 0 {
		for _, player := range players {
			player.AddCard(deck.Draw())
		}

	}
	// random index of player turn
	rand.Seed(time.Now().Unix())
	playerTurn := rand.Intn(len(players))

	g := &Game{Players: players, PlayerTurn: playerTurn}
	g.InitDealer()
	return g
}

// Set dealer as the player before player turn
func (g *Game) InitDealer() {
	g.Dealer = g.PlayerTurn - 1
	if g.Dealer < 0 {
		g.Dealer = len(g.Players) - 1
	}
}

// next player as dealer
func (g *Game) NextDealer() {
	// if dealer is not the last player
	if g.Dealer != len(g.Players)-1 {
		// set dealer to next player
		g.Dealer++
	} else {
		// set dealer to first player
		g.Dealer = 0
	}
}

// next player turn
func (g *Game) NextPlayerTurn() {
	g.PlayerTurn = (g.PlayerTurn + 1) % len(g.Players)
}

// Print players hands
func (g *Game) PrintPlayers() {
	for i := 0; i < len(g.Players); i++ {
		g.Players[i].PrintHand()
	}
	// print current player turn
	fmt.Println("Current player turn: ", g.Players[g.PlayerTurn].Name)

}

// Once all the players has 8 cards each, there is pre start game phase.
// In this phase, starting player have a choice to do one of the following:
// - Call if he has at least one card with rank 3. In this case, if other players agree with the call choice the caller nominate a card with
// rank 3 and different suit from the card with rank 3 in his hand.
// - Pass means pass the choice to next player.
func (g *Game) FirstPhase() {
	// starting player
	startPlayer := g.Players[g.PlayerTurn]
	startPlayer.MenuChoice()

	// if starting player call
	if startPlayer.Choice == CALL {
		// starting player nominate a card with rank 3
		startPlayer.Call()
		// startPlayer.NominateCard(3)
	}
	if startPlayer.Choice == PASS {
		g.NextPlayerTurn()
	}
}
