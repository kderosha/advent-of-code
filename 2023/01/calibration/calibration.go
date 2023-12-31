package calibration

import (
	"regexp"
	"log/slog"
	"strconv"
	"fmt"
)
// Regular expression matches all Digits in the input
var regexpP1 *regexp.Regexp = regexp.MustCompile(`\d`)
// Check for p2 regex
var regexpP2 *regexp.Regexp = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

var digitMap  = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
}

type Calibrations []Calibration

// Sum up all of the calibration items 
func (calibrations Calibrations) Sum() int {
	var sum int = 0
	for _, value := range calibrations{
		sum += value.GetNumber()
	}
	return sum
}

type Calibration struct {
	FirstDigit Digit
	SecondDigit Digit
}

func (c Calibration) GetNumber() int {
	// concat digits.
	var concat string = fmt.Sprintf("%s%s", c.FirstDigit, c.SecondDigit)
	// TODO: Convert concat to number
	number, err := strconv.Atoi(concat)
	if err != nil {
		slog.Error("Error converting Calibration to number")
		panic(err)
	}
	return number
}

type Digit string

func (d Digit) String() string {
	return digitMap[string(d)]
}

// Compute the value of the calibrations
func NewCalibration(value string, part string) Calibration {
	var originalValue string = value
	var matches []string 
	var matchLocation []int = make([]int,0) 
	// find all matches
	for matchLocation != nil {
		slog.Info("matchLocation != nil", "originalValue", originalValue, "value", value)
		if part == "1" {
			matchLocation = regexpP1.FindStringIndex(value)
		} else {
			matchLocation = regexpP2.FindStringIndex(value)
		}
		// Found match
		if matchLocation != nil {
			// Take the match location and get the match from the string value
			matches = append(matches, value[matchLocation[0]:matchLocation[1]])
			slog.Info("Evaluation of new value", "matchLocation[0]", matchLocation[0], "len(value)", len(value))
			if (matchLocation[0] + 1 < len(value)){
				value = value[(matchLocation[0] + 1):]
			} else {
				value = ""
			}
		}
		slog.Info("matches so far", "matches", matches)
	}
	slog.Info("matched digits", "value", originalValue, "matches", matches)
	
	matchOne := Digit(matches[0])
	lastMatch := Digit(matches[len(matches) - 1])
	
	var calibration Calibration = Calibration{
		FirstDigit: matchOne,
		SecondDigit: lastMatch,
	}
	slog.Info("Calibration's calculated number", "number", calibration.GetNumber())
	return calibration
}
