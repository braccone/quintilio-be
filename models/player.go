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

// create a menu choice with the player's hand and get the choice
func (p *Player) MenuChoice() {
	p.PrintHand()
	// if the player has not a card with rank 3, he can only pass

	if !p.HasCardWithRank3() {
		p.MakeChoice(PASS)
		fmt.Println("PASSED")
		return
	}

	fmt.Println("0: pass")
	fmt.Println("1: call")
	fmt.Println("2: exchange")

	// get the choice from user input
	var choice int
	// fmt.Scanf("%d", &choice)
	// print choice
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Printf("Scan for i failed, due to %s", err)
		return
	}
	// make the choice
	switch choice {
	case 0:
		{
			p.MakeChoice(PASS)
			break
		}
	case 1:
		{
			p.MakeChoice(CALL)
			break
		}
	case 2:
		{
			p.MakeChoice(EXCHANGE)
			break
		}
	default:
		{
			// error msg and retry
			fmt.Println("Error: invalid choice")
			p.MenuChoice()
			break
		}
	}

}

// create a menu choice with the suit of the card with rank3 not contained in the hand
func (p *Player) Call() {
	fmt.Printf("Choose a card to call: \n")

	for i, suit := range suits {
		for _, c := range p.Hand {
			if c.Rank == "3" && c.Suit != suit {
				fmt.Printf("%d: { 3 %s }\n", i, suit)
			}
		}
	}

	// get the choice from user input
	var choice int
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Printf("Scan for i failed, due to %s", err)
		return
	}
	// make the choice
	switch choice {
	case 0:
		{
			card := Card{Rank: "3", Suit: "bastoni"}
			p.CalledCard = card
			break
		}
	case 1:
		{
			card := Card{Rank: "3", Suit: "coppe"}
			p.CalledCard = card
			break
		}
	case 2:
		{
			card := Card{Rank: "3", Suit: "denari"}
			p.CalledCard = card
			break
		}
	case 3:
		{
			card := Card{Rank: "3", Suit: "spade"}
			p.CalledCard = card
			break
		}
	default:
		{
			p.Call()
			break
		}
	}
}
