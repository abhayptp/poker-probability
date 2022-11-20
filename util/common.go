package util

import (
	"sort"

	"github.com/abhayptp/poker-probability/models"
)

// Return rank of each hand (lower number => higher rank)
func RankHands(hands []*models.Hand) []int {
	index := make([]int, len(hands))
	for i := 0; i < len(hands); i++ {
		index[i] = i
	}

	sort.Slice(index, func(i, j int) bool {
		return hands[index[i]].Compare(*hands[index[j]]) > 1
	})

	rank := make([]int, len(hands))
	rank[index[0]] = 0
	for i := 1; i < len(hands); i++ {
		if hands[index[i]].Compare(*hands[index[i-1]]) == 0 {
			rank[index[i]] = rank[index[i-1]]
		} else {
			rank[index[i]] = rank[index[i-1]] + 1
		}
	}
	return rank
}

// Get best hand
func GetBestHand(hands []*models.Hand) *models.Hand {
	bestHand := hands[0]

	for _, hand := range hands {
		if hand == nil {
			continue
		}

		if bestHand.Compare(*hand) == -1 {
			bestHand = hand
		}
	}

	return bestHand
}
