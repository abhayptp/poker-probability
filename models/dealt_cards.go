package models

type DealtCards struct {
	playerName string
	cards      []*Card
}

func NewDealtCards(c ...*Card) DealtCards {
	assert(len(c) <= 2)

	if c[0] != nil && c[1] != nil {
		if c[0].LessThan(c[1]) {
			c = append(c[1:], c[:1]...)
		}
	}

	return DealtCards{
		cards: c,
	}
}

func (d DealtCards) GetCards() []*Card {
	return d.cards[:]
}

func (d DealtCards) SetUnknownDealtCards(c ...*Card) {
	assert(len(c)+len(d.cards) == 2)
	d.cards = append(d.cards, c...)
}
