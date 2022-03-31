package main

import (
	models "github.com/braccone/quintilio-be/models"
)

// Create a card game composed of a deck of 40 cards and 5 players.
// Each player draws 8 cards from the deck.
// Each player starts with a score of 0.
// Each player can only play one card at a time.
// The player who starts the game can choice the card he wants to play with from his hand
// Other players must play a card from their hands that has the same suit as the first card played.
// The owner of the played card

func main() {
	// Create a new deck of cards
	deck := models.NewDeck()
	// Shuffle the deck
	deck.Shuffle()
	// Create an array of 5 players
	players := [5]*models.Player{
		models.NewPlayer("Player 1"),
		models.NewPlayer("Player 2"),
		models.NewPlayer("Player 3"),
		models.NewPlayer("Player 4"),
		models.NewPlayer("Player 5"),
	}

	// init game
	game := models.NewGame(players)

	game.PrintPlayers()

}
