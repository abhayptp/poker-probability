package models

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []*Card
}

func NewDeck() Deck {
	c := make([]*Card, 0, 52)

	for rank, _ := range rankOrder {
		for _, suit := range suitValues {
			c = append(c, NewCard(rank, suit))
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return Deck{
		cards: c,
	}
}

func (d Deck) RemoveCards(cards []*Card) {
	for _, card := range cards {
		for j, dCard := range d.cards {
			if card.Equals(dCard) {
				d.cards = append(d.cards[:j], d.cards[j+1:]...)
				break
			}
		}
	}
}

func (d Deck) AddCards(cards []*Card) {
	d.cards = append(d.cards, cards...)
}

func (d Deck) Burn() {
	d.cards = append(d.cards[1:], d.cards[0])
}

func (d Deck) GetCardsFromTop(count int) []*Card {
	cards := d.cards[:count]
	d.cards = d.cards[count:]
	return cards
}
