package game

import (
	"testing"
)

func TestNewRoundParsing(t *testing.T) {
	// expected colors array is [red, blue, green]
	tests := []struct{
		testRound string
		expectedColors []int
	}{
		//case 1
		{
			testRound: "20 red, 80 blue, 99 green",
			expectedColors: []int{20, 80, 99},
		},
		{
			testRound: "10 red, 20 green",
			expectedColors: []int{10, 0, 20},
		},
		{
			testRound: "",
			expectedColors: []int{0, 0, 0},
		},
		{
			testRound: "10 blue",
			expectedColors: []int{0, 10, 0},
		},
	}

	for _, testCase := range tests {
		round := NewRound(testCase.testRound)
		if round.colors["red"] != testCase.expectedColors[0] {
			t.Fatalf("red color unexpected, wanted %d but got %d", testCase.expectedColors[0], round.colors["red"] )
		}
		if round.colors["blue"] != testCase.expectedColors[1] {
			t.Fatalf("blue color unexpected, wanted %d but got %d", testCase.expectedColors[0], round.colors["blue"] )
		}
		if round.colors["green"] != testCase.expectedColors[2] {
			t.Fatalf("green color unexpected, wanted %d but got %d", testCase.expectedColors[0], round.colors["green"] )
		}
	}
}