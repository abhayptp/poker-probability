package models

import (
	"reflect"
	"testing"
)

func TestNewDealtCards(t *testing.T) {
	type args struct {
		c []Card
	}
	tests := []struct {
		name string
		args args
		want *dealtCards
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDealtCards(tt.args.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDealtCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dealtCards_GetCards(t *testing.T) {
	type fields struct {
		cards []Card
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
			d := &dealtCards{
				cards: tt.fields.cards,
			}
			if got := d.GetCards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealtCards.GetCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dealtCards_SetUnknownDealtCards(t *testing.T) {
	type fields struct {
		cards []Card
	}
	type args struct {
		c []Card
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
			d := &dealtCards{
				cards: tt.fields.cards,
			}
			d.SetUnknownDealtCards(tt.args.c...)
		})
	}
}
