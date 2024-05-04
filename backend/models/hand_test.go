package models

import (
	"reflect"
	"testing"
)

func TestNewHand(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want *hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHand(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_Compare(t *testing.T) {
	type fields struct {
		Cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	type args struct {
		h2 Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h1 := &hand{
				cards:                tt.fields.Cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h1.Compare(tt.args.h2); got != tt.want {
				t.Errorf("hand.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_SetHandType(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			h.SetHandType()
		})
	}
}

func Test_hand_GetHandType(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   HandType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.GetHandType(); got != tt.want {
				t.Errorf("hand.GetHandType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_GetTieBreakerCardsOrder(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []Card
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.GetTieBreakerCardsOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hand.GetTieBreakerCardsOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_String(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				cards: []Card{
					&card{
						rank: "1",
						suit: "S",
					},
					&card{
						rank: "2",
						suit: "S",
					},
					&card{
						rank: "3",
						suit: "H",
					},
					&card{
						rank: "4",
						suit: "S",
					},
					&card{
						rank: "10",
						suit: "C",
					},
				},
			},
			want: "1S,2S,3H,4S,10C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("hand.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_royalFlush(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.royalFlush(); got != tt.want {
				t.Errorf("hand.royalFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_straightFlush(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.straightFlush(); got != tt.want {
				t.Errorf("hand.straightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_fourOfAKind(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.fourOfAKind(); got != tt.want {
				t.Errorf("hand.fourOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_fullHouse(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.fullHouse(); got != tt.want {
				t.Errorf("hand.fullHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_flush(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.flush(); got != tt.want {
				t.Errorf("hand.flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_straight(t *testing.T) {
	tests := []struct {
		name  string
		cards []Card
		want  bool
	}{
		{
			name: "success",
			cards: []Card{
				&card{
					rank: "1",
					suit: "S",
				},
				&card{
					rank: "2",
					suit: "S",
				},
				&card{
					rank: "3",
					suit: "S",
				},
				&card{
					rank: "4",
					suit: "S",
				},
				&card{
					rank: "5",
					suit: "S",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHand(tt.cards)
			if got := h.straight(); got != tt.want {
				t.Errorf("hand.straight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_threeOfAKind(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.threeOfAKind(); got != tt.want {
				t.Errorf("hand.threeOfAKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_twoPairs(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.twoPairs(); got != tt.want {
				t.Errorf("hand.twoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_pair(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.pair(); got != tt.want {
				t.Errorf("hand.pair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hand_highCard(t *testing.T) {
	type fields struct {
		cards                []Card
		cardsRankCountMap    map[Rank]int8
		handType             HandType
		tieBreakerCardsOrder []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards:                tt.fields.cards,
				cardsRankCountMap:    tt.fields.cardsRankCountMap,
				handType:             tt.fields.handType,
				tieBreakerCardsOrder: tt.fields.tieBreakerCardsOrder,
			}
			if got := h.highCard(); got != tt.want {
				t.Errorf("hand.highCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
