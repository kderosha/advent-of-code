package main

import (
	"github.com/kderosha/advent-of-code/2024/4/xmas"
	"github.com/kderosha/advent-of-code/input"
    "log/slog"
)

func main(){
    pi := input.NewInput("./input.txt")
    puzzle := xmas.New(pi)
    slog.Info("Solution to puzzle one", "solution 1", puzzle.SolutionOne())
    slog.Info("Starting solution 2")
    slog.Info("Solution to puzzle one", "solution 2", puzzle.SolutionTwo())
}
