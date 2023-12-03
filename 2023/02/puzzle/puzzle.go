package puzzle

import (
	"github.com/kderosha/advent-of-code/2023/02/game"
)

type Puzzle struct {
	RedLimit int
	GreenLimit int
	BlueLimit int
	Games []game.Game
}