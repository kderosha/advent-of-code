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
	Id int
	rolledValues []int
	winningValues []int
	wonValues []int
	subtreeSize int
}

// Parses input and creates a new card
// example input:
// Card   1: 43 19 57 13 44 22 29 20 34 33 | 34 68 13 38 32 57 20 64 42  7 44 54 16 51 33 85 43 24 86 93 83 29 25 19 22
func NewCard(input string) *Card {

	// split the string by | character
	splitInput := strings.Split(input, " | ")
	// Input the second part of the split string
	winningValues := parseWinningValues(splitInput[1])
	idAndRolledValues := strings.Split(splitInput[0], ": ")
	id := parseId(idAndRolledValues[0])
	rolledValues := parseRolledValues(idAndRolledValues[1])

	return &Card{
		Id:id, 
		rolledValues: rolledValues,
		winningValues: winningValues,
		wonValues: make([]int, 0),
	}
}

func (c *Card) CalculateScore() (int) {
	score := 0
	// Sort each of the arrays holding the values
	slices.Sort(c.rolledValues)
	slices.Sort(c.winningValues)
	for _, rolledValue := range c.rolledValues {
		// Search for the rolled values in the winning values.
		// Binary search the winning values for the current value.
		wonValuePosition, found := slices.BinarySearch(c.winningValues, rolledValue)
		if found {
			c.wonValues = append(c.wonValues, c.winningValues[wonValuePosition])
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	slog.Info("Finished processing score", "won values", c.wonValues)

	return score
}

func (c *Card) SubtreeSize(cardMap map[int]*Card) int {
	if c.subtreeSize != 0 {
		return c.subtreeSize
	} else {
		if len(c.wonValues) == 0 {
			c.subtreeSize = 1
			return 1
		}
		score := 1
		for idx, _ := range c.wonValues {
			score += cardMap[c.Id + 1 + idx].SubtreeSize(cardMap)
		}
		c.subtreeSize = score
		return score
	}
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
	slog.Info("Finished processing winning values", "values", returnSlice)
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
			slog.Error("Error parsing rolled values", "valueString", valueString)
		}
		returnSlice = append(returnSlice, integerValue)
	}
	slog.Info("Finished processing rolled values", "values", returnSlice)
	return returnSlice
}