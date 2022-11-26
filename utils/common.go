package utils

import (
	"sort"
	"strings"

	"github.com/abhayptp/poker-probability/models"
)

func CombineCardsList(c ...[]models.Card) []models.Card {
	cardsList := make([]models.Card, 0)
	for _, cl := range c {
		cardsList = append(cardsList, cl...)
	}
	return cardsList
}

func CreateCard(c string) models.Card {
	suit := c[len(c)-1]
	rank := c[:len(c)-1]

	card, _ := models.NewCard(models.Rank(rank), models.Suit(suit))

	return card
}

func CreateCardsList(c string) []models.Card {
	playersCardStr := strings.Split(c, ",")
	playerCards := make([]models.Card, 0)
	for _, playerCardStr := range playersCardStr {
		playerCards = append(playerCards, CreateCard(playerCardStr))
	}
	return playerCards
}

func CreateCardsListOrError(c string) ([]models.Card, error) {
	if len(c) == 0 {
		return []models.Card{}, nil
	}
	playersCardStr := strings.Split(c, ",")
	playerCards := make([]models.Card, 0)
	for _, c := range playersCardStr {
		suit := c[len(c)-1]
		rank := c[:len(c)-1]
		card, err := models.NewCard(models.Rank(rank), models.Suit(suit))
		if err != nil {
			return nil, err
		}
		playerCards = append(playerCards, card)
	}
	return playerCards, nil
}

func GetBestHand(hands []models.Hand) models.Hand {
	bestHand := hands[0]

	for _, hand := range hands {
		if hand == nil {
			continue
		}

		if bestHand.Compare(hand) == -1 {
			bestHand = hand
		}
	}

	return bestHand
}

func RankHands(hands []models.Hand) []int {
	index := make([]int, len(hands))
	for i := 0; i < len(hands); i++ {
		index[i] = i
	}

	sort.Slice(index, func(i, j int) bool {
		return hands[index[i]].Compare(hands[index[j]]) == 1
	})

	rank := make([]int, len(hands))
	// lower number => higher rank
	rank[index[0]] = 0
	for i := 1; i < len(hands); i++ {
		if hands[index[i]].Compare(hands[index[i-1]]) == 0 {
			rank[index[i]] = rank[index[i-1]]
		} else {
			rank[index[i]] = rank[index[i-1]] + 1
		}
	}
	return rank
}
