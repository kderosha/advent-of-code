package main

import (
	"bytes"
	"log/slog"
	"os"
	"regexp"
	"strconv"

	"github.com/kderosha/advent-of-code/2023/06/race"
)

var matcher *regexp.Regexp = regexp.MustCompile(`\d+`)

func main() {
	fileInBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}

	lineItems := bytes.Split(fileInBytes, []byte("\n"))
	slog.Info("Parsed the puzzle inputs into []bytes", "puzzle_input", lineItems)
	timeItems := matcher.FindAll(lineItems[0], -1)
	distanceItems := matcher.FindAll(lineItems[1], -1)

	var answer int = 1
	for x := 0; x < len(timeItems); x++ {
		timeValue, err := strconv.Atoi(string(timeItems[x]))
		if err != nil {
			slog.Error("Error parsing digit for time")
		}
		distanceValue, err := strconv.Atoi(string(distanceItems[x]))
		if err != nil {
			slog.Error("Error parsing digit for distance")
		}
		boatRace := race.NewBoatRace(timeValue, distanceValue)
		answer *= len(boatRace.PossibleWaysToWin())
	}
	// For each boat race get all possible ways you could win
	slog.Info("Done processing part 1", "answer", answer)
	timeValue := bytes.Join(timeItems, []byte(""))
	distanceValue := bytes.Join(distanceItems, []byte(""))
	slog.Info("Done processing part2Values", "timeValue", timeValue, "distanceValue", distanceValue)
	timeIntValue, err := strconv.Atoi(string(timeValue))
	if err != nil {
		panic(err)
	}
	distanceIntValue, err := strconv.Atoi(string(distanceValue))
	if err != nil {
		panic(err)
	}
	br := race.NewBoatRace(timeIntValue, distanceIntValue)
	possibleWaysToWin := br.PossibleWaysToWin()
	slog.Info("Possible ways to win", "possible wins", len(possibleWaysToWin))

}
