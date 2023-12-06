package race

import (
	"testing"

	"github.com/kderosha/advent-of-code/testutils"
)

func TestRaceDistanceCalculations(t *testing.T) {
	testCases := []struct {
		buttonHoldTime           int
		raceTime                 int
		expectedDistanceTraveled int
	}{{
		buttonHoldTime:           5,
		raceTime:                 10,
		expectedDistanceTraveled: 25,
	}, {
		buttonHoldTime:           7,
		raceTime:                 10,
		expectedDistanceTraveled: 21,
	}, {
		buttonHoldTime:           1,
		raceTime:                 10,
		expectedDistanceTraveled: 9,
	}}

	for _, testCase := range testCases {
		distanceCalculated := calculateDistance(testCase.raceTime, testCase.buttonHoldTime)
		if distanceCalculated != testCase.expectedDistanceTraveled {
			t.Fatalf("Unexpected calculated distance, wanted %d got %d", testCase.expectedDistanceTraveled, testCase)
		}
	}
}

func TestRacePossibleWins(t *testing.T) {
	testCases := []struct {
		time                             int
		distanceRecord                   int
		expectedCalculatedDistanceValues []int
	}{{
		time:                             5,
		distanceRecord:                   1,
		expectedCalculatedDistanceValues: []int{4, 4, 6, 6},
	}, {
		time:                             20,
		distanceRecord:                   1,
		expectedCalculatedDistanceValues: []int{19, 19, 36, 36, 51, 51, 64, 64, 75, 75, 84, 84, 91, 91, 96, 96, 99, 99, 100},
	}}

	for _, testCase := range testCases {
		br := &BoatRace{
			time:           testCase.time,
			distanceRecord: testCase.distanceRecord,
		}
		possibleWaysToWin := br.PossibleWaysToWin()
		if !testutils.ArraysAreEqual(possibleWaysToWin, testCase.expectedCalculatedDistanceValues) {
			t.Fatalf("unexpected possible ways to win %+v, expected %+v", possibleWaysToWin, testCase.expectedCalculatedDistanceValues)
		}
	}

}
