package main

import (
    "github.com/kderosha/advent-of-code/input"
    "github.com/kderosha/advent-of-code/2024/1/location"
    "log/slog"
)

func main() {
    pi := input.NewInput("./input.txt")
    lp := location.NewLocationPuzzle(pi)
    slog.Info("The solution to day 1 puzzle 1", "solution", lp.SolutionOne())
    slog.Info("The solution to day 1 puzzle 2", "solution", lp.SolutionTwo())
}
