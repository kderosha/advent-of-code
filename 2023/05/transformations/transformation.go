package transformations

import (
	"log/slog"
	"bytes"
	"regexp"
	"fmt"
	"strconv"
)

var mapRegex *regexp.Regexp = regexp.MustCompile(`(.*)-to-(.*) map`)

type TransformationChain map[Source]*transformation


// Utilize the linked map as a linked list. Transform the the number using next transformation until there are no more transformations.
// start with a source transformation
func (tc TransformationChain) Transform(number int, s Source) int {
	t, exists := tc[s]
	for exists {
		// Next transformation is the 
		number = t.transform(number)
		t, exists = tc[Source(t.d)]
	}
	return number
}

func NewTransformationChain(input []byte) TransformationChain {
	var returnT TransformationChain = make(TransformationChain, 0)

	// find next instance of double \n\n
	for idx, transformationBytes := range bytes.Split(input, []byte("\n\n")) {
		slog.Info("looping transformations parsed", "index", idx, "transformation", transformationBytes)
		parsedTransformation := newTransformation(transformationBytes)
		returnT[parsedTransformation.s] = parsedTransformation
	}

	// Take the input and iterate over the different transformation maps.
	return returnT
}

type Source string
type destination string

// Hold the Source-to-Destination mapping and the ranges of each transformation map
type transformation struct {
	s Source
	d destination
	ranges []*NumberRange
}

func newTransformation(input []byte) *transformation {
	// parse regexp of the map name in order to get the source and destination
	inputLines := bytes.Split(input, []byte("\n"))
	slog.Info("Split transformation by endlines", "split", fmt.Sprintf("%+v", inputLines))
	processedMapValues := mapRegex.FindAllSubmatch(inputLines[0], -1)
	slog.Info("Processed map source-to-dest map", "next", processedMapValues[0][0], "next2", processedMapValues[0][1], "next2", processedMapValues[0][2])
	sourceValue := Source(processedMapValues[0][1])
	destinationValue := destination(processedMapValues[0][2])
	numRanges := make([]*NumberRange, 0)
	// For each input lines we process them into ranges
	for _, rangeByteArray := range inputLines[1:] {
		if len(rangeByteArray) > 0 {
			numRanges = append(numRanges, newRange(rangeByteArray))
		}
	}

	return &transformation{
		s: sourceValue,
		d: destinationValue,
		ranges: numRanges,
	}
}

func (t *transformation) transform(number int) int {
	// find if number is in a range. If it is use range translation formula
	for _, numberRange := range t.ranges {
		// If the number exists in any of the ranges. Use that range to transform number to destination number
		if value, inRange := numberRange.transform(number); inRange {
			return value
		}
	}
	return number
}

var digitRegexp *regexp.Regexp = regexp.MustCompile(`\d+`)

// Holds the parsed out range
// example:
// 1 200 10
// Means source numbers 200-210 map to 1-11
// TODO: formula for this should be determined
type NumberRange struct {
	destinationStart int
	sourceStart int
	rangeSize int
}

func newRange(input []byte) *NumberRange{
	// Should return a [][]byte containing digits. We need to convert each index to a number and store them in the appropriate variables.
	slog.Info("Input for a new number range", "input", input)
	values := digitRegexp.FindAll(input, -1)
	slog.Info("Parsed number range values", "values", values)
	destinationStart, err := strconv.Atoi(string(values[0]))
	if err != nil {
		slog.Error("Error parsing destination start for range", "input", string(input), "attempted parse value", values[0])
	}
	sourceStart, err := strconv.Atoi(string(values[1]))
	if err != nil {
		slog.Error("Error parsing source start for range", "input", string(input), "attempted parse value", values[1])
	}
	rangeSize, err := strconv.Atoi(string(values[2]))
	if err != nil {
		slog.Error("Error parsing range size for range", "input", string(input), "attempted parse value", values[2])
	}

	return &NumberRange{
		destinationStart, sourceStart, rangeSize,
	}
}

func (r *NumberRange) inRange(number int) bool {
	if number >= r.sourceStart && number < r.sourceStart + r.rangeSize {
		return true
	}
	return false
}

// Transforms a number to the destination number
func (r *NumberRange) transform(number int) (int, bool) {
	if r.inRange(number) {
		deltaFromStart := (number - r.sourceStart)
		return r.destinationStart + (deltaFromStart), true
	} else {
		return -1, false
	}
}