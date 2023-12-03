package game

import (
	"strings"
	"log/slog"
	"fmt"
	"strconv"
)
// Holds information about the evaluation of a game.
type Evaluation struct {
	possible bool
	reason string
}

func (e Evaluation) Possible() (bool, string) {
	return e.possible, e.reason
}

// Holds the id of the game and the different rounds played
// If the game has been evaluated. 
type Game struct {
	Id int
	Rounds []Round
	Evaluated bool
	Evaluation Evaluation
}

// Evaluate the game given the limits
func (g Game) Evaluate(redLimit, blueLimit, greenLimit int) Evaluation {
	slog.Info("Evaluating game", "game", g)
	var evaluation Evaluation
	// loop through the rounds and check if the game is possible given the limits
	evaluation.possible = true 	// Assume possible unless proven otherwise
	for ridx, round := range g.Rounds {
		if possible, reason := round.possible(redLimit, blueLimit, greenLimit); !possible{
			evaluation.possible = false
			evaluation.reason = fmt.Sprintf("round %d %s", ridx + 1, reason)
			break;
		} 
	}
	g.Evaluated = true
	slog.Info("game has been evaluated", "possible", evaluation.possible, "reason", evaluation.reason)
	g.Evaluation = evaluation
	return evaluation
}

// Calculate the power of the game.
// --- Part Two ---
// The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

// As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

// Again consider the example games from earlier:

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
// In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
// Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
// Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
// Game 4 required at least 14 red, 3 green, and 15 blue cubes.
// Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.
// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

// For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
func (g Game) Power() int {
	// For every round get the check the color to get the max
	var powerArray [3]int

	for _, round := range g.Rounds {
		if round.colors["red"] > powerArray[0] {
			powerArray[0] = round.colors["red"]
		}
		if round.colors["blue"] > powerArray[1] {
			powerArray[1] = round.colors["blue"]
		}
		if round.colors["green"] > powerArray[2] {
			powerArray[2] = round.colors["green"]
		}
	}
	power := 1
	for _, colorPower := range powerArray {
		power *= colorPower
	}
	return power
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
// Example round [N, color, N, color]
func NewRound(round string) Round {
	round = strings.ReplaceAll(round, ",", "")
	parts := strings.Split(round, " ")
	slog.Info("Parsed parts", "parts", parts)
	colorIdx := 0
	var colorMap map[string]int = make(map[string]int, 0)

	for colorIdx < len(parts) / 2 {
		slog.Info("numberOfColor", "parts[colorIdx]", parts[colorIdx*2])
		number, err := strconv.Atoi(strings.Trim(parts[colorIdx*2], " "))
		if err != nil {
			slog.Info("Error parsing number of colored cubes")
		}

		slog.Info("Color of cubes", "color", parts[colorIdx * 2 + 1])
		color := parts[colorIdx * 2 + 1]

		colorIdx++
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
func (r Round) possible(redLimit, blueLimit, greenLimit int) (bool, string) {
	slog.Info("Checking if round is possible", "round", r)
	// Get red color from Round compare to red limit
	if redColor, ok := r.colors["red"] ; ok {
		if redColor > redLimit {
			return false, "to many red"
			
		}
	}

	if blueColor, ok := r.colors["blue"]; ok {
		if blueColor > blueLimit {
			return false, "to many blue"
		}
	}

	if greenColor, ok := r.colors["green"] ; ok {
		if greenColor > greenLimit {
			return false, "to many green"
		}
	}
	return true, ""
}
