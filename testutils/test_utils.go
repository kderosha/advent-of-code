package testutils

import "slices"

// Sorts and compares two arrays to test if they are equal
func ArraysAreEqual(array []int, comparedTo []int) bool {
	// Sort both
	slices.Sort(array)
	slices.Sort(comparedTo)
	if len(array) != len(comparedTo) {
		return false
	}
	for x, _ := range array {
		if array[x] != comparedTo[x] {
			return false
		}
	}

	return true

}
