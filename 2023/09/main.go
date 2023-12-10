package main

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/kderosha/advent-of-code/2023/09/sequence"
)

func main() {

	fileInBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(fileInBytes, []byte("\n"))
	p1Answer := 0
	for _, line := range lines {
		sequence := sequence.NewSequence(line)
		p1Answer += sequence.Sum()
	}
	slog.Info("Done processing all sequences for p1", "answer", p1Answer)
}
