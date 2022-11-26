package models

var _ DealtCards = &dealtCards{}

type DealtCards interface {
	GetCards() []Card
	SetUnknownDealtCards(...Card)
}

type dealtCards struct {
	cards []Card
}

func NewDealtCards(c ...Card) *dealtCards {
	if len(c) > 2 {
		return nil
	}

	if len(c) == 2 {
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
	d.cards = append(d.cards, c...)
}

func (d *dealtCards) String() string {
	if len(d.cards) == 0 {
		return "na"
	}
	res := ""
	for i, v := range d.cards {
		res += v.String()
		if i != len(d.cards)-1 {
			res += ","
		}
	}
	return res
}
