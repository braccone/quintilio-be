package models

type Suit string

type Card struct {
	Rank string
	Suit Suit
}

const (
	Bastoni Suit = "bastoni"
	Coppe   Suit = "coppe"
	Denari  Suit = "denari"
	Spade   Suit = "spade"
)

// the card have 4 suits: Spades, Hearts, Diamonds, Clubs and 10 cards per suit
var suits = []Suit{Bastoni, Coppe, Denari, Spade}

// the rank is the following order: 3, 2, 1, 10, 9, 8, 7, 6, 5, 4
var ranks = []string{"3", "2", "1", "10", "9", "8", "7", "6", "5", "4"}
