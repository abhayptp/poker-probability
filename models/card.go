package models

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

type Card struct {
	rank Rank
	suit Suit
}

func NewCard(rank Rank, suit Suit) *Card {
	c := &Card{
		rank: Rank(rank),
		suit: Suit(suit),
	}

	if !c.Valid() {
		panic("invalid card")
	}

	return c
}

func (card *Card) GetRank() *Rank {
	return &card.rank
}

func (card *Card) GetSuit() *Suit {
	return &card.suit
}

func (card *Card) Valid() bool {
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

func (card1 *Card) LessThan(card2 *Card) bool {
	if rankOrder[card1.rank] < rankOrder[card2.rank] {
		return true
	}

	return false
}

func (card1 *Card) GreaterThan(card2 *Card) bool {
	if rankOrder[card1.rank] > rankOrder[card2.rank] {
		return true
	}

	return false
}

func (card1 *Card) Equals(card2 *Card) bool {
	if card1.EqualsSuit(card2) && card1.EqualsRank(card2) {
		return true
	}

	return false
}

func (card1 *Card) EqualsSuit(card2 *Card) bool {
	if card1.suit == card2.suit {
		return true
	}

	return false
}

func (card1 *Card) EqualsRank(card2 *Card) bool {
	if card1.rank == card1.rank {
		return true
	}

	return false
}

func (card1 *Card) RankDifference(card2 *Card) int {
	return rankOrder[card1.rank] - rankOrder[card2.rank]

}

func (card *Card) EqualsString(c string) bool {
	suit := c[len(c)-1]
	rank := c[:len(c)-1]

	if string(card.rank) == rank && rune(card.suit) == rune(suit) {
		return true
	}

	return false
}
