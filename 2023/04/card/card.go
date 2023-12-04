package card

import (
	"strings"
	"strconv"
	"log/slog"
	"regexp"
	"slices"
)

var matcher *regexp.Regexp = regexp.MustCompile(`\d+`)

type Card struct {
	id int
	rolledValues []int
	winningValues []int
}

// Parses input and creates a new card
// example input:
// Card   1: 43 19 57 13 44 22 29 20 34 33 | 34 68 13 38 32 57 20 64 42  7 44 54 16 51 33 85 43 24 86 93 83 29 25 19 22
func NewCard(input string) Card {

	// split the string by | character
	splitInput := strings.Split(input, " | ")
	// Input the second part of the split string
	winningValues := parseWinningValues(splitInput[1])
	idAndRolledValues := strings.Split(splitInput[0], ": ")
	id := parseId(idAndRolledValues[0])
	rolledValues := parseRolledValues(idAndRolledValues[1])

	return Card{
		id, rolledValues, winningValues,
	}
}

func (c Card) CalculateScore() int {
	score := 0
	// Sort each of the arrays holding the values
	slices.Sort(c.rolledValues)
	slices.Sort(c.winningValues)
	for _, rolledValue := range c.rolledValues {
		// Search for the rolled values in the winning values.
		// Binary search the winning values for the current value.
		if _, found := slices.BinarySearch(c.winningValues, rolledValue); found  && score == 0{
			// If found increase score by 1
			score = 1
		} else if found {
			score *= 2
		}
	}
	// Calculate score by 2^matches
	return score
}

func parseWinningValues(input string) []int {
	var returnSlice []int = make([]int, 0)
	slog.Info("base input string", "input", input)
	matches := matcher.FindAllString(input, -1)
	slog.Info("Matched digits in input", "matches", matches)
	for _, valueString := range matches {
		if valueString != " " {
			integerValue, err := strconv.Atoi(valueString)
			if err != nil {
				slog.Error("Error parsing winning value", "valueString", valueString)
			}
			returnSlice = append(returnSlice, integerValue)
		}
	}
	return returnSlice
}


func parseId(input string) int {
	items := strings.Split(input, " ")
	id, err := strconv.Atoi(items[len(items) - 1])
	if err != nil {
		slog.Error("Error parsing card id", "input", input)
	}
	return id
}

func parseRolledValues(input string) []int{
	var returnSlice []int = make([]int, 0)
	values := strings.Split(input, " ")
	for _, valueString := range values {
		integerValue, err := strconv.Atoi(valueString)
		if err != nil {
			slog.Error("Error parsing winning value", "valueString", valueString)
		}
		returnSlice = append(returnSlice, integerValue)
	}
	return returnSlice
}