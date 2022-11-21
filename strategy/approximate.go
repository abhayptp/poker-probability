package strategy

import (
	"sync"
	"sync/atomic"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/utils"
)

type ApproximateStrategy struct {
	communityCards models.CommunityCards
	playerCards    []models.DealtCards

	playersCount     int
	simulationRounds int
	concurrency      int
}

func (a ApproximateStrategy) Run() Result {
	wg := new(sync.WaitGroup)
	wg.Add(a.concurrency)

	var winCount, drawCount int64 = 0, 0

	for i := 0; i < a.concurrency; i++ {
		go func() {
			defer wg.Done()
			playerCardsCopy := make([]models.DealtCards, len(a.playerCards))
			_ = copy(playerCardsCopy, a.playerCards)
			for i := len(playerCardsCopy); i < a.playersCount; i++ {
				playerCardsCopy = append(playerCardsCopy, models.NewDealtCards())
			}

			wins, draws := a.simulateGame(a.communityCards.Clone(), playerCardsCopy, a.playersCount, a.simulationRounds)
			atomic.AddInt64(&winCount, wins)
			atomic.AddInt64(&drawCount, draws)
		}()
	}

	wg.Wait()

	winProbability := float64(winCount*1.0) / float64(a.concurrency*a.simulationRounds)
	tieProbability := float64(drawCount*1.0) / float64(a.concurrency*a.simulationRounds)

	return Result{
		winProbability: winProbability,
		tieProbability: tieProbability,
	}
}

func (a ApproximateStrategy) simulateGame(communityCards models.CommunityCards,
	playerCards []models.DealtCards,
	playersCount int,
	simulationRounds int) (int64, int64) {
	var winCount int64
	var drawCount int64
	for round := 0; round < simulationRounds; round++ {
		deck := models.NewDeck()
		for _, dealtCards := range playerCards {
			deck.RemoveCards(dealtCards.GetCards())
		}
		deck.RemoveCards(communityCards.GetOpenedCards())

		cards := deck.PopCardsFromTop(5 - len(communityCards.GetOpenedCards()))
		communityCards.SetUnopenedCards(cards)

		var playerHands []models.Hand

		for pi := 0; pi < playersCount; pi++ {
			cnt := 2 - len(playerCards[pi].GetCards())
			if cnt != 0 {
				playerCards[pi].SetUnknownDealtCards(deck.PopCardsFromTop(cnt)...)
			}
			sevenCards := append(playerCards[pi].GetCards(), communityCards.GetAllCards()...)

			// Generate all possible hands with 7 cards
			possibleHands := make([]models.Hand, 0, 21)
			for i := 0; i < 7; i++ {
				for j := i + 1; j < 7; j++ {
					hand := models.NewHand(
						utils.CombineCardsList(
							sevenCards[:i],
							sevenCards[i+1:j],
							sevenCards[j+1:],
						),
					)
					hand.SetHandType()
					possibleHands = append(possibleHands, hand)
				}
			}

			bestHand := utils.GetBestHand(possibleHands)
			playerHands = append(playerHands, bestHand)
		}

		ranks := utils.RankHands(playerHands)
		win, draw := false, false
		if ranks[0] == 0 {
			win = true
			for i := 1; i < playersCount; i++ {
				if ranks[i] == 0 {
					draw = true
				}
			}
		}
		if win && !draw {
			winCount += 1
		}
		if draw {
			drawCount += 1
		}
	}
	return winCount, drawCount
}
