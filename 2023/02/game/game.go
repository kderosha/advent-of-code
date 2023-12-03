package game

import (
	"strings"
	"log/slog"
	"fmt"
	"strconv"

)
type Game struct {
	Id int
	Rounds []Round
}

func NewGame(game string) (Game, error) {
	colonIndex := strings.IndexByte(game, ':')
	gameSubstring := game[:colonIndex]
	roundsSubstring := game[colonIndex + 1:]
	slog.Info("game processed", "game", game, "colonIndex", colonIndex, "gameSubstring", gameSubstring, "roundsSubstring", roundsSubstring)

	// get id
	id, err := getGameId(gameSubstring)
	if err != nil {
		return Game{}, err
	}

	// create new rounds
	rounds := createRounds(roundsSubstring)

	return Game{
		Id: id,
		Rounds: rounds,
	}, nil
}

// Create rounds from a substring of format N color, N color, N color...;
func createRounds(roundsSubstring string) []Round {
	// split the string into an array of strings by semi colon
	rounds := strings.Split(roundsSubstring, ";")
	numOfRounds := len(rounds)
	slog.Info("Processed rounds substring", "rounds", rounds, "numOfRounds", numOfRounds)
	newRounds := make([]Round, 0)
	for _, round := range rounds {
		newRounds = append(newRounds, NewRound(strings.Trim(round, " ")))
	}
	return newRounds
}

// Parse the round
// Parse all the colors into a map
// Example round
func NewRound(round string) Round {
	round = strings.ReplaceAll(round, ",", "")
	parts := strings.Split(round, " ")
	slog.Info("Parsed parts", "parts", parts)
	//[N, color, N, color]
	tuple := 0
	var colorMap map[string]int = make(map[string]int, 0)

	for tuple < len(parts) / 2 {
		slog.Info("numberOfColor", "parts[tuple]", parts[tuple*2])
		number, err := strconv.Atoi(strings.Trim(parts[tuple*2], " "))
		if err != nil {
			slog.Info("Error parsing number of colored cubes")
		}

		slog.Info("Color of cubes", "color", parts[tuple * 2 + 1])
		color := parts[tuple * 2 + 1]

		tuple++
		colorMap[color] = number
	}

	return Round{
		roundString: round,
		colors: colorMap,
	}

}

// Parse out the game id from the game string
func getGameId(game string) (int, error) {
	// split string
	if idString, found := strings.CutPrefix(game, "Game "); found {
		if id, err := strconv.Atoi(idString); err != nil {
			slog.Error("Error parsing game id", "error", err.Error(), "idString", idString)
			return 0, err
		} else {
			return id, nil
		}
	}
	return 0, fmt.Errorf("Game id not found for game %s", game)
}

type Round struct {
	roundString string
	colors map[string]int
}

func (r Round) String() string {
	return fmt.Sprintf("%s, colors: %+v", r.roundString, r.colors)
}

// Determines if this round is possible with a given puzzle
func (r Round) Possible(redLimit, blueLimit, greenLimit int) bool {
	slog.Info("Checking if round is possible", "round", r)
	// Get red color from Round compare to red limit
	if redColor, ok := r.colors["red"] ; ok {
		if redColor > redLimit {
			return false
		}
	}

	if blueColor, ok := r.colors["blue"]; ok {
		if blueColor > blueLimit {
			return false
		}
	}

	if greenColor, ok := r.colors["green"] ; ok {
		if greenColor > greenLimit {
			return false
		}
	}

	return true
}
