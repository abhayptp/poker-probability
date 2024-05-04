package models

var _ CommunityCards = &communityCards{}

type CommunityCards interface {
	Clone() CommunityCards
	GetAllCards() []Card
	GetOpenedCards() []Card
	GetUnopenedCards() []Card
	SetUnopenedCards(cards []Card)
}

type communityCards struct {
	openedCards   []Card
	unopenedCards []Card
}

func NewCommunityCards(cards ...Card) *communityCards {

	return &communityCards{
		openedCards: cards,
	}
}

func (c *communityCards) Clone() CommunityCards {
	clone := *c
	return &clone
}

func (c *communityCards) GetAllCards() []Card {
	return append(c.openedCards, c.unopenedCards...)
}

func (c *communityCards) GetOpenedCards() []Card {
	return c.openedCards
}

func (c *communityCards) GetUnopenedCards() []Card {
	return c.unopenedCards
}

func (c *communityCards) SetUnopenedCards(cards []Card) {
	c.unopenedCards = cards
}
