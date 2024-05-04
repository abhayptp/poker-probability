package models

import (
	"math/rand"
	"time"
)

var _ Deck = &deck{}

type Deck interface {
	AddCards([]Card)
	Contains(Card) bool
	Burn()
	PopCardsFromTop(int) []Card
	RemoveCards([]Card)
}

type deck struct {
	cards []Card
}

func NewDeck() *deck {
	d := make([]Card, 0, 52)

	for rank := range rankOrder {
		for _, suit := range suitValues {
			c, _ := NewCard(rank, suit)
			d = append(d, c)
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return &deck{
		cards: d,
	}
}

func (d *deck) AddCards(cards []Card) {
	d.cards = append(d.cards, cards...)
}

func (d *deck) Burn() {
	d.cards = append(d.cards[1:], d.cards[0])
}

func (d *deck) Contains(c Card) bool {
	for _, card := range d.cards {
		if card.Equals(c) {
			return true
		}
	}
	return false
}

func (d *deck) PopCardsFromTop(count int) []Card {
	cards := d.cards[:count]
	d.cards = d.cards[count:]
	return cards
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
