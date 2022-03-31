package models

import "fmt"

type Choices string

const (
	CALL     Choices = "call"
	EXCHANGE Choices = "exchange"
	PASS     Choices = "pass"
)

type Player struct {
	Name       string
	Hand       []Card
	CardWon    []Card
	Choice     Choices
	Score      int
	CalledCard Card
}

// Create a new player
func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

// Add a card to the player's hand
func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
}

// Remove a card from the player's hand
func (p *Player) RemoveCard(card Card) {
	for i, c := range p.Hand {
		if c == card {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
		}
	}
}

// Add cards to the player's won cards
func (p *Player) AddCardWon(card Card) {
	p.CardWon = append(p.CardWon, card)
}

// print the player's hand
func (p *Player) PrintHand() {
	fmt.Printf("Player: %s\n", p.Name)
	for _, c := range p.Hand {
		fmt.Print(c)
	}
	fmt.Println()
}

// make a choice
func (p *Player) MakeChoice(choice Choices) {
	p.Choice = choice
}

// player has card with rank 3
func (p *Player) HasCardWithRank3() bool {
	for _, c := range p.Hand {
		if c.Rank == "3" {
			return true
		}
	}
	return false
}

// call or pass the choice
func (p *Player) CallOrPass() {
	if p.HasCardWithRank3() {
		p.MakeChoice(CALL)
	} else {
		p.MakeChoice(PASS)
	}
}

// nominate a card with rank 3 and set player called card
func (p *Player) NominateCard(rank int) {
	for _, c := range p.Hand {
		if c.Rank == fmt.Sprintf("%d", rank) {
			p.CalledCard = c
		}
	}
}
