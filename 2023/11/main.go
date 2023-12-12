package main

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/kderosha/advent-of-code/2023/11/universe"
)

func main() {
	fileBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(fileBytes, []byte("\n"))
	u := universe.NewUniverse(lines)
	answer := u.CalculateSumsOfDistanceBetweenGalaxies()
	slog.Info("P1 calculation", "answer", answer)
}
