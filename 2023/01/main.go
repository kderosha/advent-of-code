package main

import (
	"os"
	"github.com/kderosha/advent-of-code/2023/01/calibration"
	"bufio"
	"log/slog"
)

func main(){
	// Open file
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var calibrations calibration.Calibrations = make(calibration.Calibrations, 0)

	// Create file scanner
	scanner := bufio.NewScanner(file)
	// Read line
	for scanner.Scan() {
		line := scanner.Text()
		calibrations = append(calibrations, calibration.NewCalibration(line, "2"))
	}
	slog.Info("Sum of all calibration values", "answer", calibrations.Sum())
}