package models

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	tests := []struct {
		name string
		want *deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_RemoveCards(t *testing.T) {
	type fields struct {
		cards []Card
	}
	type args struct {
		cards []Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deck{
				cards: tt.fields.cards,
			}
			d.RemoveCards(tt.args.cards)
		})
	}
}

func Test_deck_AddCards(t *testing.T) {
	type fields struct {
		cards []Card
	}
	type args struct {
		cards []Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deck{
				cards: tt.fields.cards,
			}
			d.AddCards(tt.args.cards)
		})
	}
}

func Test_deck_Burn(t *testing.T) {
	type fields struct {
		cards []Card
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deck{
				cards: tt.fields.cards,
			}
			d.Burn()
		})
	}
}

func Test_deck_PopCardsFromTop(t *testing.T) {
	type fields struct {
		cards []Card
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Card
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
			args: args{
				count: 2,
			},
			want: []Card{
				&card{
					rank: "1",
					suit: "S",
				},
				&card{
					rank: "2",
					suit: "S",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &deck{
				cards: tt.fields.cards,
			}
			got := d.PopCardsFromTop(tt.args.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deck.GetCardsFromTop() = %v, want %v", got, tt.want)
				return
			}
			for _, card := range got {
				if d.Contains(card) {
					t.Errorf("deck should not contain popped cards")
				}
			}
		})
	}
}
