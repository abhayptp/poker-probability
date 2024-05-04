package models

import (
	"sort"
)

type HandType int

const (
	RoyalFlush HandType = iota
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPairs
	Pair
	HighCard
	Default
)

var _ Hand = &hand{}

type Hand interface {
	Compare(Hand) int
	GetHandType() HandType
	GetTieBreakerCardsOrder() []Card
	SetHandType()
	String() string
}

type hand struct {
	cards []Card

	cardsRankCountMap    map[Rank]int8
	handType             HandType
	tieBreakerCardsOrder []Card
}

func NewHand(cards []Card) *hand {
	if len(cards) != 5 {
		panic("hand should have 5 cards")
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].LessThan(cards[j])
	})

	m := make(map[Rank]int8)

	for _, card := range cards {
		m[card.GetRank()] += 1
	}

	tieBreakerCardsOrder := make([]Card, 5)
	_ = copy(tieBreakerCardsOrder, cards)
	sort.Slice(tieBreakerCardsOrder, func(i, j int) bool {
		c1 := tieBreakerCardsOrder[i]
		c2 := tieBreakerCardsOrder[j]

		if c1 == c2 {
			return false
		}

		if m[c1.GetRank()] > m[c2.GetRank()] {
			return true
		}

		if m[c1.GetRank()] < m[c2.GetRank()] {
			return false
		}

		return c1.GreaterThan(c2)
	})

	return &hand{
		cards: cards,

		cardsRankCountMap:    m,
		tieBreakerCardsOrder: tieBreakerCardsOrder,
		// Default hand type is HighCard
		handType: Default,
	}
}

// Return value:
// 	1  if h1 > h2
//	0 if h1 == h2
//	-1  if h1 < h2
func (h1 *hand) Compare(h2 Hand) int {
	if h1.handType == Default || h2.GetHandType() == Default {
		panic("hand type not set")
	}

	if h1.handType < h2.GetHandType() {
		return 1
	}

	if h1.handType > h2.GetHandType() {
		return -1
	}

	for i := 0; i < 5; i++ {
		if h1.tieBreakerCardsOrder[i].GreaterThan(h2.GetTieBreakerCardsOrder()[i]) {
			return 1
		}

		if h1.tieBreakerCardsOrder[i].LessThan(h2.GetTieBreakerCardsOrder()[i]) {
			return -1
		}
	}

	return 0
}

func (h *hand) SetHandType() {
	fList := [10]func() bool{
		h.royalFlush,
		h.straightFlush,
		h.fourOfAKind,
		h.fullHouse,
		h.flush,
		h.straight,
		h.threeOfAKind,
		h.twoPairs,
		h.pair,
		h.highCard,
	}

	for _, f := range fList {
		if f() {
			return
		}
	}
}

func (h *hand) GetHandType() HandType {
	if h.handType == Default {
		panic("Hand type not set")
	}

	return h.handType
}

func (h *hand) GetTieBreakerCardsOrder() []Card {
	return h.tieBreakerCardsOrder
}

func (h *hand) String() string {
	st := ""
	len := len(h.cards)
	for i, card := range h.cards {
		st += card.String()
		if i != len-1 {
			st += ","
		}
	}
	return st
}

func (h *hand) royalFlush() bool {
	if h.cards[0].GetRank() != "10" {
		return false
	}

	if h.straightFlush() {
		h.handType = RoyalFlush
		return true
	}

	return false
}

func (h *hand) straightFlush() bool {
	firstCard := h.cards[0]

	for i, card := range h.cards {
		if card.GetSuit() != firstCard.GetSuit() {
			return false
		}
		if rankOrder[card.GetRank()] != rankOrder[firstCard.GetRank()]+i {
			return false
		}
	}

	h.handType = StraightFlush

	return true
}

func (h *hand) fourOfAKind() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 4 {
			h.handType = FourOfAKind
			return true
		}
	}

	return false
}

func (h *hand) fullHouse() bool {
	cnt := 0

	for _, v := range h.cardsRankCountMap {
		if v == 3 {
			cnt += 1
		}
		if v == 2 {
			cnt += 1
		}
	}

	if cnt == 2 {
		h.handType = FullHouse
		return true
	}

	return false
}

func (h *hand) flush() bool {
	for _, card := range h.cards {
		if !card.EqualsSuit(h.cards[0]) {
			return false
		}
	}

	h.handType = Flush
	return true
}

func (h *hand) straight() bool {
	for i, card := range h.cards {
		if card.RankDifference(h.cards[0]) != i {
			return false
		}
	}

	h.handType = Straight
	return true

}

func (h *hand) threeOfAKind() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 3 {
			h.handType = ThreeOfAKind
			return true
		}
	}

	return false
}

func (h *hand) twoPairs() bool {
	cnt := 0
	for _, v := range h.cardsRankCountMap {
		if v == 2 {
			cnt += 1
		}
	}

	if cnt == 2 {
		h.handType = TwoPairs
		return true
	}

	return false
}

func (h *hand) pair() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 2 {
			h.handType = Pair
			return true
		}
	}

	return false
}

func (h *hand) highCard() bool {
	h.handType = HighCard
	return true
}
