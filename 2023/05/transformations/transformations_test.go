package transformations
import (
	"testing"
)

// Test whether values are in the ranges
func TestInRange(t *testing.T) {
	rangeObj := &NumberRange{
		sourceStart: 10,
		rangeSize: 5,
		destinationStart: 15,
	}

	if !rangeObj.inRange(10) {
		t.Fatalf("Unexpected in range value, want true got %v", rangeObj.inRange(12))
	}
	if !rangeObj.inRange(12) {
		t.Fatalf("Unexpected in range value, want true got %v", rangeObj.inRange(12))
	}
	if !rangeObj.inRange(14) {
		t.Fatalf("Unexpected in range value, want true got %v", rangeObj.inRange(12))
	}
	if rangeObj.inRange(15) {
		t.Fatalf("Unexpected value in range, wanted false but got %v", rangeObj.inRange(15))
	}
	if rangeObj.inRange(9) {
		t.Fatalf("Unexpected value in range, wanted false but got %v", rangeObj.inRange(9))
	}
}

func TestRangeTransformation(t *testing.T){
	rangeObj := &NumberRange{
		sourceStart: 10,
		rangeSize: 5,
		destinationStart: 15,
	}

	value, inRange := rangeObj.transform(10)
	if !inRange {
		t.Fatal("Not in range but expected to be in range")
	}
	if value != 15 {
		t.Fatalf("Unexpected transformed value, want 15 but got %d", value)
	}

	value, inRange = rangeObj.transform(12)
	if !inRange {
		t.Fatal("Not in range but expected to be in the range")
	}
	if value != 17 {
		t.Fatalf("Unexpected transformed value, want 17 but got %d", value)
	}

	value, inRange = rangeObj.transform(9)
	if inRange{
		t.Fatal("In range, expected to not be in range")
	}

	value, inRange = rangeObj.transform(15)
	if inRange {
		t.Fatal("In range, expected to not be in range")
	}
}