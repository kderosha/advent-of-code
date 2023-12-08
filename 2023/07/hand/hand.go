package hand

import (
	"fmt"
	"slices"
)

const (
	FiveOfAKind  = 999
	FourOfAKind  = 800
	FullHouse    = 700
	ThreeOfAKind = 600
	TwoPair      = 500
	Pair         = 400
	HighCard     = 0
)

type Hand struct {
	Cards    []Card
	handType int
	bid      int
}

// Create a new hand
func NewHand(cardsString string, bid int) *Hand {
	cards := make([]Card, 0)
	for _, card := range cardsString {
		cards = append(cards, NewCard(card))
	}

	return &Hand{cards, calculateHandType(cards), bid}
}

func (h *Hand) String() string {
	return fmt.Sprintf("cards: %+v, bid: %d, handRank: %d", h.Cards, h.bid, h.handType)
}

func (h *Hand) Bid() int {
	return h.bid
}

func (h *Hand) Type() int {
	return h.handType
}

func calculateHandType(cards []Card) int {
	// bucket the ranks of each card.
	rankBuckets := make([]int, 13)
	for _, card := range cards {
		// increment the rank of that card
		rankBuckets[card.rank]++
	}
	// Evaluate each type of special rank
	// Five of a kind
	pairs := 0
	// For pair calcaulation diregard the J wildcards
	for _, rankBucket := range rankBuckets[1:] {
		if rankBucket == 2 {
			pairs++
		}
	}
	// 5 of a kind
	if slices.Contains(rankBuckets[1:], 5) || rankBuckets[0] == 5 {
		return FiveOfAKind
	}
	// 4 of a kind
	if slices.Contains(rankBuckets[1:], 4) {
		// If there is 1 wildcard that means there is actually 5 of a kind
		if rankBuckets[0] == 1 {
			return FiveOfAKind
		}
		return FourOfAKind
	}
	// 3 of a kind processing
	if slices.Contains(rankBuckets[1:], 3) {
		// if hand contains wild cards
		// Wild cards and 3 of a kind rules
		if rankBuckets[0] == 2 {
			// 3 of a kind with 2 wild cards becomes 5 of a kind
			return FiveOfAKind
		} else if rankBuckets[0] == 1 {
			// 3 of a kind and 1 wildcard becomes a 4 of a kind
			return FourOfAKind
		} else {
			// Full house with no wild cards
			if slices.Contains(rankBuckets[1:], 2) {
				return FullHouse
			}
			return ThreeOfAKind
		}
	}

	if pairs == 2 {
		// Wild card cases
		if rankBuckets[0] == 1 {
			return FullHouse
		}
		return TwoPair
	}
	// one pair
	if pairs == 1 {
		if rankBuckets[0] == 3 {
			return FiveOfAKind
		}
		if rankBuckets[0] == 2 {
			return FourOfAKind
		}
		if rankBuckets[0] == 1 {
			return ThreeOfAKind
		}
		return Pair
	}
	// Rest of the pure wild card cases
	if rankBuckets[0] == 4 {
		return FiveOfAKind
	}
	if rankBuckets[0] == 3 {
		return FourOfAKind
	}
	if rankBuckets[0] == 2 {
		return ThreeOfAKind
	}
	if rankBuckets[0] == 1 {
		return Pair
	}
	return HighCard
}

type Card struct {
	rank int
}

func (c *Card) Rank() int {
	return c.rank
}

// Form a new card
func NewCard(cardSymbol rune) Card {
	rank := 0
	switch cardSymbol {
	case 'J':
		rank = 0
	case '2':
		rank = 1
	case '3':
		rank = 2
	case '4':
		rank = 3
	case '5':
		rank = 4
	case '6':
		rank = 5
	case '7':
		rank = 6
	case '8':
		rank = 7
	case '9':
		rank = 8
	case 'T':
		rank = 9
	case 'Q':
		rank = 10
	case 'K':
		rank = 11
	case 'A':
		rank = 12
	default:
		panic(fmt.Errorf("error determining card rank for symbol: %s", string(cardSymbol)))
	}
	return Card{rank}
}
