package models

import (
	"reflect"
	"testing"
)

func TestNewCommunityCards(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want *communityCards
	}{
		{
			name: "success",
			args: args{
				cards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
			},
			want: &communityCards{
				openedCards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommunityCards(tt.args.cards...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommunityCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_communityCards_GetOpenedCards(t *testing.T) {
	type fields struct {
		openedCards   []Card
		unopenedCards []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []Card
	}{
		{
			name: "success",
			fields: fields{
				openedCards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
			},
			want: []Card{
				&card{
					rank: "10",
					suit: "H",
				},
				&card{
					rank: "A",
					suit: "C",
				},
				&card{
					rank: "1",
					suit: "D",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &communityCards{
				openedCards:   tt.fields.openedCards,
				unopenedCards: tt.fields.unopenedCards,
			}
			if got := c.GetOpenedCards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("communityCards.GetOpenedCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_communityCards_GetUnopenedCards(t *testing.T) {
	type fields struct {
		openedCards   []Card
		unopenedCards []Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []Card
	}{
		{
			name: "success",
			fields: fields{
				openedCards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
				unopenedCards: []Card{
					&card{
						rank: "2",
						suit: "H",
					},
					&card{
						rank: "5",
						suit: "D",
					},
				},
			},
			want: []Card{
				&card{
					rank: "2",
					suit: "H",
				},
				&card{
					rank: "5",
					suit: "D",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &communityCards{
				openedCards:   tt.fields.openedCards,
				unopenedCards: tt.fields.unopenedCards,
			}
			if got := c.GetUnopenedCards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("communityCards.GetUnopenedCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_communityCards_SetUnopenedCards(t *testing.T) {
	type fields struct {
		openedCards   []Card
		unopenedCards []Card
	}
	type args struct {
		cards []Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *communityCards
	}{
		{
			name: "success",
			fields: fields{
				openedCards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
				unopenedCards: []Card{
					&card{
						rank: "2",
						suit: "H",
					},
					&card{
						rank: "5",
						suit: "D",
					},
				},
			},
			args: args{
				[]Card{
					&card{
						rank: "2",
						suit: "H",
					},
					&card{
						rank: "5",
						suit: "D",
					},
				},
			},
			want: &communityCards{
				openedCards: []Card{
					&card{
						rank: "10",
						suit: "H",
					},
					&card{
						rank: "A",
						suit: "C",
					},
					&card{
						rank: "1",
						suit: "D",
					},
				},
				unopenedCards: []Card{
					&card{
						rank: "2",
						suit: "H",
					},
					&card{
						rank: "5",
						suit: "D",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &communityCards{
				openedCards:   tt.fields.openedCards,
				unopenedCards: tt.fields.unopenedCards,
			}
			c.SetUnopenedCards(tt.args.cards)
		})
	}
}

func Test_communityCards_GetAllCards(t *testing.T) {
	type fields struct {
		openedCards   []Card
		unopenedCards []Card
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
			c := &communityCards{
				openedCards:   tt.fields.openedCards,
				unopenedCards: tt.fields.unopenedCards,
			}
			if got := c.GetAllCards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("communityCards.GetAllCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
