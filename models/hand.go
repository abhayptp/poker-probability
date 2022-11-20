package models

import "sort"

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

type Hand struct {
	Cards []*Card

	cardsRankCountMap    map[*Rank]int8
	handType             HandType
	tieBreakerCardsOrder [5]*Card
}

func NewHand(cards []*Card) *Hand {
	assert(len(cards) == 5)

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].LessThan(cards[j])
	})

	// Setting cardsRankCountMap
	var m map[*Rank]int8

	for _, card := range cards {
		m[card.GetRank()] += 1
	}

	// Setting tieBreakerCardsOrder
	tieBreakerCardsOrder := cards
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

	return &Hand{
		Cards: cards,

		cardsRankCountMap: m,
		// Default hand type is HighCard
		handType: Default,
	}
}

func assert(b bool) {
	panic("unimplemented")
}

// Return value:
// 	1  if h1 > h2
//	0 if h1 == h2
//	-1  if h1 < h2
func (h1 Hand) Compare(h2 Hand) int {
	if h1.handType == Default || h2.handType == Default {
		panic("hand type not set")
	}

	if h1.handType > h2.handType {
		return 1
	}

	if h1.handType < h2.handType {
		return -1
	}

	for i := 0; i < 5; i++ {
		if h1.tieBreakerCardsOrder[i].GreaterThan(h2.tieBreakerCardsOrder[i]) {
			return 1
		}

		if h1.tieBreakerCardsOrder[i].LessThan(h2.tieBreakerCardsOrder[i]) {
			return -1
		}
	}

	return 0
}

func (h Hand) SetHandType() {
	fList := [10]func() bool{
		h.RoyalFlush,
		h.StraightFlush,
		h.FourOfAKind,
		h.FullHouse,
		h.Flush,
		h.Straight,
		h.ThreeOfAKind,
		h.TwoPairs,
		h.Pair,
		h.HighCard,
	}

	for _, f := range fList {
		if f() {
			return
		}
	}
}

func (h Hand) RoyalFlush() bool {
	if *h.Cards[0].GetRank() != "10" {
		return false
	}

	if h.StraightFlush() {
		h.handType = RoyalFlush
		return true
	}

	return false
}

func (h Hand) StraightFlush() bool {
	firstCard := h.Cards[0]

	for i, card := range h.Cards {
		if *card.GetSuit() != *firstCard.GetSuit() {
			return false
		}
		if rankOrder[*card.GetRank()] != rankOrder[*firstCard.GetRank()]+i {
			return false
		}
	}

	h.handType = StraightFlush

	return true
}

func (h Hand) FourOfAKind() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 4 {
			h.handType = FourOfAKind
			return true
		}
	}

	return false
}

func (h Hand) FullHouse() bool {
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
		return true
	}

	return false
}

func (h Hand) Flush() bool {
	for _, card := range h.Cards {
		if !card.EqualsSuit(h.Cards[0]) {
			return false
		}
	}

	return true
}

func (h Hand) Straight() bool {
	for i, card := range h.Cards {
		if card.RankDifference(h.Cards[0]) != i {
			return false
		}
	}

	return true

}

func (h Hand) ThreeOfAKind() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 3 {
			return true
		}
	}

	return false
}

func (h Hand) TwoPairs() bool {
	cnt := 0
	for _, v := range h.cardsRankCountMap {
		if v == 2 {
			cnt += 1
		}
	}

	if cnt == 2 {
		return true
	}

	return false
}

func (h Hand) Pair() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 2 {
			return true
		}
	}

	return false
}

func (h Hand) HighCard() bool {
	return true
}
