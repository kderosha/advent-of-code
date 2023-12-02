package calibration

import (
	"testing"
)

func TestCalibrationValueCompute(t *testing.T) {
	var cal1 Calibration = NewCalibration("0aksdjfowijdf9", "1")
	var cal2 Calibration = NewCalibration("1jsi8294jaskdf", "2")
	var calStringValue Calibration = NewCalibration("1jsi8294jaskdffive", "2")

	if (cal1.GetNumber() != 9) {
		t.Fatalf("Cal1 was not expected, wanted 9 but got %d", cal1.GetNumber())
	}
	if (cal2.GetNumber() != 14) {
		t.Fatalf("Cal2 was not expected, wanted 14 but got %d", cal1.GetNumber())
	}
	if calStringValue.GetNumber() != 15 {
		t.Fatalf("Cal String value was not expected, wanted 15 but got %d", calStringValue.GetNumber())
	}
}

func TestCalibrationsCompute(t *testing.T) {
	var cal1 Calibration = NewCalibration("0aksdjfowijdf9", "1")
	var cal2 Calibration = NewCalibration("1jsi8294jaskdf", "1")
	var calibrations Calibrations = make(Calibrations, 0)
	calibrations = append(calibrations, cal1, cal2)
	if (calibrations.Sum() != 23) {
		t.Fatalf("Calibrations compute was not expected, wanted 23 but got %d", calibrations.Sum())
	}
}

func TestCalibrationsSumPart2(t *testing.T){
	var cal1 Calibration = NewCalibration("0aksdjfowijdf9", "2")
	var cal2 Calibration = NewCalibration("1jsi8294jaskdfnine", "2")
	var cal3 Calibration = NewCalibration("twojsi8294jaskdfnine", "2")

	if cal1.GetNumber() != 9 {
		t.Fatalf("Cal1 was not expected, wanted 9 but got %d", cal1.GetNumber)
	}
	if cal2.GetNumber() != 19 {
		t.Fatalf("Cal2 was not expected, wanted 19 but got %d", cal2.GetNumber())
	}
	if cal3.GetNumber() != 29 {
		t.Fatalf("Cal2 was not expected, wanted 29 but got %d", cal3.GetNumber())
	}

	var calibrations Calibrations = make(Calibrations, 0)
	calibrations = append(calibrations, cal1, cal2, cal3)
	// All 9 + 19 + 29 = 57
	if (calibrations.Sum() != 57) {
		t.Fatalf("Calibrations compute was not expected, wanted 57 but got %d", calibrations.Sum())
	}
}