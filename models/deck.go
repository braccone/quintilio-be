package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	var deck Deck
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Shuffle the deck of cards using random seed
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().Unix())
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// random draw a card from the deck
func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// print the deck of cards
func (d *Deck) Print() {
	for _, card := range d.Cards {
		fmt.Println(card)
	}
}
