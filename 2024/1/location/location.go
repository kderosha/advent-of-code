package location

import (
    "github.com/kderosha/advent-of-code/input"
    "regexp"
    "strconv"
    "sort"
    "github.com/adam-lavrik/go-imath/i64"
)


type LocationPuzzle struct {
    locations []Location
    leftNumbers []int64
    rightNumbers []int64

}

type Location struct {
    originalIndex int64
    tuple []int64 // Hold integers side by side
}

func NewLocationPuzzle(pi *input.PuzzleInput) *LocationPuzzle {
    // Parse puzzle input into slice of locations.
    locations := make([]Location, 0, len(pi.LineItems()))
    for idx, lineItem := range pi.LineItems() {
        locations = append(locations, newLocation(int64(idx), lineItem))
    }
    sortByTupleIndex(locations, 0)
    leftNumbers := getNumbersList(locations, 0)
    sortByTupleIndex(locations, 1)
    rightNumbers := getNumbersList(locations, 1)

    // Parse the puzzing input into location items and store them in the puzzle structure
    return &LocationPuzzle{
        locations: locations,
        leftNumbers: leftNumbers,
        rightNumbers: rightNumbers,
    }
}

// Inplace sort of the locations array by the number in the index of the tuple within the location
func sortByTupleIndex(l []Location, tupleIdx int) {
    sort.Slice(l, func(i, j int) bool {
        return l[i].tuple[tupleIdx] < l[j].tuple[tupleIdx]
    })
}

func getNumbersList(l []Location, tupleIdx int) []int64 {
    returnArray := make([]int64, 0, len(l))
    for _, loc := range l {
        returnArray = append(returnArray, loc.tuple[tupleIdx]) 
    }
    return returnArray
}


func newLocation(originalIndex int64, lineItem input.LineItem) Location {
    return Location{
        originalIndex: originalIndex,
        tuple: newTupleFromLineItem(lineItem),
    }
}

func newTupleFromLineItem(li input.LineItem) []int64 {
}


func (p *LocationPuzzle) SolutionOne() int64{
    sum := int64(0)
    for i, _ := range p.locations {
        sum += distanceBetween(p.leftNumbers[i], p.rightNumbers[i])
    }

    // create two sets of integers and sort in ascending order
    return sum
}

func distanceBetween(i, j int64) int64 {
    return i64.Abs(i-j) 
}

func (p *LocationPuzzle) SolutionTwo() int64 {
    sum := int64(0)
    for i, _ := range p.locations{
        sum += p.leftNumbers[i] * appearsIn(p.leftNumbers[i], p.rightNumbers)
    }
    return sum
}

func appearsIn(n int64, numbers []int64) int64 {
    appearences := int64(0)
    for _, num := range numbers {
        if num == n {
            appearences++
        }
    }
    return appearences
}
