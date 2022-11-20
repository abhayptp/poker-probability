package main

import (
	"sync"
	"sync/atomic"

	"github.com/abhayptp/poker-probability/models"
	"github.com/abhayptp/poker-probability/util"
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

	go func() {
		defer wg.Done()
		playerCardsCopy := make([]models.DealtCards, len(a.playerCards))
		_ = copy(playerCardsCopy, a.playerCards)

		wins, draws := a.simulateGame(a.communityCards, playerCardsCopy, a.playersCount, a.simulationRounds)
		atomic.AddInt64(&winCount, wins)
		atomic.AddInt64(&drawCount, draws)
	}()

	wg.Wait()

	winProbability := float64(winCount*1.0) / float64(a.concurrency*a.simulationRounds)
	tieProbability := float64(drawCount * 1.0)

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
	for i := 0; i < simulationRounds; i++ {
		deck := models.NewDeck()
		for _, dealtCards := range playerCards {
			deck.RemoveCards(dealtCards.GetCards())
		}
		deck.RemoveCards(communityCards.GetOpenedCards())

		cards := deck.GetCardsFromTop(5 - communityCards.GetOpenedCardsCount())
		communityCards.SetUnopenedCards(cards)

		var playerHands []*models.Hand

		for i := 0; i < playersCount; i++ {
			cnt := 2 - len(playerCards[i].GetCards())
			if cnt != 0 {
				playerCards[i].SetUnknownDealtCards(deck.GetCardsFromTop(cnt)...)
			}
			sevenCards := append(playerCards[i].GetCards(), communityCards.GetAllCards()...)

			// Generate all possible hands with 7 cards
			possibleHands := make([]*models.Hand, 0, 21)
			for i := 0; i < 7; i++ {
				for j := i + 1; j < 7; j++ {
					hand := models.NewHand(
						append(
							sevenCards[:i],
							append(sevenCards[i+1:j], sevenCards[j+1:]...)...,
						),
					)
					possibleHands = append(possibleHands, hand)
				}
			}

			bestHand := util.GetBestHand(possibleHands)
			playerHands = append(playerHands, bestHand)
		}

		ranks := util.RankHands(playerHands)
		win, draw := false, false
		if ranks[0] == 0 {
			win = true
			for i := 1; i < playersCount; i++ {
				if ranks[i] == 0 {
					draw = true
				}
			}
		}
		if win {
			winCount += 1
		}
		if draw {
			drawCount += 1
		}

		// Adding back the removed cards
		deck.AddCards(communityCards.GetAllCards())
		for _, dealtCards := range playerCards {
			deck.AddCards(dealtCards.GetCards())
		}
	}
	return winCount, drawCount
}
