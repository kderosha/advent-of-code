package input

import (
	"bufio"
	"os"
    "regexp"
    "strconv"
)

var numbersRegex *regexp.Regexp = regexp.MustCompile("[0-9]+")

type PuzzleInput struct {
	input []LineItem
}

type LineItem string

// Take a path for the input and return the input as an array of strings
// Each string in the array is 1 line of the input file.
func NewInput(path string) *PuzzleInput {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []LineItem = make([]LineItem, 0)

	// Create file scanner
	scanner := bufio.NewScanner(file)
	// Read line
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, LineItem(line))
	}
    return &PuzzleInput{
        input: input,
    }
}

// Return puzzle input as lines
func (pi *PuzzleInput) LineItems() []LineItem {
    return pi.input
}

func (li LineItem) ConvertToNumbersArray() []int64 {
    numbersAsStrings := numbersRegex.FindAllString(string(li), -1)
    returnTuple := make([]int64, 0, len(numbersAsStrings))
    for _, s := range numbersAsStrings {
        i, err := strconv.ParseInt(s, 10, 64)
        if err != nil {
            panic(err)
        }
        returnTuple = append(returnTuple, i)
    }
    return returnTuple
}
