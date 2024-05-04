package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/strategy"
	"github.com/abhayptp/poker-probability/utils"
)

func parseInput(playerCards []string, communityCards string, playersCount int, simulationRounds int) (strategy.ProbabilityCalculator, error) {
	pCards := make([]models.DealtCards, 0)

	for _, cards := range playerCards {
		cardList, err := utils.CreateCardsListOrError(cards)
		if err != nil {
			return nil, err
		}
		pCards = append(pCards, models.NewDealtCards(cardList...))
	}

	cardList, err := utils.CreateCardsListOrError(communityCards)
	if err != nil {
		return nil, err
	}
	cCards := models.NewCommunityCards(cardList...)

	return strategy.NewApproximate(
		cCards, pCards, playersCount, simulationRounds, 5,
	), nil
}

func printResult(res strategy.Result) {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, string(colorRed), "Player Cards", "\t", string(colorGreen), "Win Probability", "\t", string(colorYellow), "Tie Probability")
	for _, p := range res.PlayerResult {
		fmt.Fprintln(w, string(colorRed), p.Cards, "\t", string(colorGreen), p.WinProbability, "\t", string(colorYellow), p.TieProbability)
	}
	w.Flush()
}

func main() {
	communityCards := flag.String("community_cards", "", "string for opened community cards")
	playersCount := flag.Int("players_count", 2, "count of players in the game")
	simulationRounds := flag.Int("simulation_rounds", 100000, "number of simulation rounds")

	flag.Parse()
	playerCards := flag.Args()

	calculator, err := parseInput(playerCards, *communityCards, *playersCount, *simulationRounds)
	if err != nil {
		panic(err)
	}

	res := calculator.Run()
	printResult(res)
}
