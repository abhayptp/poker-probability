package utils

import (
	"reflect"
	"testing"

	"github.com/abhayptp/poker-probability/models"
)

func TestCombineCardsList(t *testing.T) {
	type args struct {
		c [][]models.Card
	}
	tests := []struct {
		name string
		args args
		want []models.Card
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineCardsList(tt.args.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombineCardsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCard(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name     string
		args     args
		wantRank models.Rank
		wantSuit models.Suit
	}{
		{
			name: "success",
			args: args{
				c: "10H",
			},
			wantRank: "10",
			wantSuit: 'H',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateCard(tt.args.c)
			if got.GetSuit() != tt.wantSuit {
				t.Errorf("Suit = %c, want %c", got.GetSuit(), tt.wantSuit)
			}
			if got.GetRank() != tt.wantRank {
				t.Errorf("Rank = %s, want %s", got.GetRank(), tt.wantRank)
			}

		})
	}
}

func TestCreateCardsList(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name      string
		args      args
		wantRanks []models.Rank
		wantSuits []models.Suit
	}{
		{
			name: "success",
			args: args{
				c: "10D,1C,AH",
			},
			wantRanks: []models.Rank{
				models.Rank("10"), models.Rank("1"), models.Rank("A"),
			},
			wantSuits: []models.Suit{
				models.Suit('D'), models.Suit('C'), models.Suit('H'),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateCardsList(tt.args.c)
			for i, card := range got {
				if card.GetRank() != tt.wantRanks[i] {
					t.Errorf("%d card rank = %s, want %s", i+1, card.GetRank(), tt.wantRanks[i])
				}
				if card.GetSuit() != tt.wantSuits[i] {
					t.Errorf("%d card rank = %c, want %c", i+1, card.GetSuit(), tt.wantSuits[i])
				}
			}
		})
	}
}

func TestGetBestHand(t *testing.T) {
	hands := []models.Hand{
		models.NewHand(CreateCardsList("AH,AD,AC,1D,2D")),
		models.NewHand(CreateCardsList("1S,2S,3S,4S,5H")),
		models.NewHand(CreateCardsList("10S,JS,KS,QS,AS")),
		models.NewHand(CreateCardsList("1H,1H,2H,2H,3H")),
	}

	type args struct {
		hands []models.Hand
	}
	tests := []struct {
		name string
		args args
		want models.Hand
	}{
		{
			name: "success",
			args: args{
				hands: hands,
			},
			want: hands[2],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, hand := range tt.args.hands {
				hand.SetHandType()
			}

			if got := GetBestHand(tt.args.hands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBestHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankHands(t *testing.T) {
	type args struct {
		hands []models.Hand
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				hands: []models.Hand{
					models.NewHand(CreateCardsList("AH,AD,AC,1D,2D")),
					models.NewHand(CreateCardsList("1S,2S,3S,4S,5H")),
					models.NewHand(CreateCardsList("10S,JS,QS,KS,AS")),
					models.NewHand(CreateCardsList("1H,1H,2H,2H,3H")),
				},
			},
			want: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, hand := range tt.args.hands {
				hand.SetHandType()
			}
			if got := RankHands(tt.args.hands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RankHands() = %v, want %v", got, tt.want)
			}
		})
	}
}
