package card

import (
	"slices"
	"testing"
)

func TestCardProcessingFromStringInput(t *testing.T) {
	testCases := []struct{
		input string
		expectedId int
		cardValues []int
		winningValues []int
	}{{
		input: "Card   1: 43 19 57 13 44 22 29 20 34 33 | 34 68 13 38 32 57 20 64 42  7 44 54 16 51 33 85 43 24 86 93 83 29 25 19 22",
		expectedId: 1,
		cardValues: []int{43, 19, 57, 13, 44, 22, 29, 20, 34, 33},
		winningValues: []int{34, 68, 13, 38, 32, 57, 20, 64, 42, 7, 44, 54, 16, 51, 33, 85, 43, 24, 86, 93, 83, 29, 25, 19, 22},
	}}

	// Create a new Card object
	for _, testCase := range testCases{
		card := NewCard(testCase.input)
		slices.Sort(testCase.cardValues)
		slices.Sort(testCase.winningValues)
		if !testArrayEquality(card.rolledValues, testCase.cardValues) {
			t.Fatalf("Cards rolled values are not expected, want %v got %v", testCase.cardValues, card.rolledValues)
		}
		if card.id != testCase.expectedId {
			t.Fatalf("Card id was not expected, want %d got %d", testCase.expectedId, card.id)
		}
		if !testArrayEquality(card.winningValues, testCase.winningValues) {
			t.Fatalf("card winning values were not expected, want %v got %v", testCase.winningValues, card.winningValues)
		}
	}
}

func TestCalculateCardScore(t *testing.T) {
	testCases := []struct{
		input string
		expectedScore int
	}{{
		input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		expectedScore: 8,
	}}

	for _, testCase := range testCases {
		card := NewCard(testCase.input)
		calculatedScore := card.CalculateScore()
		if calculatedScore != testCase.expectedScore {
			t.Fatalf("Calculated score is unexpected, want %f got %f", testCase.expectedScore, calculatedScore)
		}
	}
}

// Sorts and compares two arrays to test if they are equal
func testArrayEquality(array []int, comparedTo []int) bool {
	// Sort both
	slices.Sort(array)
	slices.Sort(comparedTo)
	if len(array) != len(comparedTo) {
		return false
	}
	for x, _ := range array {
		if array[x] != comparedTo[x] {
			return false
		}
	}

	return true;

}