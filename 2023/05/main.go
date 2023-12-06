package main

import (
	"github.com/kderosha/advent-of-code/2023/05/transformations"
	"regexp"
	"log/slog"
	"os"
	"fmt"
	"bytes"
	"strconv"
	"math"
)

var digits *regexp.Regexp = regexp.MustCompile(`\d+`)

func main(){
	puzzleInput, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}

	// Split the puzzle_input.txt into byte array
	slog.Info("puzzle_input.txt is read", "bytes", fmt.Sprintf("%+v", puzzleInput))
	firstNewLineIndex := bytes.IndexAny(puzzleInput, "\n")
	slog.Info("Slice first line of input", "line", string(puzzleInput[:firstNewLineIndex]))
	seeds := parseSeeds(puzzleInput[:firstNewLineIndex])
	var transformedSeedValues []int = make([]int, 0, len(seeds))
	// For the rest of the puzzle parse into transformation chain
	slog.Info("Rest of the puzzle", "puzzle", string(puzzleInput[firstNewLineIndex+2:]))


	// Take the rest of the lines and 
	tc := transformations.NewTransformationChain(puzzleInput[firstNewLineIndex + 2:])
	for _, seed := range seeds {
		transformedSeedValues = append(transformedSeedValues, tc.Transform(seed, transformations.Source("seed")))
	}
	slog.Info("Done transforming seed values", "transformed seeds", transformedSeedValues)
	p1Answer := findMinValueInArray(transformedSeedValues)
	slog.Info("Smallest transformed value calculated", "value", p1Answer)
	slog.Info("Process part 2")
	var p2MinValues []int = make([]int, 0)
	for x := 0; x + 1 < len(seeds); {
		p2MinValues = append(p2MinValues, findSmallestLocationInRange(tc, seeds[x], seeds[x + 1]))
		x = x + 2
	}
	part2Answer := findMinValueInArray(p2MinValues)
	slog.Info("Part 2 is done processing", "p2MinValues", p2MinValues, "answer", part2Answer)
}

func findSmallestLocationInRange(tc transformations.TransformationChain, startValue int, rangeSize int) int {
	// Make an array of integers with the range
	var values []int = make([]int, 0, rangeSize)
	endValue := startValue + rangeSize - 1
	slog.Info("Finding the smallest number by transforming every number in a given range", "startValue", startValue, "endValue", endValue)
	for seedValue := startValue; seedValue <= endValue; seedValue++{
		values = append(values, tc.Transform(seedValue, transformations.Source("seed")))
	}
	return findMinValueInArray(values)
}

// Finds the minimum value in an array of integers
func findMinValueInArray(input []int) int {
	smallestTransformedValue := math.MaxInt
	for _, value := range input {
		if value < smallestTransformedValue {
			smallestTransformedValue = value
		}
	}
	return smallestTransformedValue
}

func parseSeeds(input []byte) []int {
	var seeds []int = make([]int, 0)
	digitByteArrays := digits.FindAll(input, -1)
	for _, digitByteArray := range digitByteArrays {
		slog.Info("Digits parsed", "bytes", digitByteArray)
		digit, err := strconv.Atoi(string(digitByteArray))
		if err != nil {
			slog.Error("Error parsing digit for seed")
		}
		seeds = append(seeds, digit)
	}
	return seeds
}