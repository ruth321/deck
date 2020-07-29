package deck

import (
	"math/rand"
	"strconv"
)

type Card struct {
	Suit string
	Name string
}

type Deck interface {
	Shuffle() []Card
}

type Cards []Card

func New() []Card {
	var deck []Card
	var card Card
	for i := 1; i < 5; i++ {
		switch i {
		case 1:
			card.Suit = "spade"
		case 2:
			card.Suit = "diamond"
		case 3:
			card.Suit = "club"
		case 4:
			card.Suit = "heart"
		}
		for g := 1; g < 14; g++ {
			switch g {
			case 1:
				card.Name = "A"
			case 11:
				card.Name = "J"
			case 12:
				card.Name = "Q"
			case 13:
				card.Name = "K"
			default:
				card.Name = strconv.Itoa(g)
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func (deck *Cards) Shuffle() []Card {
	var card Card
	l := len(*deck)
	for i := 0; i < l-1; i++ {
		g := i + 1 + rand.Intn(l-i)
		card = (*deck)[i]
		(*deck)[i] = (*deck)[g]
		(*deck)[g] = card
	}
	return *deck
}
