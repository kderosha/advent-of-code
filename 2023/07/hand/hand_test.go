package hand

import "testing"

func TestCardRank(t *testing.T) {
	testCases := []struct {
		cardSymbol   rune
		expectedRank int
	}{
		{
			cardSymbol:   'J',
			expectedRank: 0,
		},
		{
			cardSymbol:   '2',
			expectedRank: 1,
		},
		{
			cardSymbol:   '3',
			expectedRank: 2,
		},
		{
			cardSymbol:   '4',
			expectedRank: 3,
		},
		{
			cardSymbol:   '5',
			expectedRank: 4,
		},
		{
			cardSymbol:   '6',
			expectedRank: 5,
		},
		{
			cardSymbol:   '7',
			expectedRank: 6,
		},
		{
			cardSymbol:   '8',
			expectedRank: 7,
		},
		{
			cardSymbol:   '9',
			expectedRank: 8,
		},
		{
			cardSymbol:   'T',
			expectedRank: 9,
		},
		{
			cardSymbol:   'Q',
			expectedRank: 10,
		},
		{
			cardSymbol:   'K',
			expectedRank: 11,
		}, {
			cardSymbol:   'A',
			expectedRank: 12,
		},
	}

	for _, testCase := range testCases {
		card := NewCard(testCase.cardSymbol)
		if card.Rank() != testCase.expectedRank {
			t.Fatalf("Unexpected card rank, wanted %d got %d", testCase.expectedRank, card.Rank())
		}
	}
}

func TestHandType(t *testing.T) {
	testCases := []struct {
		hand             string // String of cards
		expectedHandType int
	}{
		{
			hand:             "22222",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "4444J",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "JJJJJ",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "222JJ",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "22223",
			expectedHandType: FourOfAKind,
		},
		{
			hand:             "222J3",
			expectedHandType: FourOfAKind,
		},
		{
			hand:             "22233",
			expectedHandType: FullHouse,
		},
		{
			hand:             "2233J",
			expectedHandType: FullHouse,
		},
		{
			hand:             "22234",
			expectedHandType: ThreeOfAKind,
		},
		{
			hand:             "22J34",
			expectedHandType: ThreeOfAKind,
		},
		{
			hand:             "22334",
			expectedHandType: TwoPair,
		},
		{
			hand:             "22345",
			expectedHandType: Pair,
		},
		{
			hand:             "22J45",
			expectedHandType: ThreeOfAKind,
		},
		{
			hand:             "22JJ3",
			expectedHandType: FourOfAKind,
		},
		{
			hand:             "22JJJ",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "2JJJJ",
			expectedHandType: FiveOfAKind,
		},
		{
			hand:             "23JJJ",
			expectedHandType: FourOfAKind,
		},
		{
			hand:             "234JJ",
			expectedHandType: ThreeOfAKind,
		},
		{
			hand:             "2345J",
			expectedHandType: Pair,
		},
		{
			hand:             "23456",
			expectedHandType: HighCard,
		},
	}

	for _, testCase := range testCases {
		hand := NewHand(testCase.hand, 0)
		if calculateHandType(hand.Cards) != testCase.expectedHandType {
			t.Fatalf("%+v Unexpected hand type, want %d got %d", testCase, testCase.expectedHandType, calculateHandType(hand.Cards))
		}
	}
}
