package models

type Game struct {
	Players [5]*Player
}

func NewGame(players [5]*Player) *Game {
	// Create a new deck of cards
	deck := NewDeck()
	// Shuffle the deck
	deck.Shuffle()
	// while deck is not empty draw a card for each player
	for len(deck.Cards) > 0 {
		for i := 0; i < len(players); i++ {
			players[i].AddCard(deck.Draw())
		}
	}
	// print the players' hands
	g := &Game{Players: players}
	return g
}

// Print players hands
func (g *Game) PrintPlayers() {
	for i := 0; i < len(g.Players); i++ {
		g.Players[i].PrintHand()
	}
}

// Once all the players has 8 cards each, there is pre start game phase.
// In this phase, starting player have a choice to do one of the following:
// - Call if he has at least one card with rank 3. In this case, if other players agree with the call choice the caller nominate a card with
// rank 3 and different suit from the card with rank 3 in his hand.
// - Pass means pass the choice to next player.
func (g *Game) PreStartGame(players [5]*Player) {
	// starting player
	startPlayer := players[0]
	// starting player can call or pass
	startPlayer.CallOrPass()
	// if starting player call
	if startPlayer.Choice == CALL {
		// starting player nominate a card with rank 3
		startPlayer.NominateCard(3)
	}
}
