package main

import (
	"bytes"
	"log/slog"
	"os"
	"sort"
	"strconv"

	"github.com/kderosha/advent-of-code/2023/07/hand"
)

func main() {
	fileInBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	// Separate into array of array of bytes to represent each hand
	handsBytes := bytes.Split(fileInBytes, []byte("\n"))
	hands := make([]*hand.Hand, len(handsBytes))

	for x, _ := range hands {
		cardsBytes := handsBytes[x][:5]
		bidBytes := handsBytes[x][6:]
		bidInt, err := strconv.Atoi(string(bidBytes))
		if err != nil {
			slog.Error("Error processing bid", "bid", bidInt)
		}
		slog.Info("Processed the hand", "cards", cardsBytes, "bid", bidBytes)
		hands[x] = hand.NewHand(string(handsBytes[x][:5]), bidInt)
	}

	// Sort the sorted rank order
	sort.SliceStable(hands, func(i, j int) bool {
		// Same type we need to compare individual cards in hand order ranks
		if hands[i].Type() == hands[j].Type() {
			// Check the card ranks.
			for x := 0; x < 5; x++ {
				if hands[i].Cards[x].Rank() != hands[j].Cards[x].Rank() {
					return hands[i].Cards[x].Rank() < hands[j].Cards[x].Rank()
				}
			}
		}
		return hands[i].Type() < hands[j].Type()
	})

	slog.Info("hands sorted in order", "hands", hands)
	answer := 0
	for x, hand := range hands {
		slog.Info("hand is being processed into answer", "x", x, "hand", hand, "answer", answer)
		answer += hand.Bid() * (x + 1)
	}
	slog.Info("Part 1 hands are sorted and processed.", "answer", answer)

}
