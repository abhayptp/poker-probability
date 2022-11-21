package models

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	type args struct {
		rank Rank
		suit Suit
	}
	tests := []struct {
		name    string
		args    args
		want    *card
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			want: &card{
				rank: "10",
				suit: 'C',
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				rank: Rank("11"),
				suit: Suit('C'),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.args.rank, tt.args.suit)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_GetRank(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	tests := []struct {
		name   string
		fields fields
		want   Rank
	}{
		{
			name: "10C should return 10 rank",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			want: Rank("10"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card.GetRank(); got != tt.want {
				t.Errorf("card.GetRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_GetSuit(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	tests := []struct {
		name   string
		fields fields
		want   Suit
	}{
		{
			name: "10C should return 'C' suit",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			want: Suit('C'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card.GetSuit(); got != tt.want {
				t.Errorf("card.GetSuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_Valid(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "rank not valid",
			fields: fields{
				rank: Rank("11"),
				suit: Suit('C'),
			},
			want: false,
		},
		{
			name: "suit not valid",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('A'),
			},
			want: false,
		},
		{
			name: "valid",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card.Valid(); got != tt.want {
				t.Errorf("card.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_LessThan(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "1S should be less than JD",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('S'),
			},
			args: args{
				card2: &card{
					rank: "J",
					suit: 'D',
				},
			},
			want: true,
		},
		{
			name: "AH should not be less than 10C",
			fields: fields{
				rank: Rank("A"),
				suit: Suit('H'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'C',
				},
			},
			want: false,
		},
		{
			name: "10C should not be less than 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.LessThan(tt.args.card2); got != tt.want {
				t.Errorf("card.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_GreaterThan(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "JD should be greater than 1S",
			fields: fields{
				rank: Rank("J"),
				suit: Suit('D'),
			},
			args: args{
				card2: &card{
					rank: "1",
					suit: 'S',
				},
			},
			want: true,
		},
		{
			name: "10C should not be greater than AH",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "A",
					suit: 'H',
				},
			},
			want: false,
		},
		{
			name: "10C should not be greater than 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.GreaterThan(tt.args.card2); got != tt.want {
				t.Errorf("card.GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_Equals(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "10C should not be equal to 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: false,
		},
		{
			name: "1C should not be equal to 10C",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'C',
				},
			},
			want: false,
		},
		{
			name: "10C should be equal to 10C",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'C',
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.Equals(tt.args.card2); got != tt.want {
				t.Errorf("card.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_EqualsSuit(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "10C should not be equal to 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: false,
		},
		{
			name: "1C should be equal to 10C",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "1",
					suit: 'C',
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.EqualsSuit(tt.args.card2); got != tt.want {
				t.Errorf("card.EqualsSuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_EqualsRank(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "10C should not be equal to 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: true,
		},
		{
			name: "1C should be equal to 10C",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'C',
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.EqualsRank(tt.args.card2); got != tt.want {
				t.Errorf("card.EqualsRank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_RankDifference(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	type args struct {
		card2 Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "10C should not be equal to 10H",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'H',
				},
			},
			want: 0,
		},
		{
			name: "1C should be equal to 10C",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('C'),
			},
			args: args{
				card2: &card{
					rank: "10",
					suit: 'C',
				},
			},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card1 := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card1.RankDifference(tt.args.card2); got != tt.want {
				t.Errorf("card.RankDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_String(t *testing.T) {
	type fields struct {
		rank Rank
		suit Suit
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "equal",
			fields: fields{
				rank: Rank("10"),
				suit: Suit('C'),
			},
			want: "10C",
		},
		{
			name: "not equal",
			fields: fields{
				rank: Rank("1"),
				suit: Suit('C'),
			},
			want: "1C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &card{
				rank: tt.fields.rank,
				suit: tt.fields.suit,
			}
			if got := card.String(); got != tt.want {
				t.Errorf("card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
