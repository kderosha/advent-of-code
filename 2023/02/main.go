package main

import (
	"strings"
	"log/slog"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type Puzzle struct {
	redLimit int
	greenLimit int
	blueLimit int
	Games []Game
}
func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var puzzle Puzzle = Puzzle{
		redLimit: 12,
		greenLimit: 13,
		blueLimit: 14,
	}

	// Create file scanner
	scanner := bufio.NewScanner(file)
	// Read line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line items
		// Create a new game from each line
		game, err := NewGame(line)
		if err != nil {
			slog.Error("Error processing game item", "game", line, "error", err)
		}
		puzzle.Games = append(puzzle.Games, game)
	}
}