package strategy

import (
	"reflect"
	"testing"
)

func TestDeterministicStrategy_Run(t *testing.T) {
	tests := []struct {
		name string
		d    DeterministicStrategy
		want Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Run(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeterministicStrategy.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
