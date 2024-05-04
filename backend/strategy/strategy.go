package strategy

import "github.com/abhayptp/poker-probability/models"

type PlayerResult struct {
	Cards          models.DealtCards
	WinProbability float64
	TieProbability float64
}

type Result struct {
	PlayerResult []PlayerResult
}

type ProbabilityCalculator interface {
	Run() Result
}
