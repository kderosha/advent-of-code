package hand

import "slices"

type Hand struct {
	cards []Card
	rank  int
	bid   int
}

// Create a new hand
func NewHand(cardsString string, bid int) *Hand {
	cards := make([]Card, 0)
	for _, card := range cardsString {
		cards = append(cards, NewCard(card))
	}
	calculateHandRank(cards)
	return &Hand{cards, bid}
}

func calculateHandRank(cards []Card) int {
	rankBuckets := make([]int, 13)
	// Evaluate each type of special rank
	// Five of a kind
	for _, card := range cards {
		// increment the rank of that card
		rankBuckets[card.rank]++
	}
	// 5 of a kind
	if slices.Contains(rankBuckets, 5) {
		return 999
	}
	// 4 of a kind
	if slices.Contains(rankBuckets, 4) {
		return 800
	}
	// full house
	if slices.Contains(rankBuckets, 3) && slices.Contains(rankBuckets, 2) {
		return 700
	}
	// 3 of a kind
	if slices.Contains(rankBuckets, 3) {
		return 600
	}
	// two pair

	// one pair

}

type Card struct {
	rank int
}

// Form a new card
func NewCard(cardSymbol rune) Card {
	rank := 0
	switch cardSymbol {
	case '2':
		rank = 0
	case '3':
		rank = 1
	case '4':
		rank = 2
	case '5':
		rank = 3
	case '6':
		rank = 4
	case '7':
		rank = 5
	case '8':
		rank = 6
	case '9':
		rank = 7
	case 'T':
		rank = 8
	case 'J':
		rank = 9
	case 'Q':
		rank = 10
	case 'K':
		rank = 11
	case 'A':
		rank = 12
	}
	return Card{rank}
}
