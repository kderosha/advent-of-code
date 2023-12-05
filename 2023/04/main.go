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
	var cardMap map[int]*card.Card = make(map[int]*card.Card)
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line items
		// Create new card
		processedCard := card.NewCard(line)
		score := processedCard.CalculateScore()
		cardMap[processedCard.Id] = processedCard
		finalScore += score
	}
	slog.Info("Finished processing puzzle", "score", finalScore)
	part2Score := 0
	for _, value := range cardMap {
		part2Score += value.SubtreeSize(cardMap)
	}
	slog.Info("Finished part 2", "score", part2Score)

}

/*
calculate all trees root and level 1
tree structure to store solutions for each card.

            1
	2	  3	    4	   5
 3	  4  4 5    5
4 5  5   5
5

				2
			3      4
		  4  5    5
		  5  
		  */