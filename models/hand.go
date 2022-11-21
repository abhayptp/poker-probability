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

type Hand interface {
	Compare(Hand) int
	SetHandType()
	GetHandType() HandType
	GetTieBreakerCardsOrder() [5]Card
}

type hand struct {
	Cards []Card

	cardsRankCountMap    map[*Rank]int8
	handType             HandType
	tieBreakerCardsOrder [5]Card
}

func NewHand(cards []Card) *hand {
	if len(cards) != 5 {
		panic("hand should have 5 cards")
	}

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

	return &hand{
		Cards: cards,

		cardsRankCountMap: m,
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

	if h1.handType > h2.GetHandType() {
		return 1
	}

	if h1.handType < h2.GetHandType() {
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

func (h *hand) GetTieBreakerCardsOrder() [5]Card {
	return h.tieBreakerCardsOrder
}

func (h *hand) royalFlush() bool {
	if *h.Cards[0].GetRank() != "10" {
		return false
	}

	if h.straightFlush() {
		h.handType = RoyalFlush
		return true
	}

	return false
}

func (h *hand) straightFlush() bool {
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

	return cnt == 2
}

func (h *hand) flush() bool {
	for _, card := range h.Cards {
		if !card.EqualsSuit(h.Cards[0]) {
			return false
		}
	}

	return true
}

func (h *hand) straight() bool {
	for i, card := range h.Cards {
		if card.RankDifference(h.Cards[0]) != i {
			return false
		}
	}

	return true

}

func (h *hand) threeOfAKind() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 3 {
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

	return cnt == 2
}

func (h *hand) pair() bool {
	for _, v := range h.cardsRankCountMap {
		if v == 2 {
			return true
		}
	}

	return false
}

func (h *hand) highCard() bool {
	return true
}
