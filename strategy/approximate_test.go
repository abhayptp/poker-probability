package strategy

import (
	"testing"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/utils"
)

func TestApproximateStrategy_Run(t *testing.T) {
	tests := []struct {
		name string
		a    ApproximateStrategy
		want Result
	}{
		{
			name: "success",
			a: ApproximateStrategy{
				playerCards: []models.DealtCards{
					models.NewDealtCards(utils.CreateCardsList("AH,AC")...),
				},
				communityCards:   models.NewCommunityCards(),
				playersCount:     2,
				concurrency:      3,
				simulationRounds: 100000,
			},
			want: Result{
				winProbability: 0.9313,
				tieProbability: 0.0041,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Run()
			if got.tieProbability-tt.want.tieProbability > 1e-3 {
				t.Errorf("Difference in tie probability more than 1e-3, got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApproximateStrategy_simulateGame(t *testing.T) {
	type args struct {
		communityCards   models.CommunityCards
		playerCards      []models.DealtCards
		playersCount     int
		simulationRounds int
	}
	tests := []struct {
		name  string
		a     ApproximateStrategy
		args  args
		want  int64
		want1 int64
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.a.simulateGame(tt.args.communityCards, tt.args.playerCards, tt.args.playersCount, tt.args.simulationRounds)
			if got != tt.want {
				t.Errorf("ApproximateStrategy.simulateGame() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ApproximateStrategy.simulateGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
