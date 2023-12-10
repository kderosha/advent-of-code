package sequence

import (
	"bytes"
	"log/slog"
	"slices"
	"strconv"
)

type Sequence struct {
	sequence  []int
	sum       int
	sequences [][]int
}

func NewSequence(sequenceArrayBytes []byte) *Sequence {
	// Split the bytes array by spaces
	individualNumbersArray := bytes.Split(sequenceArrayBytes, []byte(" "))
	processedNumbers := make([]int, 0, len(individualNumbersArray))

	for _, numberBytes := range individualNumbersArray {
		value, err := strconv.Atoi(string(numberBytes))
		if err != nil {
			slog.Error("Error processing numbers", "numbers", value)
		}
		processedNumbers = append(processedNumbers, value)
	}
	slog.Info("numbers for sequence has been processed", "numbers", processedNumbers)
	slices.Reverse(processedNumbers)
	sequence := &Sequence{sequence: processedNumbers}
	sequence.sequences = append(sequence.sequences, processedNumbers)
	sequence.calculateSequences()
	return sequence
}

func (s *Sequence) Sum() int {
	return s.sum
}

// Calculate the sum of all the added values
func (seq *Sequence) calculateSequences() {
	currentSequence := seq.sequence
	allZeros := false
	for !allZeros {
		newSequence := make([]int, len(currentSequence)-1)
		allZeros = true
		for x := range newSequence {
			newSequence[x] = -1 * (currentSequence[x+1] - currentSequence[x])
			if newSequence[x] != 0 {
				allZeros = false
			}
		}
		seq.sequences = append(seq.sequences, newSequence)
		currentSequence = newSequence
	}
	slog.Info("All sequences are calculated", "sequences", seq.sequences)
	// append a 0 to the current sequence and continue
	currentSequence = append(currentSequence, 0)
	// Work backwards from the sequences and add numbers to the next sequence
	runningSum := 0
	for x := len(seq.sequences) - 1; x > 0; x-- {
		nextValueForPreviousSequence := lastValueOfSlice(seq.sequences[x-1]) - lastValueOfSlice(currentSequence)
		seq.sequences[x-1] = append(seq.sequences[x-1], nextValueForPreviousSequence)
		runningSum += nextValueForPreviousSequence
		currentSequence = seq.sequences[x-1]
	}
	slog.Info("Additions of values for each sequence is done", "sequences", seq.sequences)
	seq.sum = lastValueOfSlice(currentSequence)
	slog.Info("Sum of all added values", "sum", seq.sum)
}

func lastValueOfSlice(slice []int) int {
	return slice[len(slice)-1]
}
