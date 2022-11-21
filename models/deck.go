package models

import (
	"math/rand"
	"time"
)

type Deck interface {
	RemoveCards([]Card)
	AddCards([]Card)
	Burn()
	GetCardsFromTop(int) []Card
}

type deck struct {
	cards []Card
}

func NewDeck() *deck {
	c := make([]Card, 0, 52)

	for rank := range rankOrder {
		for _, suit := range suitValues {
			c = append(c, NewCard(rank, suit))
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return &deck{
		cards: c,
	}
}

func (d *deck) RemoveCards(cards []Card) {
	for _, card := range cards {
		for j, dCard := range d.cards {
			if card.Equals(dCard) {
				d.cards = append(d.cards[:j], d.cards[j+1:]...)
				break
			}
		}
	}
}

func (d *deck) AddCards(cards []Card) {
	d.cards = append(d.cards, cards...)
}

func (d *deck) Burn() {
	d.cards = append(d.cards[1:], d.cards[0])
}

func (d *deck) GetCardsFromTop(count int) []Card {
	cards := d.cards[:count]
	d.cards = d.cards[count:]
	return cards
}
