package main

import (
    "log/slog"
    "github.com/kderosha/advent-of-code/input"
    "github.com/kderosha/advent-of-code/2024/3/corruption"
)

func main() {
	pi := input.NewInput("./input.txt")
    reports := corruption.NewPuzzle(pi)
	slog.Info("The solution to day 2 puzzle 1", "solution", reports.SolutionOne())
	slog.Info("The solution to day 2 puzzle 2", "solution", reports.SolutionTwo())

}
