package report

import (
    "log/slog"
	"github.com/adam-lavrik/go-imath/i64"
	"github.com/kderosha/advent-of-code/input"
)

type Reports struct {
	reports []Report
}

func NewReports(pi *input.PuzzleInput) *Reports {
	// Parse puzzle input into slice of locations.
	reports := make([]Report, 0, len(pi.LineItems()))
	for _, lineItem := range pi.LineItems() {
		reports = append(reports, newReport(lineItem))
	}
	return &Reports{
		reports: reports,
	}
}

type Report struct {
	levels []int64
}

func (r *Report) IsSafe() bool {
    slog.Info("Checking report is safe", "report", r.levels)
	// Detemine if increasing or decreasing state
	increasing := r.levels[0] < r.levels[1]
	for x := 1; x < len(r.levels); x++ {
		if increasing && r.levels[x] < r.levels[x-1] {
			return false
		}
		if !increasing && r.levels[x] > r.levels[x-1] {
			return false
		}
		if !isSafeDifference(r.levels[x], r.levels[x-1]) {
			return false
		}
	}
	return true
}

func isSafeDifference(x, y int64) bool {
	difference := i64.Abs(x - y)
	if difference < 1 || difference > 3 {
		return false
	}
	return true
}

// Create a new report from a line
func newReport(s input.LineItem) Report {
	levels := s.ConvertToNumbersArray()
	return Report{
		levels: levels,
	}
}

func (r *Report) duplicateReportWithMissingIndex(skipIdx int) *Report {
    newReportLevels := make([]int64, 0, len(r.levels) - 1)
    for idx, level := range r.levels {
        if idx != skipIdx {
            newReportLevels = append(newReportLevels, level)
        }
    }
    return &Report{
        levels: newReportLevels,
    }
}

func (r *Reports) SolutionOne() int64 {
    slog.Info("Calculating solution one")
	safeReports := int64(0)
	for _, rs := range r.reports {
		if rs.IsSafe() {
			safeReports++
		}
	}
	return safeReports
}

func (rl *Reports) SolutionTwo() int64 {
    slog.Info("Calculating solution two")
	safeReports := int64(0)
	for _, r := range rl.reports {
        // if original report is safe increase number of safe reports
        if r.IsSafe() {
            safeReports++
        } else {
            // Create a new reports out of each report and check if that new report is safe
            for x := 0; x < len(r.levels); x++ {
                rd := r.duplicateReportWithMissingIndex(x)
                slog.Info("Testing duplicate of report", "duplicate", x, "report permutation", rd.levels) 
                if rd.IsSafe() {
                    safeReports++
                    break;
                }
            }
        }
	}
	return safeReports
}
