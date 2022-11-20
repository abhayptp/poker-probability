package models

type CommunityCards struct {
	openedCards   []*Card
	unopenedCards []*Card
}

func NewCommunityCards(cards ...*Card) CommunityCards {

	return CommunityCards{
		openedCards: cards,
	}
}

func (c CommunityCards) GetOpenedCards() []*Card {
	return c.openedCards
}

func (c CommunityCards) GetOpenedCardsCount() int {
	return len(c.openedCards)
}

func (c CommunityCards) SetUnopenedCards(cards []*Card) {
	c.unopenedCards = cards
}

func (c CommunityCards) GetAllCards() []*Card {
	return append(c.openedCards, c.unopenedCards...)
}
