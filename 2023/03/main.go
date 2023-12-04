package main

import (
	"os"
	"bufio"
	"unicode"
	"strconv"
	"log/slog"
)

type encodedMatrix [][]int

func main() {
	// We are going to read each character by character
	// -1 represents a symbol
	var matrix encodedMatrix = make([][]int, 0)

	// For each line process line into an []int
	// Open puzzle input
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	currentNumberIdx := 1
	var numbers []int = make([]int, 0)
	// Create file scanner
	scanner := bufio.NewScanner(file)
	// Read line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line items
		// Create a new game from each line
		var processed []int = make([]int, len(line), len(line))
		number := ""
		lengthOfLine := len(line)
		for x := 0; x < lengthOfLine; x++ {
			character := rune(line[x])
			if unicode.IsNumber(character) {
				processed[x] = currentNumberIdx
				number += string(character)
			} else {
				// Process number here 
				if len(number) > 0 {
					numberInt, err := strconv.Atoi(number)
					if err != nil {
						slog.Error("Error parsing the number")
					}
					numbers = append(numbers, numberInt)
					number = ""
					currentNumberIdx++
				}
				if character == '.' {
					processed[x] = 0
				} else {
					processed[x] = -1
				}
			} 
		}
		// line ended in number
		if len(number) > 0 {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				slog.Error("Error parsing the number")
			}
			numbers = append(numbers, numberInt)
			currentNumberIdx++
		}
		matrix = append(matrix, processed)
	}
	slog.Info("Encoded version of the matrix ", "matrix", matrix, "numbers", numbers)
	slog.Info("Procesing encoded matrix")
	indexes := make([]int, 0)
	for rowIdx, row := range matrix {
		for columnIdx, _ := range row {
			// Found symbol
			if matrix[rowIdx][columnIdx] == -1 {
				for subRowIdx := rowIdx - 1; subRowIdx < rowIdx + 2; subRowIdx++ {
					for subColumnIdx := columnIdx - 1; subColumnIdx < columnIdx + 2; subColumnIdx++ {
						// Index out of bounds prevention
						if subRowIdx >= 0 && subRowIdx < len(matrix) && subColumnIdx >= 0 && subColumnIdx < len(row){
							value := matrix[subRowIdx][subColumnIdx] 
							if value > 0 {
								// Add value to an array of number indexes if it doesn't yet exist.
								exists := false
								for _,index := range indexes {
									if value == index {
										exists = true
									}
								}
								if !exists {
									indexes = append(indexes, value)
								}
							}
						}
					}
				}
			}
		}
	}
	slog.Info("Indexes of numbers that touch a symbol", "indexes", indexes)
	sum := 0
	for _, index := range indexes {
		sum += numbers[index-1]
	}
	slog.Info("Done processing", "sum", sum)
}