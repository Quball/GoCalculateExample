package calculate

import "testing"
import "../cmd/calculate"

func TestCalculateDivision(t *testing.T) {
	expected := float32(7.5)
	actual, err := calculate.Calculate(15, 2, "DIVIDE")
	if actual != expected {
		t.Error("Test failed")
	}
	if err != nil {
		t.Error("Calculate returned error")
	}
}

func TestCalculateMultiply(t *testing.T) {
	expected := float32(30)
	actual, err := calculate.Calculate(15, 2, "MULTIPLY")
	if actual != expected {
		t.Error("Test failed")
	}
	if err != nil {
		t.Error("Calculate returned error")
	}
}

func TestCalculateAdd(t *testing.T) {
	expected := float32(3)
	actual, err := calculate.Calculate(1, 2, "ADD")
	if actual != expected {
		t.Error("Test failed")
	}
	if err != nil {
		t.Error("Calculate returned error")
	}
}

func TestCalculateSubtract(t *testing.T) {
	expected := float32(-1)
	actual, err := calculate.Calculate(1, 2, "SUBTRACT")
	if actual != expected {
		t.Error("Test failed")
	}
	if err != nil {
		t.Error("Calculate returned error")
	}
}
