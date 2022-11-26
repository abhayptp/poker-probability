package strategy

import (
	"sync"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/utils"
)

type Approximate struct {
	sync.Mutex
	communityCards models.CommunityCards
	playerCards    []models.DealtCards

	playersCount     int
	simulationRounds int
	concurrency      int
}

func NewApproximate(communityCards models.CommunityCards,
	playerCards []models.DealtCards,
	playersCount int,
	simulationRounds int,
	concurrency int) *Approximate {
	return &Approximate{
		communityCards:   communityCards,
		playerCards:      playerCards,
		playersCount:     playersCount,
		simulationRounds: simulationRounds,
		concurrency:      concurrency,
	}
}

func (a *Approximate) Run() Result {
	wg := new(sync.WaitGroup)
	wg.Add(a.concurrency)

	winCount := make([]int64, a.playersCount)
	drawCount := make([]int64, a.playersCount)

	for i := 0; i < a.concurrency; i++ {
		go func() {
			defer wg.Done()
			playerCardsCopy := make([]models.DealtCards, len(a.playerCards))
			_ = copy(playerCardsCopy, a.playerCards)
			for i := len(playerCardsCopy); i < a.playersCount; i++ {
				playerCardsCopy = append(playerCardsCopy, models.NewDealtCards())
			}

			wins, draws := a.simulateGame(a.communityCards.Clone(), playerCardsCopy, a.playersCount, a.simulationRounds)
			a.Lock()
			for i, v := range wins {
				winCount[i] += v
			}
			for i, v := range draws {
				drawCount[i] += v
			}
			a.Unlock()
		}()
	}

	wg.Wait()

	pRes := make([]PlayerResult, a.playersCount)

	for i := 0; i < a.playersCount; i++ {
		var pCards models.DealtCards
		if i < len(a.playerCards) {
			pCards = a.playerCards[i]
		} else {
			pCards = models.NewDealtCards()
		}
		pRes[i] = PlayerResult{
			Cards:          pCards,
			WinProbability: float64(winCount[i]) / float64(a.concurrency*a.simulationRounds),
			TieProbability: float64(drawCount[i]) / float64(a.concurrency*a.simulationRounds),
		}
	}

	return Result{
		PlayerResult: pRes,
	}
}

func (a *Approximate) simulateGame(communityCards models.CommunityCards,
	playerCards []models.DealtCards,
	playersCount int,
	simulationRounds int) ([]int64, []int64) {
	winCount := make([]int64, playersCount)
	drawCount := make([]int64, playersCount)
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
		winners := make([]int, 0)
		for i, v := range ranks {
			if v == 0 {
				winners = append(winners, i)
			}
		}
		if len(winners) == 1 {
			winCount[winners[0]] += 1
		} else {
			for _, w := range winners {
				drawCount[w] += 1
			}
		}
	}
	return winCount, drawCount
}
