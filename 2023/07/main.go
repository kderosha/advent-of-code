package main

import (
	"bytes"
	"log/slog"
	"os"
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

}
