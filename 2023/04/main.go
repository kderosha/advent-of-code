package main

import (
	"github.com/kderosha/advent-of-code/2023/04/card"
	"os"
	"bufio"
	"log/slog"
)

func main() {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	// Create file scanner
	scanner := bufio.NewScanner(file)

	finalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line items
		// Create new card
		processedCard := card.NewCard(line)
		score := processedCard.CalculateScore()
		finalScore += score
	}
	slog.Info("Finished processing puzzle", "score", finalScore)
}
