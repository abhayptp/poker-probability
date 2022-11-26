package strategy

import (
	"reflect"
	"sync"
	"testing"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/utils"
)

func TestApproximate_Run(t *testing.T) {
	tests := []struct {
		name string
		a    *Approximate
		want Result
	}{
		{
			name: "success",
			a: &Approximate{
				playerCards: []models.DealtCards{
					models.NewDealtCards(utils.CreateCardsList("AH,AC")...),
				},
				communityCards:   models.NewCommunityCards(),
				playersCount:     2,
				concurrency:      3,
				simulationRounds: 100000,
			},
			want: Result{
				PlayerResult: []PlayerResult{
					{
						Cards:          models.NewDealtCards(utils.CreateCardsList("AH,AC")...),
						WinProbability: 0.92,
						TieProbability: 0.0035,
					},
					{
						Cards:          models.NewDealtCards(),
						WinProbability: 0,
						TieProbability: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Run()
			for i, pRes := range got.PlayerResult {
				if pRes.TieProbability-tt.want.PlayerResult[i].TieProbability > 2e-2 {
					t.Errorf("Difference in tie probability more than 1e-3, got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestApproximate_simulateGame(t *testing.T) {
	type fields struct {
		Mutex            sync.Mutex
		communityCards   models.CommunityCards
		playerCards      []models.DealtCards
		playersCount     int
		simulationRounds int
		concurrency      int
	}
	type args struct {
		communityCards   models.CommunityCards
		playerCards      []models.DealtCards
		playersCount     int
		simulationRounds int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int64
		want1  []int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Approximate{
				Mutex:            tt.fields.Mutex,
				communityCards:   tt.fields.communityCards,
				playerCards:      tt.fields.playerCards,
				playersCount:     tt.fields.playersCount,
				simulationRounds: tt.fields.simulationRounds,
				concurrency:      tt.fields.concurrency,
			}
			got, got1 := a.simulateGame(tt.args.communityCards, tt.args.playerCards, tt.args.playersCount, tt.args.simulationRounds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approximate.simulateGame() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Approximate.simulateGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
