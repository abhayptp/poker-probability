package models

import "errors"

type Rank string

type Suit rune

var rankOrder = map[Rank]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

var suitValues = [4]Suit{'S', 'H', 'D', 'C'}

var _ Card = &card{}

type Card interface {
	Equals(Card) bool
	EqualsRank(Card) bool
	EqualsSuit(Card) bool
	GetRank() Rank
	GetSuit() Suit
	GreaterThan(Card) bool
	LessThan(Card) bool
	RankDifference(Card) int
	String() string
	Valid() bool
}

type card struct {
	rank Rank
	suit Suit
}

func NewCard(rank Rank, suit Suit) (*card, error) {
	c := &card{
		rank: Rank(rank),
		suit: Suit(suit),
	}

	if !c.Valid() {
		return nil, errors.New("rank or suit not valid")
	}

	return c, nil
}

func (card1 *card) Equals(card2 Card) bool {
	if card1.EqualsSuit(card2) && card1.EqualsRank(card2) {
		return true
	}

	return false
}

func (card1 *card) EqualsRank(card2 Card) bool {
	return card1.rank == card2.GetRank()
}

func (card1 *card) EqualsSuit(card2 Card) bool {
	return card1.suit == card2.GetSuit()
}

func (card *card) GetRank() Rank {
	return card.rank
}

func (card *card) GetSuit() Suit {
	return card.suit
}

func (card1 *card) GreaterThan(card2 Card) bool {
	return rankOrder[card1.rank] > rankOrder[card2.GetRank()]
}

func (card1 *card) LessThan(card2 Card) bool {
	return rankOrder[card1.rank] < rankOrder[card2.GetRank()]
}

func (card1 *card) RankDifference(card2 Card) int {
	return rankOrder[card1.rank] - rankOrder[card2.GetRank()]
}

func (card *card) String() string {
	return string(card.rank) + string(card.suit)
}

func (card *card) Valid() bool {
	if card == nil {
		return false
	}

	if _, ok := rankOrder[card.rank]; !ok {
		return false
	}

	for _, possibleSuit := range suitValues {
		if card.suit == possibleSuit {
			return true
		}
	}

	return false
}
