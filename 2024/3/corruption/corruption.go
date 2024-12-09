package corruption

import (
	"github.com/kderosha/advent-of-code/input"
	"log/slog"
	"regexp"
	"strconv"
)

var re *regexp.Regexp = regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
var do *regexp.Regexp = regexp.MustCompile("do\\(\\)")
var dont *regexp.Regexp = regexp.MustCompile("don't\\(\\)")

var multRe = regexp.MustCompile(`[0-9]+`)

type CorruptionPuzzle struct {
	pi *input.PuzzleInput
}

func NewPuzzle(pi *input.PuzzleInput) *CorruptionPuzzle {
	return &CorruptionPuzzle{pi}
}

func (cp *CorruptionPuzzle) SolutionOne() int64 {
	answer := int64(0)
	allInstancesOfMultRegexp := make([]string, 0)
	slog.Info("Parsing input file", "input", cp.pi.LineItems())

	for _, li := range cp.pi.LineItems() {
		allInstancesOfMultRegexp = append(allInstancesOfMultRegexp, re.FindAllString(string(li), -1)...)
	}

	slog.Info("all found instances of the mult function", "instances", allInstancesOfMultRegexp)

	for _, mult := range allInstancesOfMultRegexp {
		matches := multRe.FindAllString(mult, -1)
		arg1, err := strconv.ParseInt(matches[0], 10, 64)
		if err != nil {
			panic(err)
		}
		arg2, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			panic(err)
		}
		answer += arg1 * arg2
	}
	return answer
}

func (p *CorruptionPuzzle) SolutionTwo() int64 {
	answer := int64(0)
	content := ""
	for _, li := range p.pi.LineItems() {
		content += string(li)
	}
	slog.Info("Combined content together", "content", content)
    dos := make([][]int, 0)
    donts := make([][]int, 0)
    dos = append(dos, []int{0,0})
    dos = append(dos, do.FindAllStringIndex(content, -1)...)
    donts = append(donts, []int{-1, -1})
	donts = append(donts, dont.FindAllStringIndex(content, -1)...)
	// Find all instances of mul()
	multIdxs := re.FindAllStringIndex(content, -1)
	multInstances := re.FindAllString(content, -1)

	for i, multInstance := range multInstances {
		multIdx := multIdxs[i][0]
		doIdx := findNearestBelow(dos, multIdx)
		dontIdx := findNearestBelow(donts, multIdx)
		if doIdx > dontIdx {
			matches := multRe.FindAllString(multInstance, -1)

			arg1, err := strconv.ParseInt(matches[0], 10, 64)
			if err != nil {
				panic(err)
			}
			arg2, err := strconv.ParseInt(matches[1], 10, 64)
			if err != nil {
				panic(err)
			}
			answer += arg1 * arg2
		}
	}

	return answer
}

// Find the nearest integer lower than the target
func findNearestBelow(arr [][]int, target int) int {
    slog.Info("finding nearest below", "arr", arr, "target", target)
	nearest := -9999
	for _, intArr := range arr {
		compareIdx := intArr[0]
		if compareIdx < target && compareIdx >= nearest {
			nearest = compareIdx
		}
	}
    slog.Info("found nearest", "nearest", nearest)
	return nearest
}
