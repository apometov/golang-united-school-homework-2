package square

import (
	"testing"
)

func TestCalcSquare(t *testing.T) {
	const delta = 0.0000000001
	var result float64

	// test for 0 sides
	result = CalcSquare(10.0, SidesCircle)
	if result-314.159265358979323 > delta {
		t.Errorf("Expected 314.159265358979323, but got %f", result)
	}

	// test for 3 sides
	result = CalcSquare(10.0, SidesTriangle)
	if result-43.30127018922193 > delta {
		t.Errorf("Expected 43.30127018922193, but got %f", result)
	}

	// test for 4 sides
	result = CalcSquare(10.0, SidesSquare)
	if result-100 > delta {
		t.Errorf("Expected 100, but got %f", result)
	}

	// test for other count of sides
	result = CalcSquare(10.0, 5)
	if result != 0 {
		t.Errorf("Expected 0, but got %f", result)
	}
}
