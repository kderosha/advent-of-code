package main

import (
	"log/slog"
	"bufio"
	"os"
	"github.com/kderosha/advent-of-code/2023/02/game"
	"github.com/kderosha/advent-of-code/2023/02/puzzle"
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var puzzle puzzle.Puzzle = puzzle.Puzzle{
		RedLimit: 12,
		GreenLimit: 13,
		BlueLimit: 14,
	}

	// Create file scanner
	scanner := bufio.NewScanner(file)
	// Read line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line items
		// Create a new game from each line
		game, err := game.NewGame(line)
		if err != nil {
			slog.Error("Error processing game item", "game", line, "error", err)
		}
		puzzle.Games = append(puzzle.Games, game)
	}
	sum := 0
	for _, game := range puzzle.Games {
		id := game.Id
		possible := true
		for _, round := range game.Rounds {
			if !round.Possible(puzzle.RedLimit, puzzle.BlueLimit, puzzle.GreenLimit) {
				possible = false
			}
		}
		slog.Info("Outputting if game was possible", "game", game.Id, "rounds", game.Rounds, "possible", possible)
		if possible {
			sum += id
		}
	}
	slog.Info("The final sum of all the possible games is.", "sum", sum)
}