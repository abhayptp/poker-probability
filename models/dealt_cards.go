package models

type DealtCards interface {
	GetCards() []Card
	SetUnknownDealtCards(...Card)
}

type dealtCards struct {
	playerName string
	cards      []Card
}

func NewDealtCards(c ...Card) *dealtCards {
	if len(c) > 2 {
		panic("known dealt cards can not be greater than 2")
	}

	if c[0] != nil && c[1] != nil {
		if c[0].LessThan(c[1]) {
			c = append(c[1:], c[:1]...)
		}
	}

	return &dealtCards{
		cards: c,
	}
}

func (d *dealtCards) GetCards() []Card {
	return d.cards[:]
}

func (d *dealtCards) SetUnknownDealtCards(c ...Card) {
	if len(c)+len(d.cards) != 2 {
		panic("total dealt cards should be 2")
	}

	d.cards = append(d.cards, c...)
}
