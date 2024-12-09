package xmas

import (
	"fmt"
	"github.com/kderosha/advent-of-code/input"
	"log/slog"
)

var xmasLetters []rune = []rune{'X', 'M', 'A', 'S'}
var xDashMasLetters [][]rune = [][]rune{
    {'M', 'M', 'S', 'S'},
    {'S', 'S', 'M', 'M'},
    {'M', 'S', 'S', 'M'},
    {'S', 'M', 'M', 'S'},
}
/*
m s s s 
 a   a
m s m m
*/

var n Direction = Direction{x: 0, y: 1}

var directions []Direction = []Direction{
	n,              // n
	{x: 1, y: 1},   // ne
	{x: 1, y: 0},   // e
	{x: 1, y: -1},  //se
	{x: 0, y: -1},  // s
	{x: -1, y: -1}, // sw
	{x: -1, y: 0},  // w
	{x: -1, y: 1},  // nw
}

var xDashMasDirections []Direction = []Direction{
	{x: -1, y: -1}, // nw
	{x: 1, y: -1},   // ne
	{x: 1, y: 1},  //se
	{x: -1, y: 1}, // sw
}

type Xmas struct {
	pi           *input.PuzzleInput
	letterMatrix *LetterMatrix
}

func New(pi *input.PuzzleInput) *Xmas {
	letterMatrix := &LetterMatrix{lm: make([][]Letter, 0)}
	for y, li := range pi.LineItems() {
		letterLine := make([]Letter, 0)
		for x, character := range string(li) {
			letterLine = append(letterLine, Letter{l: character, v: Vertex{x: x, y: y}})
		}
		letterMatrix.lm = append(letterMatrix.lm, letterLine)
	}
	return &Xmas{
		pi:           pi,
		letterMatrix: letterMatrix,
	}
}

type LetterMatrix struct {
	lm [][]Letter
}

func (a *LetterMatrix) getRowLength() int {
	return len(a.lm[0])
}

func (a *LetterMatrix) getNumOfRows() int {
	return len(a.lm)
}

func (a *LetterMatrix) GetLetter(v Vertex) (Letter, error) {
	if a.InBounds(v) {
		return a.lm[v.y][v.x], nil
	}
	return Letter{}, nil
}

func (a *LetterMatrix) InBounds(v Vertex) bool {
	return v.x >= 0 && v.x < a.getRowLength() && v.y >= 0 && v.y < a.getNumOfRows()
}

type Vertex struct {
	x, y int
}

// Move the vertex to a direction
func (v Vertex) Move(d Direction) Vertex {
	return Vertex{x: v.x + d.x, y: v.y + d.y}
}

type Direction Vertex

type Letter struct {
	l rune
	v Vertex
}

func (let Letter) String() string {
	return fmt.Sprintf("%s, %+v", string(let.l), let.v)
}

func (let Letter) evaluateDirection(letterIdx int, d Direction, lm *LetterMatrix, letters []rune) int {
	slog.Info("Evaluating direction for letter", "letter", let, "direction", d)
	if let.l != letters[letterIdx] {
		slog.Info("Letter does not match letter in xmasLetters", "letter", let, "xmasIdx", letterIdx)
		return 0
	}
	if letterIdx == len(letters)-1 {
		slog.Info("Letter matches and it's the base case index")
		return 1
	}
	newLetter, err := let.moveInDirection(d, lm)
	if err != nil {
		slog.Info("Found error moving in direction", "direction", d, "error", err)
		return 0
	}
	slog.Info("Evaluting new letter in same direction", "newLetter", newLetter, "direction", d)
	return newLetter.evaluateDirection(letterIdx+1, d, lm, letters)
}

// Move in the direction d from vertex v within lm
func (let Letter) moveInDirection(d Direction, lm *LetterMatrix) (Letter, error) {
	// Calculate new vertex in a given direction
	newVertex := let.v.Move(d)
	slog.Info("Moving letter in direction.", "letter", let, "direction", d, "newVertexInDirection", newVertex)
	// Check if that new vertex is in bounds
	if !lm.InBounds(newVertex) {
		return Letter{}, fmt.Errorf("Vertex out of bounds")
	}
	return lm.GetLetter(newVertex)
}

func (let Letter) IsXMas(lm *LetterMatrix, letters []rune) bool {
	slog.Info("Evaluating if letter is xmas", "letter", let)
    for dIdx, direction := range xDashMasDirections {
        slog.Info("Checking letter in direction", "dIdx", dIdx, "direction", direction)
        newVertex := let.v.Move(direction)
        letter, err := lm.GetLetter(newVertex)
        slog.Info("New Letter at new vertex", "newVertex", newVertex, "newLetter", letter)
        if err != nil || letter.l != letters[dIdx]{
            return false
        }
    }
    return true
}

func (xmas *Xmas) SolutionOne() int {
	answer := 0
	for _, letters := range xmas.letterMatrix.lm {
		for _, letter := range letters {
			slog.Info("Evaluating all directions", "letter", letter)
			for _, direction := range directions {
				answer += letter.evaluateDirection(0, direction, xmas.letterMatrix, xmasLetters)
				slog.Info("Result for letter in direction", "answer", answer, "letter", letter, "direction", direction)
			}
		}
	}
	return answer
}

func (xmas *Xmas) SolutionTwo() int {
	answer := 0
	for _, letters := range xmas.letterMatrix.lm {
		for _, letter := range letters {
			if letter.l == 'A' {
                for _, xdlArr := range xDashMasLetters {
                    if letter.IsXMas(xmas.letterMatrix, xdlArr) {
                        answer++
                    } 
                }
			}
		}
	}
	return answer
}
