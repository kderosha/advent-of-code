package main

import (
    "github.com/kderosha/advent-of-code/input"
    "github.com/kderosha/advent-of-code/2024/2/report"
    "log/slog"
)

func main() {
	pi := input.NewInput("./input.txt")
    reports := report.NewReports(pi)
	slog.Info("The solution to day 2 puzzle 1", "solution", reports.SolutionOne())
	slog.Info("The solution to day 2 puzzle 2", "solution", reports.SolutionTwo())
}
